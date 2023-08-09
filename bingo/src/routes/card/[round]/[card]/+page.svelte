<script>
    import { onMount } from "svelte";
    import { getCard, playAudio, redirectToNextRound } from "../../../bingo.js";
    import { goto } from "$app/navigation";
    import Card from "./Card.svelte";

    export let data;
    let card = data;

    export async function updateCard() {
        let response = await getCard("/card/" + card.Round + "/" + card.Card);
        if (card.Card == 0) {
            card = response;
            goto("/card/" + card.Round + "/" + card.Card);
        } else {
            card = response;
            redirectToNextRound();
            playAudio();
            setTimeout(updateCard, 2000);
        }
    }

    onMount(updateCard);
</script>

<div>
    <h2>Cartela de Bingo #{card.Card} rodada #{card.Round}</h2>
    <Card bind:card />
</div>

<style>
    div {
        text-align: center;
    }
</style>
