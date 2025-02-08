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
    :root {
        --black: #000;
        font-size: 62.5%;
    }
    button {
        margin: 0 0.4em;
        padding: 2rem;
        border-radius: 3rem;
        border: 1px solid black;
        background-color: #fcdcc0;
        font-size: 1.8rem;
        color: var(--black);
    }

    button[aria-pressed="true"] {
        background-color: #00ff00;
    }

    @media (max-width: 518px) {
        button {
            margin: 0 0.3rem;
        }
    }

    @media (max-width: 490px) {
        button {
            width: 8.5rem;
            height: 8.5rem;
            margin: 0.5rem 2rem;
            padding: 0;
            font-size: 4rem;
            border-radius: 50%;
            border: 3px solid black;
            font-weight: 700;
        }
    }
</style>
