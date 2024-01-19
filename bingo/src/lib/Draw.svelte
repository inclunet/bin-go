<script>
    import { card, getCard } from "./bingo";

    async function drawNumber() {
        $card = await getCard("/card/" + $card.Round + "/" + $card.Card + "/0");
    }
</script>

{#if $card.Card == 1 && $card.Checked < $card.Type}
    <button class="btn btn-primary" on:click={drawNumber}
        ><strong>Sortear</strong></button
    >
{/if}

<style>
    button {
        padding: 10px;
        font-size: 1.3em;
        color: #000;
    }
    @media (max-width: 450px) {
        button {
            height: 48px;
            padding: 7px;
            font-size: 1em;
        }
    }
</style>
