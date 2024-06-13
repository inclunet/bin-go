<script>
    import { onMount } from "svelte";
    import { goto } from "$app/navigation";
    import CardHeader from "$lib/CardHeader.svelte";
    import Card from "$lib/Card.svelte";
    import { card } from "../../../../lib/bingo";
    import PageTitle from "$lib/PageTitle.svelte";
    import Adds from "$lib/Adds.svelte";
    import Play from "$lib/Play.svelte";
    import { callApi, getWSEndpoint } from "$lib/api";
    import Completions from "$lib/bingo/Completions.svelte";
    import TableModalCard from "$lib/TableModalCard.svelte";

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
            if (!event.detail.Checked || $card.Card == 1) {
                await updateCard(
                    `/api/bingo/${$card.Round}/${$card.Card}/${event.detail.Number}`
                );
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

    onMount(loadCard);
</script>

<PageTitle
    title="Inclubingo - Cartela {$card.Card}, rodada {$card.Round}"
    game="Inclubingo"
/>

<div class="container container-card">
    <div class="mx-3 info-card">
        <div class="info-card-header">
            <h2>Cartela de Bingo #{$card.Card}</h2>
            <h3 class="info-card-header-round">Rodada #{$card.Round}</h3>
        </div>
        <CardHeader
            on:autoplay={handleAutoplayEvent}
            on:openConfig={handleOpenConfigEvent}
            on:draw={handleDrawEvent}
            on:newRound={handleNewRoundEvent}
            on:stopBingoAlert={handleCancelBingoAlertEvent}
        />
    </div>
    <div class="table-card">
        <Card
            on:checkNumber={handleCheckNumberEvent}
            on:playCheckSound={handlePlayCheckSoundEvent}
        />
    </div>
    <div class="anuncio">
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
    .info-card,
    .table-card {
        flex-grow: 2;
    }
    .anuncio {
        flex-grow: 1;
    }
    .anuncio {
        padding: 0;
        flex-basis: 21%;
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
        .info-card {
            width: 90%;
        }
        .container-card {
            margin-top: 17px;
        }
    }
</style>
