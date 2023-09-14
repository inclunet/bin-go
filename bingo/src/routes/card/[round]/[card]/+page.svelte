<script>
    import { onMount } from "svelte";
    import { goto } from "$app/navigation";
    import CardHeader from "$lib/CardHeader.svelte";
    import Card from "$lib/Card.svelte";
    import { getCard } from "../../../../lib/bingo";
    import PageTitle from "$lib/PageTitle.svelte";
    import Adds from "$lib/Adds.svelte";
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
    <div class="anuncio"><Adds/></div>
</div>

<div class=" container-card-mobile">
    <div class="info-card">
        <h2 class="text-center mt-2">Cartela de Bingo #{card.Card} rodada #{card.Round}</h2>
        <CardHeader bind:card />
    </div>
    <div class="">
        <div class="table-card">
            <Card bind:card />
        </div>
        <div class="container anuncio">
            <Adds/>
        </div>
    </div>
</div>

<style>
    .container-card{
        display: flex;
        margin-top: 35px;
        justify-content: space-between;
    }
    .anuncio{
        width: 20%;
    }
    .container-card-mobile{
        display: none;
    }
    @media(max-width: 762px){
        .container-card{
            display: none;
        }
        .container-card-mobile{
            display: flex;
            flex-direction: column;
        }
        .anuncio{
            width: 100%;
            margin-top: 10px
        }
    }
    @media(max-width: 450px){
    }
</style>
