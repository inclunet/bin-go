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
    aria-label={`Ponto ${brailleDot}`}
    aria-pressed={Boolean(brailleCell & (1 << (brailleDot - 1)))}
    on:click={handleClick}
>
    <span aria-hidden="true">
        {brailleDot}
    </span>
</button>

<style>
    button {
        height: 8rem;
        margin: 0.3rem 1.5rem !important;
        padding: 2rem;
        border-radius: 3rem;
        border: 1px solid var(--black);
        background-color: #fcdcc0;
        font-size: 3rem !important;
        color: var(--black);
    }

    button[aria-pressed="true"] {
        background-color: #00ff00;
    }

    @media (max-width: 518px) {
        button {
            margin: 0 0.3rem !important;
        }
    }

    @media (max-width: 490px) {
        button {
            width: 8.5rem;
            height: 8.5rem;
            margin: 0.5rem 2rem !important;
            padding: 0;
            font-size: 4rem;
            border-radius: 50%;
            border: 3px solid black;
            font-weight: 700;
        }
    }
</style>
