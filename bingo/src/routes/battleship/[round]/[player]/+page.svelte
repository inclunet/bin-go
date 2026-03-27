<script>
	// @ts-nocheck
	import { page } from '$app/stores';
	import { get } from 'svelte/store';
	import { onMount, tick } from 'svelte';
	import PageTitle from '$lib/PageTitle.svelte';
	import BattleshipBoard from '$lib/battleship/BattleshipBoard.svelte';
	import BattleshipScoreboard from '$lib/battleship/BattleshipScoreboard.svelte';
	import BattleshipStatus from '$lib/battleship/BattleshipStatus.svelte';
	import AdManager from '$lib/ads/AdManager.svelte';
	import AdProtection from '$lib/ads/AdProtection.svelte';
	import { adConfig } from '$lib/ads/adConfig.js';

	const BOARD_SIZE = 10;
	const p = get(page);
	let roundParam = p.params.round;
	const playerParam = (p.params.player || 'a').toLowerCase();
	const localPlayer = playerParam === 'b' ? 'b' : 'a';

	function createEmptyBoard() {
		return Array.from({ length: BOARD_SIZE }, () =>
			Array.from({ length: BOARD_SIZE }, () => ({ state: 'empty' }))
		);
	}

	function normalizeBoard(boardData) {
		if (!Array.isArray(boardData)) {
			return createEmptyBoard();
		}

		return Array.from({ length: BOARD_SIZE }, (_, rowIdx) => {
			const row = boardData[rowIdx] || [];
			return Array.from({ length: BOARD_SIZE }, (_, colIdx) => {
				const cell = row[colIdx];
				if (cell && typeof cell === 'object') {
					return { ...cell };
				}
				return { state: 'empty' };
			});
		});
	}

	// Atualiza tabuleiro existente mutando células para evitar recriação de nós DOM e perda de foco
	// Abordagem mais idiomática Svelte: gerar um novo board mantendo objetos/células iguais quando não mudaram.
	// Retorna o mesmo array original se nada mudou (evita render). Senão, retorna novo array com linhas e células novas só onde necessário.
	function mergeBoard(oldBoard, incoming) {
		if (!Array.isArray(incoming)) return oldBoard;
		let anyChange = false;
		const next = new Array(BOARD_SIZE);
		for (let r = 0; r < BOARD_SIZE; r++) {
			const oldRow = oldBoard[r];
			const inRow = incoming[r] || [];
			let rowChanged = false;
			const newRow = new Array(BOARD_SIZE);
			for (let c = 0; c < BOARD_SIZE; c++) {
				const inCell = inRow[c];
				if (!inCell || typeof inCell !== 'object') {
					const prev = oldRow[c];
					if (prev.state !== 'empty' || prev.shipId || prev.name) {
						rowChanged = true;
						newRow[c] = { state: 'empty' };
					} else {
						newRow[c] = prev; // reusa objeto
					}
					continue;
				}
				const prev = oldRow[c];
				const ns = inCell.state || 'empty';
				const sid = inCell.shipId;
				const nname = inCell.name;
				if (prev.state === ns && prev.shipId === sid && prev.name === nname) {
					newRow[c] = prev; // sem mudança
				} else {
					rowChanged = true;
					newRow[c] = { state: ns, shipId: sid, name: nname };
				}
			}
			if (!rowChanged) {
				next[r] = oldRow; // preserva referência de linha inteira
			} else {
				anyChange = true;
				next[r] = newRow;
			}
		}
		return anyChange ? next : oldBoard;
	}

	let myBoard = createEmptyBoard();
	let enemyBoard = createEmptyBoard();
	let currentBoard = myBoard;
	let viewMode = 'my';

	let row = 0;
	let col = 0;

	let phase = 'setup';
	let currentPlayer = 'a';
	let winner = '';
	let shipsRemaining = { a: 5, b: 5 };
	// Contadores derivados para exibir nas abas
	$: myRemaining = (localPlayer === 'a') ? shipsRemaining.a : shipsRemaining.b;
	$: enemyRemaining = (localPlayer === 'a') ? shipsRemaining.b : shipsRemaining.a;
	let playerAReady = false;
	let playerBReady = false;
	let scoreA = 0;
	let scoreB = 0;
	let scoreDraw = 0;
	let hasPlayerA = false;
	let hasPlayerB = false;
	let redirecting = false;

	let shipOrientation = 'horizontal';
	let currentShipIndex = 0;
	let shipsToPlace = [
		{ id: 1, name: 'Porta-aviões', size: 5, placed: false },
		{ id: 2, name: 'Encouraçado', size: 4, placed: false },
		{ id: 3, name: 'Cruzador', size: 3, placed: false },
		{ id: 4, name: 'Submarino', size: 3, placed: false },
		{ id: 5, name: 'Destroyer', size: 2, placed: false }
	];
	// Declarar antes de uso reativo em $: shipNames = ...
	let shipNames = {}; // { id: nome } de navios já posicionados para acessibilidade

	let liveAnnounce = '';
	let boardInitialized = false;
	let lastAction = '';
	let lastSunkShip = null; // dados do último navio afundado vindos do servidor
	let statusMessage = '';
	let ws;
	let wsStatus = 'connecting'; // connecting | open | closed | reconnecting
	let wsAttempts = 0;
	let wsMaxAttempts = 10;
	let wsHeartbeat; // interval id
	let lastWsMessage = Date.now();
	let wsConnecting = false; // trava reentrância
	let wsSeq = 0; // sequence para comandos
	const pendingCmds = new Map(); // id -> {resolve,reject,timer,type}
	const wsOpenWaiters = []; // aguardando abertura do WS

	function whenWsOpen(maxWait=4000) {
		if (ws && ws.readyState === WebSocket.OPEN) return Promise.resolve();
		return new Promise((resolve, reject) => {
			const timer = setTimeout(() => {
				const idx = wsOpenWaiters.findIndex(w => w.resolve === resolve);
				if (idx >= 0) wsOpenWaiters.splice(idx,1);
				reject(new Error('timeout ws open'));
			}, maxWait);
			wsOpenWaiters.push({ resolve, reject, timer });
			if (!wsConnecting) connectWebSocket(false);
		});
	}

	function setWsStatus(s) { wsStatus = s; }

	function scheduleHeartbeat() {
		clearInterval(wsHeartbeat);
		// Intervalo mais longo; não fecha agressivamente para evitar reconexões em excesso
		wsHeartbeat = setInterval(() => {
			if (!ws || ws.readyState !== WebSocket.OPEN) return;
			const silentMs = Date.now() - lastWsMessage;
			// Apenas loga silêncio prolongado; deixa servidor decidir encerrar.
			if (silentMs > 120000 && debugMode) {
				console.debug('WS silencioso há', Math.round(silentMs/1000),'s');
			}
			// Envia ping leve (servidor deve ignorar se não implementado)
			try { ws.send('{"type":"ping"}'); } catch {}
		}, 25000);
	}

	function connectWebSocket(initial=false) {
		if (wsConnecting) return; // reentrada
		// Evita múltiplas conexões: se já existe em estado CONNECTING (0), OPEN (1) ou CLOSING (2), aguarda.
		if (ws && (ws.readyState === WebSocket.OPEN || ws.readyState === WebSocket.CONNECTING || ws.readyState === WebSocket.CLOSING)) return;
		wsConnecting = true;
		if (initial) { wsAttempts = 0; }
		const attempt = wsAttempts++;
		const backoff = Math.min(1000 * Math.pow(2, attempt), 15000);
		setWsStatus(attempt === 0 ? 'connecting' : 'reconnecting');
		try {
			ws = new WebSocket(`${location.protocol === 'https:' ? 'wss' : 'ws'}://${location.host}/ws/battleship/${roundParam}/${playerParam}`);
			ws.onopen = () => {
				wsConnecting = false;
				setWsStatus('open');
				lastWsMessage = Date.now();
				scheduleHeartbeat();
				// Sync completo após (re)conexão para garantir que hits recebidos durante desconexão apareçam
				loadGameData();
				while (wsOpenWaiters.length) { // resolver aguardando
					const w = wsOpenWaiters.shift();
					clearTimeout(w.timer);
					w.resolve();
				}
			};
			ws.onmessage = (event) => {
				lastWsMessage = Date.now();
				try {
					const raw = event.data;
					let data;
					try { data = JSON.parse(raw); } catch { return; }
					if (debugMode) {
						console.debug('[WS payload]', data);
						if (data?.enemyBoard) {
							const counts = { hit:0, miss:0, sunk:0, ship:0 };
							for (const row of data.enemyBoard || []) for (const cell of row || []) if (cell && counts.hasOwnProperty(cell.state)) counts[cell.state]++;
							console.debug('[WS enemyBoard counts]', counts);
						}
						if (data?.lastShot) console.debug('[WS lastShot]', data.lastShot, 'lastSunkShip', data.lastSunkShip);
					}
					// Roteamento de mensagens
					if (data && data.type) {
						const { type, id } = data;
						if (type === 'shootResult' || type === 'placeResult') {
							if (data.state) applyRoundState(data.state);
							if (id && pendingCmds.has(id)) {
								const { resolve, timer } = pendingCmds.get(id);
								clearTimeout(timer);
								pendingCmds.delete(id);
								resolve(data);
							}
							return;
						} else if (type === 'error') {
							if (id && pendingCmds.has(id)) {
								const { reject, timer } = pendingCmds.get(id);
								clearTimeout(timer);
								pendingCmds.delete(id);
								reject(new Error(data.message || 'erro'));
							}
							liveAnnounce = data.message || 'Erro de comando';
							return;
						} else if (type === 'state') {
							applyRoundState(data);
							return;
						}
					}
					// Broadcast antigo sem type
					if (data && data.phase && data.currentPlayer) {
						applyRoundState(data);
					}
				} catch (err) {
					console.error('Falha WS msg', err);
				}
			};
			ws.onerror = () => {
				if (wsStatus === 'open') { console.warn('WS error - mantendo tentativa'); }
			};
			ws.onclose = () => {
				wsConnecting = false;
				clearInterval(wsHeartbeat);
				setWsStatus('closed');
				if (wsAttempts <= wsMaxAttempts) {
					setTimeout(() => connectWebSocket(false), backoff);
				} else {
					console.error('WS reconnection stopped after max attempts');
				}
			};
		} catch (e) {
			wsConnecting = false;
			console.error('Erro iniciando WS', e);
			setTimeout(() => connectWebSocket(false), backoff);
		}
	}

	// ======== BLOCO CORRIGIDO (variáveis e utilitários antes estavam dentro de onMount) ========
	let boardMode = 'setup';
	let announcedAllShips = false; // evita repetir anúncio final
	let placingShip = false; // trava reentrância
	let debugMode = false; // modo debug (ativado via ?debug=1)
	let pendingShots = new Set(); // tiros aguardando resposta
	let isMyTurn = false; // calculado reativamente
	let prevViewMode = viewMode; // para detectar mudança
	let manualViewOverride = false; // usuário alterou manualmente a visão (inibe autoswitch até próximo evento relevante)
	let lastLocalShot = null; // {row,col} do último tiro enviado localmente aguardando broadcast
	let prevIsMyTurn = false; // rastrear transição de turno
	let prevLastShotSig = ''; // evita anunciar o mesmo tiro duas vezes (HTTP + WS)

	function getCurrentShip() { return shipsToPlace[currentShipIndex]; }
	function allShipsPlaced() { return shipsToPlace.every(s => s.placed); }
	function isMyTurnCheck() { return phase === 'playing' && currentPlayer === localPlayer; }
	function updateTurnView() {
		boardMode = phase === 'setup' ? 'setup' : 'playing';
	}

	// Reatividade derivada centralizada
	$: shipNames = shipsToPlace.filter(s=>s.placed).reduce((acc,s)=>{acc[s.id]=s.name;return acc;},{});
	$: if (viewMode !== prevViewMode) { prevViewMode = viewMode; }
	$: currentBoard = viewMode === 'my' ? myBoard : enemyBoard;
	$: isMyTurn = isMyTurnCheck();
	$: updateTurnView();

	onMount(() => {
		// Detecta modo debug via query (?debug=1)
		try {
			const sp = new URLSearchParams(window.location.search);
			debugMode = sp.get('debug') === '1';
		} catch {}
		(async () => {
			await ensureRound();
			await loadGameData();
			connectWebSocket(true);
		})();
		return () => { try { ws?.close(); } catch {}; clearInterval(wsHeartbeat); };
	});
	// ======== FIM BLOCO CORRIGIDO ========

	function updateCurrentShipIndex() {
		const nextIndex = shipsToPlace.findIndex((ship) => !ship.placed);
		currentShipIndex = nextIndex === -1 ? shipsToPlace.length : nextIndex;
	}

	function focusActiveCell() {
		const element = document.getElementById(`battleship-cell-${row}-${col}`);
		element?.focus();
	}

	function handleCellFocus(event) {
		const { row: newRow, col: newCol } = event.detail;
		row = newRow;
		col = newCol;
	}

	function announceShipStatus() {
		const remaining = shipsToPlace.filter((ship) => !ship.placed);
		if (!remaining.length) {
			return 'Todos os navios já foram posicionados';
		}

		const current = getCurrentShip();
		const remainingCount = remaining.length;
		const base = remainingCount === 1
			? 'Resta somente um navio para posicionar'
			: `Restam ${remainingCount} navios para posicionar`;

		if (!current) {
			return base;
		}

		return `${base}. Agora posicione o ${current.name} com tamanho ${current.size}.`;
	}

	function handleKeydown(event) {
		const nativeEvent = event.detail.event;
		if (!nativeEvent) return;

		let nextRow = row;
		let nextCol = col;
		let handled = false;

		switch (nativeEvent.key) {
			case 'ArrowUp':
				nextRow = Math.max(0, row - 1);
				handled = true;
				break;
			case 'ArrowDown':
				nextRow = Math.min(BOARD_SIZE - 1, row + 1);
				handled = true;
				break;
			case 'ArrowLeft':
				nextCol = Math.max(0, col - 1);
				handled = true;
				break;
			case 'ArrowRight':
				nextCol = Math.min(BOARD_SIZE - 1, col + 1);
				handled = true;
				break;
			case 'Home':
				nextRow = 0;
				nextCol = 0;
				handled = true;
				break;
			case 'End':
				nextRow = BOARD_SIZE - 1;
				nextCol = BOARD_SIZE - 1;
				handled = true;
				break;
			case 'Enter':
			case ' ':
				if (phase === 'setup') {
					placeShip(row, col);
				} else if (phase === 'playing' && viewMode === 'enemy') {
					shoot(row, col);
				}
				handled = true;
				break;
			case 'r':
			case 'R':
				if (phase === 'setup') {
					rotateShip();
					handled = true;
				}
				break;
			case 't':
			case 'T':
				if (phase === 'playing') {
					toggleView();
					handled = true;
				}
				break;
			case 's':
			case 'S':
				if (phase === 'setup') {
					liveAnnounce = announceShipStatus();
					handled = true;
				}
				break;
			case 'h':
			case 'H':
				if (phase === 'setup') {
					liveAnnounce = 'Use as setas para navegar, Enter para posicionar, R para rotacionar, S para status dos navios e H para ajuda.';
				} else if (phase === 'playing') {
					liveAnnounce = 'Use as setas para navegar, Enter para atirar, T para alternar o tabuleiro e H para ajuda.';
				}
				handled = true;
				break;
		}

		if (handled) {
			nativeEvent.preventDefault();
			if (nextRow !== row || nextCol !== col) {
				row = nextRow;
				col = nextCol;
				focusActiveCell();
			}
		}
	}

	function handleCellClick(event) {
		const { row: clickedRow, col: clickedCol } = event.detail;
		row = clickedRow;
		col = clickedCol;

		if (phase === 'setup') {
			placeShip(clickedRow, clickedCol);
		} else if (phase === 'playing' && viewMode === 'enemy') {
			shoot(clickedRow, clickedCol);
		}

		focusActiveCell();
	}

	function getShipCells(startRow, startCol, size, orientation) {
		const cells = [];
		for (let index = 0; index < size; index++) {
			if (orientation === 'horizontal') {
				cells.push([startRow, startCol + index]);
			} else {
				cells.push([startRow + index, startCol]);
			}
		}
		return cells;
	}

	// Injetar metadados de navios diretamente no myBoard evitando mapa duplicado
	function augmentBoardWithShips() {
		// Clonar raso para disparar reatividade só se mudar algo
		let changed = false;
		for (const ship of shipsToPlace) {
			if (!ship?.placed) continue;
			if (typeof ship.row !== 'number' || typeof ship.col !== 'number') continue;
			const orientation = ship.orientation === 'vertical' ? 'vertical' : 'horizontal';
			const cells = getShipCells(ship.row, ship.col, ship.size, orientation);
			for (const [r, c] of cells) {
				const cell = myBoard[r][c];
				// Não sobrescrever estados de combate caso já existam
				if (cell.state === 'hit' || cell.state === 'miss' || cell.state === 'sunk') continue;
				if (cell.shipId !== ship.id || cell.state !== 'ship' || cell.name !== ship.name) {
					myBoard[r][c] = { state: 'ship', shipId: ship.id, name: ship.name };
					changed = true;
				}
			}
		}
		if (changed) {
			// Forçar reatribuição para reatividade Svelte
			myBoard = myBoard.map(row => row.map(cell => ({ ...cell })));
		}
	}

	function canPlaceShip(startRow, startCol, size, orientation) {
		const cells = getShipCells(startRow, startCol, size, orientation);

		const isOutsideBoard = cells.some(([testRow, testCol]) =>
			testRow < 0 || testRow >= BOARD_SIZE || testCol < 0 || testCol >= BOARD_SIZE
		);
		if (isOutsideBoard) {
			return false;
		}

		const overlapsShip = cells.some(([testRow, testCol]) => myBoard[testRow][testCol].state !== 'empty');
		if (overlapsShip) {
			return false;
		}

		return true;
	}

 	async function placeShip(targetRow, targetCol) {
		if (placingShip) return; // bloqueia reentrância
		if (phase !== 'setup' || allShipsPlaced()) return;

		const ship = getCurrentShip();
		if (!ship || ship.placed) return;

		// Guardar índice e id antes de qualquer await para evitar condição de corrida
		const snapshotShipId = ship.id;
		const snapshotIndex = shipsToPlace.findIndex(s => s.id === snapshotShipId);

		if (!canPlaceShip(targetRow, targetCol, ship.size, shipOrientation)) {
			liveAnnounce = 'Não é possível posicionar o navio nesse local.';
			return;
		}

		placingShip = true;
		const coordTxt = getCoordinate(targetRow, targetCol);
		liveAnnounce = `Posicionando ${ship.name} em ${coordTxt}...`;

		let serverOk = false;
		try {
			const resp = await sendShipPlacement(ship.id, targetRow, targetCol, shipOrientation);
			// WebSocket retorna placeResult com state já aplicado via onmessage; considerar sucesso se ok
			serverOk = !!resp?.ok;
		} catch (error) {
			console.error('Falha ao sincronizar posicionamento de navio:', error);
			if (String(error?.message).includes('WebSocket não conectado')) {
				liveAnnounce = 'Reconectando... tente novamente em instantes.';
				if (!wsConnecting) connectWebSocket(false);
			} else {
				liveAnnounce = 'Erro ao sincronizar com o servidor. Tente novamente.';
			}
			placingShip = false;
			return;
		}

		if (!serverOk) {
			liveAnnounce = 'Servidor não confirmou posicionamento.';
			placingShip = false;
			return;
		}

		// Verificar se o broadcast já marcou este navio como placed (evita duplicação / marcação errada de outro índice)
		const existing = shipsToPlace.find(s => s.id === snapshotShipId);
		const alreadyPlacedByServer = existing?.placed;

		if (!alreadyPlacedByServer) {
			// Atualizar tabuleiro localmente apenas se servidor ainda não refletiu (modo defensivo)
			const occupiedCells = getShipCells(targetRow, targetCol, ship.size, shipOrientation);
			const occupiedSet = new Set(occupiedCells.map(([r, c]) => `${r}-${c}`));
			myBoard = myBoard.map((boardRow, rowIndex) =>
				boardRow.map((cell, colIndex) => (occupiedSet.has(`${rowIndex}-${colIndex}`)
					? { state: 'ship', shipId: snapshotShipId, name: ship.name }
					: cell))
			);

			shipsToPlace = shipsToPlace.map((current, index) =>
				index === snapshotIndex
					? {
						...current,
						placed: true,
						row: targetRow,
						col: targetCol,
						orientation: shipOrientation
					}
					: current
			);
		}

		// Atualiza índice atual (próximo navio não-posicionado)
		updateCurrentShipIndex();

		await tick();
		focusActiveCell();
		liveAnnounce = `${ship.name} posicionado em ${coordTxt}`;

		if (allShipsPlaced() && !announcedAllShips) {
			liveAnnounce = 'Todos os navios posicionados. Aguardando oponente.';
			announcedAllShips = true;
		}

		updateTurnView();

		placingShip = false;
	}

	function rotateShip() {
		shipOrientation = shipOrientation === 'horizontal' ? 'vertical' : 'horizontal';
		liveAnnounce = `Orientação do navio: ${shipOrientation === 'horizontal' ? 'horizontal' : 'vertical'}`;
	}

	function getCoordinate(targetRow, targetCol) {
		const letter = String.fromCharCode(65 + targetCol);
		return `${letter}${targetRow + 1}`;
	}

	async function sendWsCommand(type, payload, timeout=6000) {
		if (!ws || ws.readyState !== WebSocket.OPEN) {
			try { await whenWsOpen(); } catch { throw new Error('WebSocket não conectado'); }
		}
		return new Promise((resolve, reject) => {
			const id = String(++wsSeq);
			const timer = setTimeout(() => {
				if (pendingCmds.has(id)) {
					pendingCmds.delete(id);
					reject(new Error('timeout comando '+type));
				}
			}, timeout);
			pendingCmds.set(id, { resolve, reject, timer, type });
			try {
				ws.send(JSON.stringify({ ...payload, type, id }));
			} catch (e) {
				clearTimeout(timer);
				pendingCmds.delete(id);
				reject(e);
			}
		});
	}

	// Reset override quando encerramos turno por shot confirmado do oponente (miss nosso) será feito em applyRoundState.

	// ===== Sistema de anúncio enxuto =====
	let pendingWaitTimer = null;
	function scheduleWait(msg='Aguardando servidor...', delay=450) {
		clearTimeout(pendingWaitTimer);
		pendingWaitTimer = setTimeout(()=>{ liveAnnounce = msg; }, delay);
	}
	function cancelWait() { clearTimeout(pendingWaitTimer); pendingWaitTimer = null; }

    async function shoot(targetRow, targetCol) {
		const pendingKey = `${targetRow}-${targetCol}`;
		if (pendingShots.has(pendingKey)) {
			// Não repetir anúncio; mantemos silêncio para não poluir fila do leitor
			return;
		}
		if (!isMyTurn) {
			liveAnnounce = 'Não é seu turno.';
			return;
		}
		const currentState = enemyBoard[targetRow]?.[targetCol]?.state;
		if (currentState && currentState !== 'empty') {
			liveAnnounce = 'Célula já atacada anteriormente.';
			return;
		}

		const contextDebug = { phase, isMyTurn, viewMode, targetRow, targetCol, timestamp: new Date().toISOString() };
		console.debug('Solicitando tiro', contextDebug);
		lastLocalShot = { row: targetRow, col: targetCol };
		// Adiamos mensagem de espera; se resposta vier rápido, não há ruído.
		scheduleWait();
		pendingShots.add(pendingKey);
		try {
			if (wsStatus !== 'open') {
				cancelWait();
				liveAnnounce = 'Conexão perdida. Tentando reconectar...';
				return;
			}
			const result = await sendShot(targetRow, targetCol);
			if (!result || !result.ok) {
				cancelWait();
				liveAnnounce = 'Erro ao enviar tiro';
				return;
			}
			// applyRoundState já ocorre via onmessage (shootResult)
		} catch (error) {
			console.warn('Erro ao processar tiro (catch)', { error: error?.message, ...contextDebug });
		} finally {
			pendingShots.delete(pendingKey);
		}
	}

	function toggleView() {
		if (phase !== 'playing') return;
		const next = viewMode === 'my' ? 'enemy' : 'my';
		setViewMode(next, true);
	}

	function setViewMode(mode, manual=false) {
		if (viewMode === mode) return;
		viewMode = mode;
		if (manual) manualViewOverride = true; else manualViewOverride = false;
		liveAnnounce = `Visualizando ${viewMode === 'my' ? 'seu tabuleiro' : 'tabuleiro inimigo'}`;
	}

	// ===== Acessibilidade: abas (tablist) para alternância ataque/defesa =====
	let tabMyEl; let tabEnemyEl; let tablistEl; // refs
	function focusActiveTab() {
		const el = viewMode === 'my' ? tabMyEl : tabEnemyEl;
		if (el && typeof el.focus === 'function') el.focus();
	}
	function handleTabListKeydown(e) {
		if (e.key === 'ArrowLeft' || e.key === 'ArrowRight') {
			// Alterna mantendo padrão roving tabindex
			const next = viewMode === 'my' ? 'enemy' : 'my';
			setViewMode(next, true);
			// aguardar flush reativo mínimo
			setTimeout(focusActiveTab, 0);
			e.preventDefault();
		}
	}
	$: if (phase === 'playing') {
		// Ajustar roving tabindex quando modo mudar
		if (tabMyEl) tabMyEl.tabIndex = viewMode === 'my' ? 0 : -1;
		if (tabEnemyEl) tabEnemyEl.tabIndex = viewMode === 'enemy' ? 0 : -1;
	}

	function applyServerShipData(serverShips) {
		if (!Array.isArray(serverShips)) return;

		const serverMap = new Map(serverShips.map((ship) => [ship.id, ship]));
		shipsToPlace = shipsToPlace.map((ship) => {
			const serverShip = serverMap.get(ship.id);
			if (!serverShip) return ship;
			return {
				...ship,
				placed: !!serverShip.placed,
				row: serverShip.row ?? ship.row,
				col: serverShip.col ?? ship.col,
				orientation: serverShip.orientation ?? ship.orientation
			};
		});
		updateCurrentShipIndex();
	}

	function applyRoundState(data) {
		if (!data || typeof data !== 'object') return;

		const priorCurrentPlayer = currentPlayer;
		const priorPhase = phase;
		const priorIsMyTurn = isMyTurn;


		let phaseChanged = false;
		if (data.phase && phase !== data.phase) { phase = data.phase; phaseChanged = true; }
		if (data.currentPlayer) currentPlayer = data.currentPlayer;
		if (typeof data.playerAReady === 'boolean') playerAReady = data.playerAReady;
		if (typeof data.playerBReady === 'boolean') playerBReady = data.playerBReady;
		if (typeof data.winner === 'string') winner = data.winner;
		if (typeof data.scoreA === 'number') scoreA = data.scoreA;
		if (typeof data.scoreB === 'number') scoreB = data.scoreB;
		if (typeof data.scoreDraw === 'number') scoreDraw = data.scoreDraw;
		if (data.playerA !== undefined) hasPlayerA = !!data.playerA;
		if (data.playerB !== undefined) hasPlayerB = !!data.playerB;
		if (data.shipsRemaining) {
			shipsRemaining = {
				a: data.shipsRemaining.a ?? shipsRemaining.a,
				b: data.shipsRemaining.b ?? shipsRemaining.b
			};
		}
		if (typeof data.statusMessage === 'string') {
			statusMessage = data.statusMessage;
		}

		if (Array.isArray(data.myBoard)) {
			const merged = mergeBoard(myBoard, data.myBoard);
			if (merged !== myBoard) {
				myBoard = merged;
			}
		}
		if (Array.isArray(data.enemyBoard)) {
			const merged = mergeBoard(enemyBoard, data.enemyBoard);
			if (merged !== enemyBoard) {
				enemyBoard = merged;
			}
		}
		if (Array.isArray(data.myShips)) {
			applyServerShipData(data.myShips);
		}
		if (data.lastShot && typeof data.lastShot === 'object') {
			const { row: lastRow, col: lastCol, result } = data.lastShot;
			if (typeof lastRow === 'number' && typeof lastCol === 'number' && typeof result === 'string') {
				const shotSig = `${lastRow}-${lastCol}-${result}`;
				lastAction = `${result} em ${getCoordinate(lastRow, lastCol)}`;
				if (shotSig !== prevLastShotSig) {
					// Cancelar mensagem de espera se era confirmação de tiro local
					cancelWait();
					if (result === 'sunk' && data.lastSunkShip && Array.isArray(data.lastSunkShip.cells)) {
						lastSunkShip = data.lastSunkShip;
						liveAnnounce = `Afundou navio (${data.lastSunkShip.size || data.lastSunkShip.cells.length}).`; 
					} else if (result === 'hit') {
						liveAnnounce = `Acerto ${getCoordinate(lastRow, lastCol)}`;
					} else if (result === 'miss') {
						liveAnnounce = `Água ${getCoordinate(lastRow, lastCol)}`;
					} else if (result === 'sunk') {
						liveAnnounce = `Afundou em ${getCoordinate(lastRow, lastCol)}`;
					}
					prevLastShotSig = shotSig;
				}
				// Se perdemos o turno (miss) mudar automaticamente para defesa
				if (result === 'miss' && currentPlayer !== localPlayer) {
					setViewMode('my', false); // perda de turno -> visão defesa
				}

				// Se este broadcast confirmar o último tiro local, limpar lastLocalShot
				if (lastLocalShot && lastLocalShot.row === lastRow && lastLocalShot.col === lastCol) {
					lastLocalShot = null;
					manualViewOverride = false; // confirmação encerra qualquer override ligado ao tiro
				}

				// Fallback: se resultado sunk mas servidor não marcou todo navio, tentar completar cluster local
				// Sem fallback local: dependemos do servidor enviar todas as células sunk via board/lastSunkShip
			}
		}
		if (data.lastSunkShip && !data.lastShot) {
			// Caso raro: sunk ship broadcast sem lastShot
			lastSunkShip = data.lastSunkShip;
			if (Array.isArray(data.lastSunkShip.cells)) {
				liveAnnounce = `Navio afundado (${data.lastSunkShip.size || data.lastSunkShip.cells.length} células).`;
			}
		}

		// Sem fallback de mudança de fase: aguardamos broadcast oficial do servidor

		// Força recalcular turno antes de ajustar visão para evitar defasagem em disabled
		isMyTurn = isMyTurnCheck();
		const gainedTurn = !priorIsMyTurn && isMyTurn;

		// Ajustar visão após qualquer atualização significativa de turno/fase
		updateTurnView();

		// Se acabamos de entrar em 'playing', definir visão inicial adequada
		if (phaseChanged && phase === 'playing') {
			if (currentPlayer === localPlayer) {
				// Jogador local começa atacando (auto)
				setViewMode('enemy', false);
				liveAnnounce = 'Fase de combate iniciada. Seu turno, selecione uma célula para atirar.';
			} else {
				setViewMode('my', false);
				liveAnnounce = 'Fase de combate iniciada. Aguarde o turno do oponente.';
			}
		}

		// Auto switch ao ganhar turno (desde que usuário não tenha feito override manual desde última troca)
		if (!phaseChanged && phase === 'playing' && gainedTurn) {
			if (!manualViewOverride) {
				setViewMode('enemy', false);
				// Mensagem de turno só se não houve lastShot do oponente na mesma atualização
				if (!data.lastShot || (data.lastShot && data.lastShot.result === 'miss')) {
					liveAnnounce = 'Seu turno. Selecione uma célula para atirar.';
				}
			}
			// Reset override para próxima janela de decisão
			manualViewOverride = false;
		}

		// Foco preservado por referência estável; refoco explícito não mais necessário.

		// (Fallback de fase removido para expor problema real de broadcast)
	}

