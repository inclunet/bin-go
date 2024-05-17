<script>
    import { createEventDispatcher } from "svelte";

    export let brailleWord = "";
    export let brailleKeyboard;

    let audio;

    const dispatch = createEventDispatcher();

    const handleKeyDown = (
        /** @type {{ key: string; preventDefault: () => void; }} */ event
    ) => {
        if (brailleKeyboard) {
            if (event.key === "f") {
                dispatch("brailleKey", { key: 1 });
                event.preventDefault();
            }

            if (event.key === "d") {
                dispatch("brailleKey", { key: 2 });
                event.preventDefault();
            }

            if (event.key === "s") {
                dispatch("brailleKey", { key: 3 });
                event.preventDefault();
            }

            if (event.key === "j") {
                dispatch("brailleKey", { key: 4 });
                event.preventDefault();
            }

            if (event.key === "k") {
                dispatch("brailleKey", { key: 5 });
                event.preventDefault();
            }

            if (event.key === "l") {
                dispatch("brailleKey", { key: 6 });
                event.preventDefault();
            }

            if (event.key === "a") {
                dispatch("brailleKey", { key: 7 });
                event.preventDefault();
            }

            if (event.key === "ç") {
                dispatch("brailleKey", { key: 8 });
                event.preventDefault();
            }

            if (event.key === " ") {
                dispatch("brailleKey", { key: 0 });
                event.preventDefault();
            }
        }
        if (event.key === "Enter") {
            handleSubmit();
            event.preventDefault();
        }
    };

    const handleKeyUp = (/** @type {{ key: any; }} */ event) => {
        if (brailleKeyboard) {
            switch (event.key) {
                case "a":
                case "s":
                case "d":
                case "f":
                case "j":
                case "k":
                case "l":
                case "ç":
                    dispatch("brailleEnter", "dots");
                    audio.play();
                    break;
                case " ":
                    dispatch("brailleEnter", "space");
                    audio.play();
                    break;
            }
        }
    };

    const handleSubmit = () => {
        dispatch("submitChallenge");
    };
</script>

<section id="answer">
    <label for="word">Resposta:</label>
    <input
        type="text"
        id="word"
        bind:value={brailleWord}
        on:keydown={handleKeyDown}
        on:keyup={handleKeyUp}
        autocomplete="off"
        autofocus
    />
    <button class="btn button-color" on:click={handleSubmit}>Enviar</button>
</section>

<audio src="/key.mp3" bind:this={audio} />

<style>
    :root {
        --primary-button-color: #2b7ef4;
        --secondary-button-color: #2868c2;
        --white: #fff;
        --black: #000;
    }

    section#answer {
        display: flex;
        width: 40vw;
        margin: 1.5em 0;
        align-items: center;
    }

    section input {
        width: 60%;
        margin-left: 0.8em;
        padding: 0.2em 0;
        padding-left: 0.8em;
    }
    section button {
        padding: 0.2em 0.6em;
        font-size: 1.2em;
        margin-left: 2em;
    }

    .button-color {
        background-color: var(--primary-button-color);
        color: var(--black);
    }

    .button-color:hover {
        background-color: var(--secondary-button-color);
        color: var(--white);
    }
    .button-color:active {
        background-color: var(--primary-button-color);
    }
</style>
