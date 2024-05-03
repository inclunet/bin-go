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
    on:click={handleClick}>{brailleDot}</button
>
