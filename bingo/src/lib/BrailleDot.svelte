<script>
    import { createEventDispatcher } from "svelte";

    export let brailleCell = 0;
    export let brailleDot = 0;

    const dispatcher = createEventDispatcher();

    const handleClick = () => {
        dispatcher("brailleKey", { key: brailleDot });
    };
</script>

<button
    aria-pressed={Boolean(brailleCell & (1 << (brailleDot - 1)))}
    on:click={handleClick}
    >{brailleDot}
</button>

<style>
    button {
        margin: 0 0.4em;
        padding: 1em;
        border-radius: 2em;
        border: 1px solid black;
        background-color: #fcdcc0;
    }

    button[aria-pressed="true"] {
        background-color: #00ff00;
    }
</style>
