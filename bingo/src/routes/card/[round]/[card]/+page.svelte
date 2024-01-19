<script>
    import { onMount } from "svelte";
    import { goto } from "$app/navigation";
    import CardHeader from "$lib/CardHeader.svelte";
    import Card from "$lib/Card.svelte";
    import { card, getCard, getEndpointUrl } from "../../../../lib/bingo";
    import PageTitle from "$lib/PageTitle.svelte";
    import Adds from "$lib/Adds.svelte";
    import Play from "$lib/Play.svelte";

    export let data;

    function redirectToNextRound() {
        if ($card.Round == 0) {
            goto("/");
        }

        if ($card.Card == 0) {
            goto("/card/" + $card.Round + "/new");
        }

        if ($card.NextRound > 0 && $card.Card > 1) {
            goto("/card/" + $card.NextRound + "/new");
        }
    }

    const liveUpdater = async () => {
        const socket = new window.WebSocket(
            getEndpointUrl(
                "/card/" + $card.Round + "/" + $card.Card + "/live",
                "ws",
            ),
        );

        socket.addEventListener("open", (event) => {
            //this.update(JSON.parse(event.data));
        });

        socket.addEventListener("message", (event) => {
            $card = JSON.parse(event.data);
            redirectToNextRound();
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
        $card = await getCard("/card/" + $card.Round + "/" + $card.Card);
        redirectToNextRound();
    };

    onMount(loadCard);
</script>

<PageTitle title="Inclubingo - Cartela {$card.Card}, rodada {$card.Round}" />

<div class="container container-card">
    <div class="mx-3 info-card">
        <div class="info-card-header">
            <h2>Cartela de Bingo #{$card.Card}</h2>
            <h3 class="info-card-header-round">Rodada #{$card.Round}</h3>
        </div>
        <CardHeader />
    </div>
    <div class="table-card">
        <Card  />
    </div>
    <div class="anuncio">
        <Adds />
    </div>
</div>

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
