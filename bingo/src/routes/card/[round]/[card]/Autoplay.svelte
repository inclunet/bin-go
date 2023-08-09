<script>

    import { toggleAutoplay } from "../../../bingo";
    import Card from "./Card.svelte";
export let card = {Autoplay: false, Card: 0};
</script>
{#if card.Card > 1}
<button aria-pressed={card.Autoplay} on:click={toggleAutoplay}>
                {#if card.Autoplay}
                    Automático
                {:else}
                    Manual
                {/if}
            </button>
        {/if}