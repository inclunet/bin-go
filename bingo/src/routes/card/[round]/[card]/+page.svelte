<script>
    import { onMount } from "svelte";
    import { goto } from "$app/navigation";
    import CardHeader from "$lib/CardHeader.svelte";
    import Card from "$lib/Card.svelte";
    import { getCard } from "../../../../lib/bingo";
    import PageTitle from "$lib/PageTitle.svelte";

    export let data;
    let card = data;

    function redirectToNextRound() {
        if (card.Round == 0) {
            window.location.href = "/";
        }

        if (card.Card == 0) {
            window.location.href = "/card/" + card.Round + "/new";
        }

        if (card.NextRound > 0) {
            window.location.href = "/card/" + card.NextRound + "/new";
        }
    }

    async function updateCard() {
        redirectToNextRound();
        card = await getCard("/card/" + card.Round + "/" + card.Card);
        setTimeout(updateCard, 2000);
    }

    onMount(updateCard);
</script>

<PageTitle title="Cartela de Bingo" />

<div class=" mx-5 container-card">
    <div class="info-card">
        <h2>Cartela de Bingo #{card.Card}</h2>
        <h3 class="mt-4">Rodada #{card.Round}</h3>
        <CardHeader bind:card />
    </div>
    <div class="table-card">
        <Card bind:card />
    </div>
    <div class="anuncio">Anúncio</div>
</div>

<style>
    .container-card{
        margin-top: 35px;
        display: flex;
        justify-content: space-between;
    }
    .info-card{
        /* border: 1px solid black; */
    }
    .table-card{
    }
    .anuncio{
        width: 20%;
        border: 1px solid black;
    }
</style>
