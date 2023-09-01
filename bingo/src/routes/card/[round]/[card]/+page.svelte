<script>
    import { onMount } from "svelte";
    import { goto } from "$app/navigation";
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

<div class="container-fluid text-center d-flex align-items-center flex-column">
    <h2>Cartela de Bingo #{card.Card} rodada #{card.Round}</h2>
    <Card bind:card />
</div>

<style>
</style>
