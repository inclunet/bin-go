<script>
    import { onDestroy, onMount } from "svelte";
    import CardHeader from "$lib/bingo/CardHeader.svelte";
    import Card from "$lib/bingo/Card.svelte";
    import {
        getPendingLobbyURL,
        rememberPendingLobbyURL,
    } from "$lib/bingo/playerCards";
    import { card } from "../../../../lib/bingo";
    import PageTitle from "$lib/PageTitle.svelte";
    import Adds from "$lib/Adds.svelte";
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
    let pollingInFlight = false;
    let leavingPage = false;
    let cardLoaded = false;
    let loadError = "";
    let actionError = "";

    const isValidCard = (value) =>
        value.ID &&
        value.RoundID &&
        value.Round > 0 &&
        value.Card > 0;

    const startPolling = () => {
        if (!leavingPage && !pollingInterval) {
            pollingInterval = setInterval(poolingUpdater, 1000);
        }
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

    const handleNewRoundEvent = async () => {
        const updated = await updateCard(
            `/api/bingo/${$card.RoundID}/${$card.ID}/new/${$card.Type}`
        );
        if (!updated) {
            return;
        }
        const lobbyURL = `/bingo/${$card.RoundID}/${$card.ID}/lobby`;
        rememberPendingLobbyURL(lobbyURL);
        leavingPage = true;
        socket?.close();
        window.location.assign(lobbyURL);
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
        const result = await callApiResult(
            $card,
            `/api/bingo/${$card.RoundID}/${$card.ID}/cancel`,
            "GET"
        );
        if (!result.ok || !isValidCard(result.data)) {
            actionError =
                "Não foi possível atualizar a cartela. Tente novamente.";
            return;
        }
        actionError = "";
        $card = result.data;
        isBingo();
    };

    const handleSaveCompletions = async () => {
        const result = await callApiResult(
            $card,
            `/api/bingo/${$card.RoundID}/${$card.ID}/completions`,
            "POST",
            $card.Completions
        );
        if (!result.ok || !isValidCard(result.data)) {
            actionError =
                "Não foi possível salvar as configurações. Tente novamente.";
            return;
        }
        actionError = "";
        $card = result.data;
    };

    const isBingo = () => {
        if ($card.Card > 1 && bingoAudio) {
            if ($card.Bingo) {
                bingoAudio.play();
            } else {
                bingoAudio.pause();
                bingoAudio.currentTime = 0;
            }
        }
    };

    const liveUpdater = async () => {
        socket = new window.WebSocket(
            getWSEndpoint(`/ws/bingo/${$card.RoundID}/${$card.ID}`)
        );

        socket.addEventListener("open", (event) => {
            //this.update(JSON.parse(event.data));
        });

        socket.addEventListener("message", (event) => {
            try {
                const updated = JSON.parse(event.data);
                if (!isValidCard(updated)) {
                    throw new Error("invalid card update");
                }
                $card = updated;
                actionError = "";
                loadError = "";
                redirectToNextRound();
                isBingo();
            } catch (error) {
                actionError =
                    "A atualização em tempo real foi interrompida. Tentando reconectar.";
                socket?.close();
                startPolling();
            }
        });

        socket.addEventListener("close", (event) => {
            startPolling();
        });

        socket.addEventListener("error", (event) => {
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
        if (pollingInFlight) {
            return;
        }
        pollingInFlight = true;
        try {
            await updateCard(
                `/api/bingo/${$card.RoundID}/${$card.ID}`,
                false
            );
        } finally {
            pollingInFlight = false;
        }
    };

    const redirectToNextRound = () => {
        if (leavingPage) {
            return;
        }

        if ($card.Round == 0) {
            leavingPage = true;
            socket?.close();
            window.location.assign("/");
            return;
        }

        if ($card.Card == 0) {
            leavingPage = true;
            socket?.close();
            window.location.assign(`/bingo/${$card.RoundID}/new`);
            return;
        }

        if ($card.NextRoundID && $card.Card > 1) {
            leavingPage = true;
            socket?.close();
            window.location.assign(`/bingo/${$card.NextRoundID}/new`);
        }
    };

    const updateCard = async (path = "", reportError = true) => {
        const result = await callApiResult($card, path, "GET");
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
        const pendingLobbyURL = getPendingLobbyURL();
        if (pendingLobbyURL) {
            leavingPage = true;
            window.location.assign(pendingLobbyURL);
            return;
        }
        loadCard();
    });
    onDestroy(() => {
        leavingPage = true;
        socket?.close();
        if (pollingInterval) {
            clearInterval(pollingInterval);
        }
    });
</script>

{#if loadError}
    <PageTitle title="Inclubingo - Cartela indisponível" game="Inclubingo" />
    <p class="alert alert-danger text-center" role="alert">{loadError}</p>
{:else if cardLoaded}
    <PageTitle
        title="Inclubingo - Cartela {$card.Card}, rodada {$card.Round}"
        game="Inclubingo"
    />
    {#if actionError}
        <p class="alert alert-danger text-center" role="alert">{actionError}</p>
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
                            Rodada #{$card.Round}
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
                            Rodada #{$card.Round}
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
                    Rodada #{$card.Round}
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
    <div
        class="anuncio"
        class:table_client
        class:table_draw
        data-table_client={table_client}
    >
        <Adds />
    </div>
    </div>

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
                "tableCard tableCard tableCard"
                "anuncio anuncio anuncio";
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
        .anuncio {
            grid-area: anuncio;
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
        .anuncio {
            flex-grow: 0;
            width: 30%;
        }

        .anuncio[data-table_client="true"] {
            width: 25%;
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
        .anuncio[data-table_client="true"] {
            margin-top: 1rem;
            flex-basis: 100%;
        }
        .info-card-client,
        .table-card[data-table_client="true"] {
            flex-basis: 45%;
        }
    }

    @media (max-width: 910px) {
        .table-card {
            justify-content: center;
        }

        .anuncio {
            margin: 0;
            flex-grow: 1;
        }
    }
    @media (max-width: 767px) {
        .anuncio {
            flex-basis: 50%;
        }
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
        .anuncio {
            display: flex;
            justify-content: center;
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
