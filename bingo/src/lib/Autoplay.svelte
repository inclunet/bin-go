<script>
    import { getCard } from "./bingo";
    export let card = { Autoplay: false, Card: 0 };

    async function toggleAutoplay() {
        if (card.Autoplay) {
            card.Autoplay = false;
        } else {
            card.Autoplay = true;
        }

        card = await getCard(
            "/card/" + card.Round + "/" + card.Card + "/autoplay"
        );
    }
</script>

{#if card.Card > 1}
    {#if card.Autoplay}
        <button class="btn btn-success" on:click={toggleAutoplay}> <strong>Automático</strong> </button>
    {:else}
        <button class="btn btn-warning" on:click={toggleAutoplay}> Manual </button>
    {/if}
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