/* Fallbacks removidos: toda responsabilidade de marcar navio afundado completo e iniciar fase playing é do backend */

	async function loadRound() {
		try {
			const response = await fetch(`/api/battleship/${roundParam}`);
			if (!response.ok) return;
			const data = await response.json();
			applyRoundState(data);
		} catch (error) {
			console.error('Erro ao carregar rodada', error);
		}
	}

	async function loadGameData() {
		try {
			const [myBoardRes, enemyBoardRes, shipsRes, roundMetaRes] = await Promise.all([
				fetch(`/api/battleship/${roundParam}/${playerParam}/board`),
				fetch(`/api/battleship/${roundParam}/${playerParam}/board?enemy=true`),
				fetch(`/api/battleship/${roundParam}/${playerParam}/ships`),
				fetch(`/api/battleship/${roundParam}`)
			]);

			if (myBoardRes.ok) {
				const boardData = await myBoardRes.json();
				if (boardData.board) {
					myBoard = normalizeBoard(boardData.board);
					boardInitialized = true;
				}
			}

			if (enemyBoardRes.ok) {
				const enemyData = await enemyBoardRes.json();
				if (enemyData.board) {
					enemyBoard = normalizeBoard(enemyData.board);
				}
			}

			if (shipsRes.ok) {
				const shipsData = await shipsRes.json();
				if (Array.isArray(shipsData.ships)) {
					applyServerShipData(shipsData.ships);
				}
			}
			if (roundMetaRes.ok) {
				try {
					const roundMeta = await roundMetaRes.json();
					applyRoundState(roundMeta);
				} catch(e) { console.warn('Falha meta round', e); }
			}
		} catch (error) {
			console.error('Erro ao carregar dados do jogo', error);
		}
	}

	async function ensureRound() {
		try {
			const response = await fetch(`/api/battleship/${roundParam}`);
			if (response.status === 404) {
				const createResponse = await fetch(`/api/battleship/${roundParam}/new`);
				if (createResponse.ok) {
					const data = await createResponse.json();
					if (data?.round && Number(data.round) !== Number(roundParam)) {
						redirecting = true;
						window.location.href = `/battleship/${data.round}/${playerParam}`;
						return;
					}
					await loadRound();
				}
			} else if (response.ok) {
				await loadRound();
			}
		} catch (error) {
			console.error('Erro garantindo partida', error);
		}
	}

 	async function sendShipPlacement(shipId, row, col, orientation) {
		return sendWsCommand('placeShip', { shipId, row, col, orientation });
	}

	async function sendShot(row, col) {
		return sendWsCommand('shoot', { row, col });
	}

	function reset() {
		window.location.reload();
	}

	async function newRound() {
		if (redirecting || !winner) return;
		try {
			const res = await fetch(`/api/battleship/${roundParam}/new`, { method: 'GET', cache: 'no-store' });
			if (res.ok) {
				const data = await res.json();
				if (data?.round && Number(data.round) !== Number(roundParam)) {
					redirecting = true;
					window.location.href = `/battleship/${data.round}/${playerParam}`;
				}
			}
		} catch(e) {
			console.error('Erro nova rodada', e);
		}
	}

	// (Removido: bloco duplicado de onMount com WebSocket manual - unificado na versão superior)
