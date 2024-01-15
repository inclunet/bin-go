<script>
    import { card } from "./bingo";
    import { createEventDispatcher } from "svelte";

    const dispatch = createEventDispatcher();

    function stopBingoAlert() {
        dispatch("stopBingoAlert");
    }
</script>

{#if $card.Card > 1 && $card.Bingo}
    <button class="btn btn-warning btn-lg" on:click={stopBingoAlert}
        >Bingo!</button
    >
{/if}
