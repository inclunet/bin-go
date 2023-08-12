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
        <button on:click={toggleAutoplay}> Automático </button>
    {:else}
        <button on:click={toggleAutoplay}> Manual </button>
    {/if}
{/if}
