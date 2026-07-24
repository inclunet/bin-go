<script>
    import { onDestroy, onMount } from "svelte";
    import CardHeader from "$lib/bingo/CardHeader.svelte";
    import BingoCelebrationModal from "$lib/bingo/BingoCelebrationModal.svelte";
    import Card from "$lib/bingo/Card.svelte";
    import {
        getPendingLobbySourceCardID,
        getPendingLobbyURL,
        rememberPendingLobbyURL,
    } from "$lib/bingo/playerCards";
    import { card } from "../../../../lib/bingo";
    import PageTitle from "$lib/PageTitle.svelte";
    import Play from "$lib/Play.svelte";
    import { callApiResult, getWSEndpoint } from "$lib/api";
    import Completions from "$lib/bingo/Completions.svelte";
    import TableModalCard from "$lib/TableModalCard.svelte";
    import MediaQuery from "$lib/MediaQuery.svelte";
    import Button from "$lib/Button.svelte";

    export let data;

    /**
     * @type {HTMLAudioElement}
     */
    let bingoAudio;

    /**
     * @type {HTMLAudioElement}
     */
    let checkAudio;

    let config = false;
    /** @type {WebSocket | undefined} */
    let socket;
    /** @type {ReturnType<typeof setInterval> | undefined} */
    let pollingInterval;
    /** @type {ReturnType<typeof setTimeout> | undefined} */
    let reconnectTimeout;
    /** @type {ReturnType<typeof setTimeout> | undefined} */
    let connectionTimeout;
    /** @type {ReturnType<typeof setTimeout> | undefined} */
    let navigationRecoveryTimeout;
    let pollingInFlight = false;
    let leavingPage = false;
    let navigationPending = false;
    let cardLoaded = false;
    let loadError = "";
    let actionError = "";
    let soundStatus = "idle";
    let bingoDismissError = "";
    let dismissingBingo = false;
    let bingoSoundSilenced = false;
    let soundGeneration = 0;
    let syncGeneration = 0;
    let updateGeneration = 0;
    let foregroundUpdates = 0;
    /** @type {WakeLockSentinel | undefined} */
    let screenWakeLock;
    let wakeLockNeedsAction = false;
    let requestingWakeLock = false;

    const isValidCard = (value) =>
        value.ID &&
        value.RoundID &&
        value.Round > 0 &&
        value.Card > 0;

    const requestScreenWakeLock = async () => {
        if (screenWakeLock?.released) {
            screenWakeLock = undefined;
        }
        if (
            leavingPage ||
            requestingWakeLock ||
            screenWakeLock ||
            document.visibilityState !== "visible" ||
            !("wakeLock" in navigator)
        ) {
            return;
        }

        requestingWakeLock = true;
        try {
            const wakeLock = await navigator.wakeLock.request("screen");
            if (leavingPage) {
                await wakeLock.release();
                return;
            }
            screenWakeLock = wakeLock;
            wakeLockNeedsAction = false;
            wakeLock.addEventListener(
                "release",
                () => {
                    if (screenWakeLock === wakeLock) {
                        screenWakeLock = undefined;
                        if (
                            !leavingPage &&
                            document.visibilityState === "visible"
                        ) {
                            wakeLockNeedsAction = true;
                        }
                    }
                },
                { once: true }
            );
            if (wakeLock.released) {
                screenWakeLock = undefined;
                wakeLockNeedsAction =
                    document.visibilityState === "visible";
            }
        } catch {
            wakeLockNeedsAction = true;
        } finally {
            requestingWakeLock = false;
        }
    };

    const handleVisibilityChange = () => {
        if (document.visibilityState === "visible") {
            requestScreenWakeLock();
        }
    };

    const startPolling = () => {
        if (!leavingPage && !navigationPending && !pollingInterval) {
            poolingUpdater();
            pollingInterval = setInterval(poolingUpdater, 1000);
        }
    };

    const stopPolling = () => {
        if (pollingInterval) {
            clearInterval(pollingInterval);
            pollingInterval = undefined;
        }
    };

    const scheduleReconnect = () => {
        if (
            leavingPage ||
            navigationPending ||
            reconnectTimeout ||
            !("WebSocket" in window)
        ) {
            return;
        }
        reconnectTimeout = setTimeout(() => {
            reconnectTimeout = undefined;
            liveUpdater();
        }, 3000);
    };

    const handleAutoplayEvent = async () => {
        if (
            !(await updateCard(
                `/api/bingo/${$card.RoundID}/${$card.ID}/autoplay`
            ))
        ) {
            return;
        }
    };

    const handleCheckNumberEvent = async (event = {}) => {
        if (!$card.Autoplay || $card.Card == 1) {
            if (
                !(await updateCard(
                    `/api/bingo/${$card.RoundID}/${$card.ID}/${event.detail.Number}`
                ))
            ) {
                return;
            }
        }
    };

    const handleOpenConfigEvent = () => {
        config = true;
        setTimeout(() => {
            const dialog = document.querySelector("[role=dialog] h1");
            if (dialog) {
                dialog.focus();
            }
        }, 100);
    };

    const handleDrawEvent = async () => {
        if (
            !(await updateCard(
                `/api/bingo/${$card.RoundID}/${$card.ID}/0`
            ))
        ) {
            return;
        }
    };

    const navigateTo = (target) => {
        if (leavingPage || navigationPending) {
            return;
        }
        navigationPending = true;
        stopPolling();
        try {
            window.location.assign(target);
        } catch {
            navigationPending = false;
            actionError =
                "Não foi possível abrir a próxima página. Tente novamente.";
            return;
        }
        navigationRecoveryTimeout = setTimeout(() => {
            // A successful navigation destroys this component. If the browser
            // cancels it, keep the current card connected and retryable.
            navigationRecoveryTimeout = undefined;
            navigationPending = false;
            requestScreenWakeLock();
            if (!cardLoaded) {
                loadCard();
            } else if ("WebSocket" in window) {
                liveUpdater();
            } else {
                startPolling();
            }
        }, 1000);
    };

    const handleNewRoundEvent = async () => {
        const pendingLobbyURL = getPendingLobbyURL();
        const pendingSourceCardID = getPendingLobbySourceCardID();
        if (
            pendingLobbyURL &&
            pendingSourceCardID === String(data.Card)
        ) {
            navigateTo(pendingLobbyURL);
            return;
        }

        const sourceCardID = $card.ID;
        const result = await callApiResult(
            $card,
            `/api/bingo/${$card.RoundID}/${$card.ID}/new/${$card.Type}`,
            "GET"
        );
        if (
            !result.ok ||
            !isValidCard(result.data) ||
            result.data.Card !== 1
        ) {
            actionError =
                "Não foi possível criar a próxima rodada. Tente novamente.";
            return;
        }
        actionError = "";
        const lobbyURL = `/bingo/${result.data.RoundID}/${result.data.ID}/lobby`;
        rememberPendingLobbyURL(lobbyURL, sourceCardID);
        navigateTo(lobbyURL);
    };

    const handlePlayCheckSoundEvent = async () => {
        if (!checkAudio) {
            return;
        }
        checkAudio.pause();
        checkAudio.currentTime = 0;
        checkAudio.play();
    };

    const handleCancelBingoAlertEvent = async () => {
        if (dismissingBingo) {
            return;
        }
        bingoSoundSilenced = true;
        soundGeneration++;
        if (bingoAudio) {
            bingoAudio.pause();
            bingoAudio.currentTime = 0;
        }
        soundStatus = "stopped";
        bingoDismissError = "";
        dismissingBingo = true;
        foregroundUpdates++;
        const requestUpdateGeneration = ++updateGeneration;
        const result = await callApiResult(
            $card,
            `/api/bingo/${$card.RoundID}/${$card.ID}/cancel`,
            "GET"
        );
        foregroundUpdates--;
        if (!result.ok || !isValidCard(result.data)) {
            actionError =
                "Não foi possível atualizar a cartela. Tente novamente.";
            bingoDismissError =
                "Não foi possível fechar o aviso. O som foi interrompido; tente novamente.";
            dismissingBingo = false;
            return;
        }
        actionError = "";
        bingoDismissError = "";
        dismissingBingo = false;
        const hasNewerCardState =
            requestUpdateGeneration !== updateGeneration;
        syncGeneration++;
        updateGeneration++;
        if (hasNewerCardState) {
            $card = {
                ...$card,
                Bingo: result.data.Bingo,
                LastCompletion: result.data.LastCompletion,
            };
        } else {
            $card = result.data;
        }
        isBingo();
        const previousSocket = socket;
        socket = undefined;
        previousSocket?.close();
        startPolling();
        scheduleReconnect();
    };

    const handleSaveCompletions = async () => {
        foregroundUpdates++;
        const requestUpdateGeneration = ++updateGeneration;
        const result = await callApiResult(
            $card,
            `/api/bingo/${$card.RoundID}/${$card.ID}/completions`,
            "POST",
            $card.Completions
        );
        foregroundUpdates--;
        if (!result.ok || !isValidCard(result.data)) {
            actionError =
                "Não foi possível salvar as configurações. Tente novamente.";
            return;
        }
        actionError = "";
        if (requestUpdateGeneration !== updateGeneration) {
            return;
        }
        $card = result.data;
    };

    const playBingoSound = async () => {
        const playbackGeneration = ++soundGeneration;
        bingoSoundSilenced = false;
        soundStatus = "starting";
        if (!bingoAudio) {
            soundStatus = "blocked";
            return;
        }
        try {
            await bingoAudio.play();
            if (
                playbackGeneration !== soundGeneration ||
                bingoSoundSilenced ||
                !$card.Bingo
            ) {
                bingoAudio.pause();
                bingoAudio.currentTime = 0;
                return;
            }
            soundStatus = "playing";
        } catch {
            if (playbackGeneration === soundGeneration) {
                soundStatus = "blocked";
            }
        }
    };

    const isBingo = () => {
        if ($card.Card > 1) {
            if (
                $card.Bingo &&
                !bingoSoundSilenced &&
                soundStatus === "idle"
            ) {
                void playBingoSound();
            } else if (!$card.Bingo) {
                bingoSoundSilenced = false;
                soundGeneration++;
                soundStatus = "idle";
                if (bingoAudio) {
                    bingoAudio.pause();
                    bingoAudio.currentTime = 0;
                }
            }
        }
    };

    const liveUpdater = () => {
        if (
            leavingPage ||
            navigationPending ||
            socket?.readyState === WebSocket.OPEN ||
            socket?.readyState === WebSocket.CONNECTING
        ) {
            return;
        }

        startPolling();
        const connection = new window.WebSocket(
            getWSEndpoint(`/ws/bingo/${$card.RoundID}/${$card.ID}`)
        );
        const connectionGeneration = syncGeneration;
        socket = connection;
        connectionTimeout = setTimeout(() => {
            if (
                socket === connection &&
                connection.readyState === WebSocket.CONNECTING
            ) {
                connection.close();
            }
        }, 10000);

        connection.addEventListener("open", () => {
            if (socket !== connection || leavingPage) {
                connection.close();
                return;
            }
            if (connectionTimeout) {
                clearTimeout(connectionTimeout);
                connectionTimeout = undefined;
            }
            actionError = "";
        });

        connection.addEventListener("message", (event) => {
            if (
                socket !== connection ||
                leavingPage ||
                connectionGeneration !== syncGeneration
            ) {
                return;
            }
            try {
                const updated = JSON.parse(event.data);
                if (!isValidCard(updated)) {
                    throw new Error("invalid card update");
                }
                updateGeneration++;
                $card = updated;
                stopPolling();
                actionError = "";
                loadError = "";
                redirectToNextRound();
                isBingo();
            } catch (error) {
                actionError =
                    "A atualização em tempo real foi interrompida. Tentando reconectar.";
                connection.close();
                startPolling();
            }
        });

        connection.addEventListener("close", () => {
            if (connectionTimeout) {
                clearTimeout(connectionTimeout);
                connectionTimeout = undefined;
            }
            if (socket !== connection || leavingPage) {
                return;
            }
            socket = undefined;
            startPolling();
            scheduleReconnect();
        });

        connection.addEventListener("error", () => {
            connection.close();
            startPolling();
        });
    };

    const loadCard = async () => {
        $card.ID = String(data.Card);
        $card.RoundID = String(data.Round);
        const loaded = await updateCard(
            `/api/bingo/${$card.RoundID}/${$card.ID}`
        );
        if (!loaded) {
            actionError = "";
            cardLoaded = false;
            loadError =
                "Não foi possível carregar esta cartela. Verifique o link e tente novamente.";
            return;
        }
        cardLoaded = true;
        loadError = "";

        if ("WebSocket" in window) {
            liveUpdater();
        } else {
            startPolling();
        }
    };

    const poolingUpdater = async () => {
        if (pollingInFlight || foregroundUpdates > 0) {
            return;
        }
        pollingInFlight = true;
        try {
            await updateCard(
                `/api/bingo/${$card.RoundID}/${$card.ID}`,
                false,
                true
            );
        } finally {
            pollingInFlight = false;
        }
    };

    const redirectToNextRound = () => {
        if (leavingPage || navigationPending) {
            return;
        }

        if ($card.Round == 0) {
            navigateTo("/");
            return;
        }

        if ($card.Card == 0) {
            navigateTo(`/bingo/${$card.RoundID}/new`);
            return;
        }

        if ($card.NextRoundID && $card.Card > 1) {
            navigateTo(`/bingo/${$card.NextRoundID}/new`);
        }
    };

    const updateCard = async (
        path = "",
        reportError = true,
        background = false
    ) => {
        const requestGeneration = syncGeneration;
        if (!background) {
            foregroundUpdates++;
            updateGeneration++;
        }
        const requestUpdateGeneration = updateGeneration;
        const result = await callApiResult($card, path, "GET");
        if (!background) {
            foregroundUpdates--;
        }
        if (
            requestGeneration !== syncGeneration ||
            requestUpdateGeneration !== updateGeneration
        ) {
            return true;
        }
        if (!result.ok) {
            if (reportError) {
                actionError =
                    "A ação não pôde ser concluída. Verifique sua conexão e tente novamente.";
            }
            return false;
        }
        const updated = result.data;
        if (!isValidCard(updated)) {
            if (reportError) {
                actionError =
                    "A API retornou uma cartela inválida. Atualize a página e tente novamente.";
            }
            return false;
        }
        actionError = "";
        $card = updated;
        redirectToNextRound();
        isBingo();
        return true;
    };

    export let table_client = false;

    $: table_client = $card.Card > 1 ? true : false;

    export let table_draw = false;

    $: table_draw = $card.Card == 1 ? true : false;

    onMount(() => {
        document.addEventListener("visibilitychange", handleVisibilityChange);
        const pendingLobbyURL = getPendingLobbyURL();
        const pendingSourceCardID = getPendingLobbySourceCardID();
        if (
            pendingLobbyURL &&
            pendingSourceCardID &&
            pendingSourceCardID === String(data.Card)
        ) {
            navigateTo(pendingLobbyURL);
            return;
        }
        requestScreenWakeLock();
        loadCard();
    });
    onDestroy(() => {
        leavingPage = true;
        document.removeEventListener(
            "visibilitychange",
            handleVisibilityChange
        );
        screenWakeLock?.release();
        screenWakeLock = undefined;
        socket?.close();
        stopPolling();
        if (reconnectTimeout) {
            clearTimeout(reconnectTimeout);
        }
        if (connectionTimeout) {
            clearTimeout(connectionTimeout);
        }
        if (navigationRecoveryTimeout) {
            clearTimeout(navigationRecoveryTimeout);
        }
    });
