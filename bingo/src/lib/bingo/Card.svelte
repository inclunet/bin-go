<script>
    import CardDefault from "$lib/bingo/CardDefault.svelte";
    import CardHorizontal from "$lib/bingo/CardHorizontal.svelte";
    import MediaQuery from "$lib/MediaQuery.svelte";

    export let card = { Card: 0, Numbers: [] };
</script>

{#if card.Card == 1}
    <MediaQuery query="(min-width: 1150px)" let:matches>
        {#if matches}
            <CardHorizontal {card} on:checkNumber on:playCheckSound />
        {/if}
    </MediaQuery>
    <MediaQuery query="(max-width: 1149px)" let:matches>
        {#if matches}
            <CardDefault {card} on:checkNumber on:playCheckSound />
        {/if}
    </MediaQuery>
{:else}
    <CardDefault {card} on:checkNumber on:playCheckSound />
{/if}
