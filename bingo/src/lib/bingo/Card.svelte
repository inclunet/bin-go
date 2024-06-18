<script>
    import CardDefault from "$lib/bingo/CardDefault.svelte";
    import CardHorizontal from "$lib/bingo/CardHorizontal.svelte";

    export let card = { Card: 0, Numbers: [] };
</script>

{#if card.Card == 1}
    <CardHorizontal {card} on:checkNumber on:playCheckSound />
{:else}
    <CardDefault {card} on:checkNumber on:playCheckSound />
{/if}