</script>

{#if loadError}
    <PageTitle title="Inclubingo - Cartela indisponível" game="Inclubingo" />
    <p class="alert alert-danger text-center" role="alert">{loadError}</p>
{:else if cardLoaded}
    <PageTitle
        title="Inclubingo - Cartela {$card.Card}, sorteio {$card.Round}"
        game="Inclubingo"
    />
    {#if actionError}
        <p class="alert alert-danger text-center" role="alert">{actionError}</p>
    {/if}
    {#if wakeLockNeedsAction}
        <div class="wake-lock-notice" aria-live="polite">
            <span>A tela pode apagar durante a partida.</span>
            <button
                type="button"
                class="btn btn-sm btn-outline-light"
                disabled={requestingWakeLock}
                on:click={requestScreenWakeLock}
            >
                {requestingWakeLock ? "Ativando…" : "Manter tela ligada"}
            </button>
        </div>
    {/if}
    <div class="container container-card">
    {#if $card.Card == 1}
        <MediaQuery query="(min-width: 1150px)" let:matches>
            {#if matches}
                <div class=" info-card table-horizontal" class:table_draw>
                    <div class="container-qr_code">
                        <img src="/qr/bingo/{$card.RoundID}" alt="QR-Code" />
                    </div>
                    <div class="info-card-header">
                        <h2>Cartela de Bingo #{$card.Card}</h2>
                        <h3 class="info-card-header-round">
                            Sorteio #{$card.Round}
                        </h3>
                    </div>
                    <div class="cardHeader">
                        <CardHeader
                            on:autoplay={handleAutoplayEvent}
                            on:openConfig={handleOpenConfigEvent}
                            on:draw={handleDrawEvent}
                            on:newRound={handleNewRoundEvent}
                            on:stopBingoAlert={handleCancelBingoAlertEvent}
                        />
                    </div>
                </div>
            {/if}
        </MediaQuery>
        <MediaQuery query="(max-width: 1149px)" let:matches>
            {#if matches}
                <div class=" info-card">
                    <div class="info-card-header">
                        <h2>Cartela de Bingo #{$card.Card}</h2>
                        <h3 class="info-card-header-round">
                            Sorteio #{$card.Round}
                        </h3>
                    </div>
                    <div class="cardHeader">
                        <CardHeader
                            on:autoplay={handleAutoplayEvent}
                            on:openConfig={handleOpenConfigEvent}
                            on:draw={handleDrawEvent}
                            on:newRound={handleNewRoundEvent}
                            on:stopBingoAlert={handleCancelBingoAlertEvent}
                        />
                    </div>
                </div>
            {/if}
        </MediaQuery>
    {/if}
    {#if $card.Card > 1}
        <div class="info-card-client" class:table_client class:table_draw>
            <div class="info-card-header">
                <h2>Cartela de Bingo #{$card.Card}</h2>
                <h3 class="info-card-header-round">
                    Sorteio #{$card.Round}
                </h3>
            </div>
            <div class="cardHeader">
                <CardHeader
                    on:autoplay={handleAutoplayEvent}
                    on:openConfig={handleOpenConfigEvent}
                    on:draw={handleDrawEvent}
                    on:newRound={handleNewRoundEvent}
                    on:stopBingoAlert={handleCancelBingoAlertEvent}
                />
            </div>
        </div>
    {/if}

    <div
        class="table-card"
        class:table_client
        class:table_draw
        data-table_client={table_client}
    >
        <Card
            card={$card}
            on:checkNumber={handleCheckNumberEvent}
            on:playCheckSound={handlePlayCheckSoundEvent}
        />
    </div>
    </div>

    {#if $card.Card > 1 && $card.Bingo}
        <BingoCelebrationModal
            cardNumber={$card.Card}
            completion={$card.LastCompletion}
            {soundStatus}
            dismissError={bingoDismissError}
            dismissing={dismissingBingo}
            on:playSound={playBingoSound}
            on:dismiss={handleCancelBingoAlertEvent}
        />
    {/if}

    <TableModalCard title="Configurações">
        <Completions
            card={$card.Card}
            bind:config
            bind:completions={$card.Completions}
            on:saveCompletions={handleSaveCompletions}
        />
    </TableModalCard>

{:else}
    <PageTitle title="Inclubingo - Carregando cartela" game="Inclubingo" />
    <div class="loading-card" role="status" aria-live="polite">
        <span class="loading-indicator" aria-hidden="true"></span>
        <div>
            <h2>Preparando sua cartela</h2>
            <p>Aguarde enquanto carregamos os dados desta rodada.</p>
        </div>
    </div>
{/if}

<audio bind:this={bingoAudio} src="/sms.mp3" loop></audio>
<audio bind:this={checkAudio} src="/pop.mp3"></audio>

<style>
    :root {
        font-size: 62.5%;
    }
    h2 {
        margin-top: 0;
        font-size: 2.8rem;
    }
    h3 {
        font-size: 2.3rem;
    }

    .wake-lock-notice {
        display: flex;
        width: fit-content;
        max-width: calc(100% - 3rem);
        align-items: center;
        gap: 1rem;
        margin: 1rem auto;
        padding: 0.8rem 1.2rem;
        border-radius: 0.8rem;
        background: #174c86;
        color: #fff;
        font-size: 1.4rem;
    }

    .loading-card {
        display: flex;
        width: min(60rem, calc(100% - 3rem));
        min-height: 14rem;
        align-items: center;
        gap: 2rem;
        margin: 4rem auto;
        padding: 2.4rem;
        border-left: 0.6rem solid var(--primary-color, #2b7ef4);
        border-radius: 0.8rem;
        background: var(--white, #fff);
        box-shadow: 0 0.6rem 1.8rem rgba(29, 29, 29, 0.14);
    }

    .loading-card h2,
    .loading-card p {
        margin: 0;
    }

    .loading-card p {
        line-height: 1.6;
    }

    .loading-indicator {
        width: 4rem;
        height: 4rem;
        flex: 0 0 auto;
        border: 0.5rem solid #d8e5f8;
        border-top-color: var(--primary-color, #2b7ef4);
        border-radius: 50%;
        animation: loading-spin 0.9s linear infinite;
    }

    @keyframes loading-spin {
        to {
            transform: rotate(360deg);
        }
    }

    @media (prefers-reduced-motion: reduce) {
        .loading-indicator {
            animation: none;
            border-color: var(--primary-color, #2b7ef4);
        }
    }

    .table_draw {
        flex-grow: 1;
    }

    .table_client {
        width: 30%;
        flex-grow: 0 !important;
    }

    .container-card {
        margin-top: 35px;
        padding: 0;
        display: flex;
        flex-direction: row;
        flex-wrap: wrap;
        justify-content: space-between;
    }
    .info-card-header-round {
        margin-top: 30px;
    }
    .table-card {
        display: flex;
        justify-content: flex-start;
    }

    @media (min-width: 1150px) {
        .table-horizontal {
            display: grid;
            grid-template-areas:
                "container-qr_code info-card-header cardHeader"
                "tableCard tableCard tableCard";
            gap: 3.5rem 2rem;
        }
        .table-horizontal .info-card-header h2 {
            font-size: 2.4rem;
        }
        .info-card-header {
            grid-area: info-card-header;
        }
        .cardHeader {
            grid-area: cardHeader;
        }
        .table-card {
            grid-area: tableCard;
        }
        .container-qr_code {
            grid-area: container-qr_code;
        }

        .table-horizontal .info-card-header {
            display: flex;
            flex-direction: column;
            justify-content: space-between;
            width: 23rem;
            height: 7rem;
        }
        .table-horizontal .info-card-header-round {
            margin: 0;
            font-size: 2.4rem;
        }

        .table-horizontal .container-qr_code {
            width: 12rem;
            height: 7rem;
        }

        .table-horizontal .container-qr_code img {
            width: 100%;
        }
    }

    @media (max-width: 1149px) {
        .table_draw {
            justify-content: center;
        }
        .table-card[data-table_client="true"] {
            width: auto;
        }
    }
    @media (max-width: 991px) {
        .container {
            min-width: 90vw;
        }
    }

    @media (max-width: 970px) {
        .info-card-client,
        .table-card[data-table_client="true"] {
            flex-basis: 45%;
        }
    }

    @media (max-width: 910px) {
        .table-card {
            justify-content: center;
        }

    }
    @media (max-width: 767px) {
        .info-card-header {
            display: flex;
            flex-direction: row;
            justify-content: space-between;
        }
        .info-card-header-round {
            margin: 0px;
        }
    }
    @media (max-width: 450px) {
        h2,
        h3 {
            font-size: 2.2rem;
            margin: 0;
        }
        .info-card {
            width: 100%;
            margin: 0 1rem;
        }
        .container-card {
            margin-top: 17px;
        }
        .table-card[data-table_client="true"],
        .info-card-client {
            flex-basis: 100%;
        }

        .info-card-header {
            display: flex;
            margin: 0 1rem;
        }
    }

    @media (max-width: 410px) {
    }
    @media (max-width: 320px) {
        h2,
        h3 {
            font-size: 1.8rem;
        }
    }
</style>
