Resumo
Implementa o jogo Batalha Naval completo na plataforma:
- Frontend Svelte: layout dedicado, página inicial de partidas, seleção de jogador (com QR e compartilhamento), página de jogo com placar, status, tabuleiros Defesa/Ataque e botão Nova rodada.
- Backend Go: rotas REST (estado, nova rodada, QR, tabuleiro, navios, partidas abertas) + WebSocket para comandos (placeShip, shoot, ping) e broadcast de estado.
- Acessibilidade: rótulos de célula estado/nome antes da coordenada, regiões aria-live separadas, tablist acessível, anúncios enxutos (Acerto / Água / Afundado / turno).
- Responsividade: ajustes de layout e board para desktop/tablet/mobile.
- Compartilhamento: QR code, copiar link, suporte ao Web Share API em dispositivos móveis.
- Continuidade de partidas: botão “Nova rodada” preservando placar acumulado.
- Ads: slots top / rectangle / bottom com proteção de áreas interativas.

Detalhes técnicos
Frontend (principais arquivos):
- src/routes/battleship/+layout.svelte
- src/routes/battleship/+page.svelte
- src/routes/battleship/[round]/+page.svelte
- src/routes/battleship/[round]/[player]/+page.svelte
- src/lib/battleship/BattleshipBoard.svelte
- src/lib/battleship/BattleshipStatus.svelte
- src/lib/battleship/BattleshipScoreboard.svelte
- src/lib/battleship/labels.ts

Backend:
- pkg/battleship/game.go (motor de tabuleiro e tiros)
- pkg/battleship/round.go (estado da rodada, broadcast, turnos)
- pkg/battleship/battleship.go (registro de rotas REST + WebSocket)

Testes
- Frontend Vitest: BattleshipBoard.test.ts, labels.test.ts, __tests__/broadcast.test.ts (todos passando).
- Backend Go: game_test.go, round_test.go (cobrem posicionamento, tiros, afundamento, LastSunkShip, máscara de tabuleiro, vitória, readiness).

Acessibilidade
- Labels únicos por célula (estado → coord), evitando redundância.
- Anúncios de ação separados de status de fase/turno para reduzir sobrecarga auditiva.
- Tablist com roving tabindex e alternância por setas ou tecla T.
- QR e feedbacks de copiar/compartilhar com aria-live="polite".

WebSocket
- Mensagens tipadas (type: state, shootResult, placeResult, error, pong).
- Correlação de comandos via id para respostas assíncronas.
- Máscara de tabuleiro inimigo mostrando somente células reveladas (hit/miss/sunk).
- Registro completo de navio afundado (LastSunkShip) com coordenadas para acessibilidade.

Placar e continuidade
- Manutenção de ScoreA, ScoreB, ScoreDraw entre rodadas.
- Rota /battleship/{round}/new gera próxima rodada herdando placar.
- Botão “Nova rodada” no frontend se winner definido.

Ads
- Uso de AdManager e AdProtection com slots configurados (top, rectangle, bottom) e bloqueio de sobreposição em áreas sensíveis do jogo.

Pendências futuras (não inclusas neste PR)
- Extração de tokens CSS responsivos comuns.
- Componente reutilizável de InvitePanel (QR + share).
- Internacionalização (labels e mensagens).
- Testes adicionais de reconexão e espectadores simultâneos.

Checklist
- [x] Acessibilidade base (rótulos, aria-live, tablist)
- [x] Responsividade
- [x] Backend WS + rotas REST
- [x] Compartilhamento (QR / copiar / Web Share)
- [x] Continuidade de placar
- [x] Ads integrados
- [x] Testes frontend/back-end
- [ ] Tokens CSS (próximo PR)
- [ ] InvitePanel reutilizável
- [ ] i18n

Validação manual
1. Abrir /battleship → criar nova partida.
2. Selecionar jogador A ou B.
3. Posicionar navios (R para rotacionar, S status).
4. Início automático quando ambos prontos.
5. Atacar alternando para aba Ataque (T ou setas).
6. Após vitória, usar “Nova rodada” para continuar placar.
7. Testar QR e copiar link em página de seleção.

Riscos
- Volume grande de arquivos novos (revisão detalhada recomendada).
- Ajustes futuros de logging de broadcast.

Solicitação
Revisão de acessibilidade, segurança WS, performance de mergeBoard e layout responsivo.

Último commit
18534434ba8f27f831244a942e2f03cd98e8e44d
