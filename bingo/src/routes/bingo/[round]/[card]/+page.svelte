<script>
    import { onMount } from "svelte";
    import { goto } from "$app/navigation";
    import CardHeader from "$lib/bingo/CardHeader.svelte";
    import Card from "$lib/bingo/Card.svelte";
    import { card } from "../../../../lib/bingo";
    import PageTitle from "$lib/PageTitle.svelte";
    import Adds from "$lib/Adds.svelte";
    import Play from "$lib/Play.svelte";
    import { callApi, getWSEndpoint } from "$lib/api";
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

    const handleAutoplayEvent = async () => {
        await updateCard(`/api/bingo/${$card.Round}/${$card.Card}/autoplay`);
    };

    const handleCheckNumberEvent = async (event = {}) => {
        if (!$card.Autoplay || $card.Card == 1) {
            await updateCard(
                `/api/bingo/${$card.Round}/${$card.Card}/${event.detail.Number}`
            );
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
        await updateCard(`/api/bingo/${$card.Round}/${$card.Card}/0`);
    };

    const handleNewRoundEvent = async () => {
        await updateCard(`/api/bingo/${$card.Round}/new/${$card.Type}`);
        goto(`/bingo/${$card.Round}`);
    };

    const handlePlayCheckSoundEvent = async () => {
        checkAudio.pause();
        checkAudio.currentTime = 0;
        checkAudio.play();
    };

    const handleCancelBingoAlertEvent = async () => {
        $card = await callApi(
            $card,
            `/api/bingo/${$card.Round}/${$card.Card}/cancel`,
            "GET"
        );

        isBingo();
    };

    const handleSaveCompletions = async () => {
        $card = await callApi(
            $card,
            `/api/bingo/${$card.Round}/${$card.Card}/completions`,
            "POST",
            $card.Completions
        );
    };

    const isBingo = () => {
        if ($card.Card > 1) {
            if ($card.Bingo) {
                bingoAudio.play();
            } else {
                bingoAudio.pause();
                bingoAudio.currentTime = 0;
            }
        }
    };

    const liveUpdater = async () => {
        const socket = new window.WebSocket(
            getWSEndpoint(`/ws/bingo/${$card.Round}/${$card.Card}`)
        );

        socket.addEventListener("open", (event) => {
            //this.update(JSON.parse(event.data));
        });

        socket.addEventListener("message", (event) => {
            $card = JSON.parse(event.data);
            redirectToNextRound();
            isBingo();
        });

        socket.addEventListener("close", (event) => {
            setInterval(poolingUpdater, 1000);
        });

        socket.addEventListener("error", (event) => {
            setInterval(poolingUpdater, 1000);
        });
    };

    const loadCard = async () => {
        $card.Card = Number(data.Card);
        $card.Round = Number(data.Round);

        if ("WebSocket" in window) {
            liveUpdater();
        } else {
            setInterval(poolingUpdater, 1000);
        }
    };

    const poolingUpdater = async () => {
        await updateCard(`/api/bingo/${$card.Round}/${$card.Card}`);
    };

    const redirectToNextRound = () => {
        if ($card.Round == 0) {
            goto("/");
        }

        if ($card.Card == 0) {
            goto(`/bingo/${$card.Round}/new`);
        }

        if ($card.NextRound > 0 && $card.Card > 1) {
            goto(`/bingo/${$card.NextRound}/new`);
        }
    };

    const updateCard = async (path = "") => {
        $card = await callApi($card, path, "GET");
        redirectToNextRound();
        isBingo();
    };

    export let table_client = false;

    $: table_client = $card.Card > 1 ? true : false;

    export let table_draw = false;

    $: table_draw = $card.Card == 1 ? true : false;

    onMount(loadCard);
</script>

<PageTitle
    title="Inclubingo - Cartela {$card.Card}, rodada {$card.Round}"
    game="Inclubingo"
/>
<div class="container container-card">
    {#if $card.Card == 1}
        <MediaQuery query="(min-width: 1150px)" let:matches>
            {#if matches}
                <div class=" info-card table-horizontal" class:table_draw>
                    <div class="container-qr_code">
                        <img src="/qr/bingo/{$card.Round}" alt="QR-Code" />
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