</script>

<PageTitle title="Batalha Naval" game={`Partida ${roundParam}`} />

<!-- Proteção contra anúncios intrusivos -->
<AdProtection protection="strict" />

<div class="container battleship-container py-4 d-flex flex-column align-items-center">
	<!-- Anúncio topo (desktop/mobile permitido conforme config) -->
	<AdManager
		placement="top"
		format="banner"
		adSlot={adConfig.slots.topBanner}
		priority="low"
	/>
	<div class="scoreboard-container mb-3">
		<BattleshipScoreboard
			{scoreA}
			{scoreB}
			{scoreDraw}
			{currentPlayer}
			{phase}
			{winner}
		/>
	</div>

	<div class="status-container mb-3">
		<BattleshipStatus
			{phase}
			{currentPlayer}
			{localPlayer}
			{isMyTurn}
			message={statusMessage}
			{shipsRemaining}
			{playerAReady}
			{playerBReady}
			{lastAction}
		/>
	</div>

	{#if phase === 'setup' && !allShipsPlaced()}
		<div class="setup-info mb-3">
			<div class="current-ship">
				<strong>Posicionando:</strong> {getCurrentShip()?.name}
				<span>(tamanho: {getCurrentShip()?.size})</span>
			</div>
			<div class="setup-controls">
				<span>Orientação: <strong>{shipOrientation === 'horizontal' ? 'Horizontal' : 'Vertical'}</strong></span>
				<span class="text-muted">• R: rotacionar • S: status • H: ajuda</span>
			</div>
			<div class="ships-progress">
				<small class="text-muted">
					Navios posicionados: {shipsToPlace.filter((ship) => ship.placed).length} de {shipsToPlace.length}
				</small>
			</div>
		</div>
	{/if}

	{#if phase === 'playing'}
		<div class="view-controls mb-3" role="tablist" aria-label="Tabuleiros" bind:this={tablistEl} on:keydown={handleTabListKeydown} tabindex="0">
			<button
				id="tab-my"
				bind:this={tabMyEl}
				role="tab"
				class={`btn ${viewMode === 'my' ? 'btn-primary' : 'btn-outline-primary'} tab-btn`}
				aria-selected={viewMode === 'my'}
				aria-controls="panel-board"
				tabindex={viewMode === 'my' ? 0 : -1}
				type="button"
				aria-label={`Defesa ${myRemaining}`}
				title={`Meus navios restantes: ${myRemaining}`}
				on:click={() => { setViewMode('my', true); focusActiveTab(); }}
			>
				<span aria-hidden="true">🚢 Defesa <span class="badge">{myRemaining}</span></span>
			</button>
			<button
				id="tab-enemy"
				bind:this={tabEnemyEl}
				role="tab"
				class={`btn ${viewMode === 'enemy' ? 'btn-primary' : 'btn-outline-primary'} tab-btn`}
				aria-selected={viewMode === 'enemy'}
				aria-controls="panel-board"
				tabindex={viewMode === 'enemy' ? 0 : -1}
				type="button"
				aria-label={`Ataque ${enemyRemaining}`}
				title={`Navios inimigos restantes: ${enemyRemaining}`}
				on:click={() => { setViewMode('enemy', true); focusActiveTab(); }}
			>
				<span aria-hidden="true">🎯 Ataque <span class="badge">{enemyRemaining}</span></span>
			</button>
			<span class="view-hint text-muted">Setas: alternar • T: alternar • H: ajuda</span>
		</div>
	{/if}

	<div class="board-container" role="tabpanel" id="panel-board" aria-labelledby={viewMode === 'my' ? 'tab-my' : 'tab-enemy'}>
		<BattleshipBoard
			board={currentBoard}
			mode={boardMode}
			{isMyTurn}
			{row}
			{col}
			{roundParam}
			{shipNames}
			pendingShots={pendingShots}
			setupShip={phase === 'setup' ? getCurrentShip() : null}
			setupOrientation={shipOrientation}
			on:cellClick={handleCellClick}
			on:cellFocus={handleCellFocus}
			on:keydown={handleKeydown}
		/>
	</div>

	<!-- Anúncio retângulo (após tabuleiro) -->
	<AdManager
		placement="auto"
		format="rectangle"
		adSlot={adConfig.slots.rectangle}
		priority="medium"
	/>

	<div class="actions mt-4">
		{#if winner}
			<button class="btn btn-primary btn-lg" on:click={newRound} disabled={redirecting || !winner} aria-disabled={!winner} title={!winner ? 'Aguarde terminar' : 'Criar nova rodada baseada no placar atual'}>
				Nova rodada
			</button>
		{:else if phase === 'setup' && allShipsPlaced()}
			<button class="btn btn-success btn-lg" disabled>
				Aguardando oponente...
			</button>
		{/if}
	</div>

	<!-- Anúncio inferior -->
	<AdManager
		placement="bottom"
		format="banner"
		adSlot={adConfig.slots.bottomBanner}
		priority="low"
	/>

	<div class="sr-only" aria-live="polite">{liveAnnounce}</div>
</div>

{#if debugMode}
<div class="bs-debug-overlay">
 	<strong>DEBUG</strong>
 	<div>phase: {phase}</div>
 	<div>currentPlayer: {currentPlayer}</div>
 	<div>localPlayer: {localPlayer}</div>
 	<div>playerAReady: {playerAReady ? 'Y' : 'N'} | playerBReady: {playerBReady ? 'Y' : 'N'}</div>
 	<div>isMyTurn: {isMyTurn ? 'Y' : 'N'}</div>
 	<div>viewMode: {viewMode}</div>
 	<div>boardMode: {boardMode}</div>
 	<div>placingShip: {placingShip ? 'Y' : 'N'}</div>
 	<div>ships placed: {shipsToPlace.filter(s=>s.placed).length}/{shipsToPlace.length}</div>
 	<div>lastAction: {lastAction}</div>
</div>
{/if}

<style>
	.battleship-container {
		max-width: 800px;
		width: 100%;
	}

	.scoreboard-container,
	.status-container {
		width: 100%;
		max-width: 600px;
	}

	.setup-info {
		text-align: center;
		background: rgba(40, 167, 69, 0.1);
		border: 1px solid rgba(40, 167, 69, 0.3);
		border-radius: 0.5rem;
		padding: 1rem;
		color: #e8f4fd;
	}

	.current-ship {
		font-size: 1.1rem;
		margin-bottom: 0.5rem;
		display: flex;
		gap: 0.5rem;
		justify-content: center;
		flex-wrap: wrap;
	}

	.setup-controls {
		font-size: 0.9rem;
		display: flex;
		gap: 1rem;
		justify-content: center;
		align-items: center;
		flex-wrap: wrap;
	}

	.ships-progress {
		margin-top: 0.5rem;
		text-align: center;
	}

	.view-controls {
		display: flex;
		gap: 1rem;
		align-items: center;
		flex-wrap: wrap;
		justify-content: center;
	}

	.view-hint {
		font-size: 0.85rem;
		margin-left: 1rem;
	}

	.board-container {
		width: 100%;
		max-width: 600px;
	}

	.actions {
		text-align: center;
	}

	.sr-only {
		position: absolute;
		width: 1px;
		height: 1px;
		padding: 0;
		margin: -1px;
		overflow: hidden;
		clip: rect(0 0 0 0);
		white-space: nowrap;
		border: 0;
	}

	@media (max-width: 768px) {
		.battleship-container {
			padding-left: 1rem;
			padding-right: 1rem;
		}

		.setup-controls {
			flex-direction: column;
			gap: 0.5rem;
		}

		.view-controls {
			gap: 0.75rem;
		}

		.view-hint {
			margin-left: 0;
			margin-top: 0.5rem;
		}
	}

	/* Debug overlay */
	.bs-debug-overlay {
		position: fixed;
		top: .5rem;
		right: .5rem;
		background: rgba(0,0,0,0.8);
		color: #0f0;
		font: 12px/1.3 monospace;
		padding: .6rem .8rem;
		border: 1px solid #0f0;
		border-radius: .4rem;
		z-index: 9999;
		max-width: 260px;
		pointer-events: none;
	}
	.bs-debug-overlay strong { display:block; margin-bottom:.3rem; color:#fff; }
</style>

<!-- estilos consolidados: debug overlay movido para o bloco principal -->