<script>
    import Card from "./Card.svelte";
    import { getCard } from "./bingo";

    export let card = {
        Card: 0,
        Checked: 0,
        Round: 0,
        Type: 0,
        Numbers: [],
    };

    async function drawNumber() {
        card = await getCard("/card/" + card.Round + "/1/0");
    }
</script>

{#if card.Card == 1 && card.Checked < card.Type}
    <button on:click={drawNumber}>Sortear</button>
{/if}
