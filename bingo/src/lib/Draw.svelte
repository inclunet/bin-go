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
    <button class="btn btn-primary" on:click={drawNumber}><strong>Sortear</strong></button>
{/if}

<style>
    button{
        padding: 10px;
        font-size: 1.3em;
        color: #000;
    }
    @media(max-width: 450px){
        button{
            padding: 7px;
            font-size: 1em;
        }
    }
</style>