<script>
    import { createEventDispatcher } from "svelte";
    import { braille } from "$lib/braille.js";
    import { onMount, onDestroy } from "svelte";

    export let brailleCell = 0;
    export let brailleWord = "";
    export let brailleKeyboard;

    let audio;
    let isClient = typeof window !== "undefined";
    let buttonEnviar;
    let elementFocus;

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
        if (brailleWord.length == 0) {
            dispatch("enableSpaceTip");
        } else {
            dispatch("submitChallenge");
        }
    };

    const convertToLowerCase = (event) => {
        /**Todo:
         * verificar qual é a class (nível, ou conjunto de exercícios) que corresponde ao sinal de letra maiúscula para poder realizar a validação correta
         * */
        if ($braille.CurrentClass + 1 != 0) {
            const value = event.target.value;
            brailleWord = value.toLowerCase();
        }
    };

    const handleHotKeys = (event) => {
        if (event.ctrlKey && event.altKey && event.key.toLowerCase() === "n") {
            event.preventDefault();
            buttonEnviar.click();
        }
    };

    onMount(() => {
        if (isClient) {
            document.addEventListener("keydown", handleHotKeys);
        }
        elementFocus.focus();
    });

    onDestroy(() => {
        if (isClient) {
            document.removeEventListener("keydown", handleHotKeys);
        }
    });
</script>

<section id="answer">
    <label for="word">Resposta:</label>
    <input
        type="text"
        id="word"
        bind:value={brailleWord}
        bind:this={elementFocus}
        on:keydown={handleKeyDown}
        on:keyup={handleKeyUp}
        on:input={convertToLowerCase}
        autocomplete="off"
    />
    <button
        disabled={brailleWord.length == 0 && brailleCell == 0}
        class="btn button-color"
        on:click={handleSubmit}
        bind:this={buttonEnviar}>Enviar</button
    >
</section>

<audio src="/key.mp3" bind:this={audio} />

<style>
    :root {
        --primary-button-color: #2b7ef4;
        --secondary-button-color: #2868c2;
        --white: #fff;
        --black: #000;

        font-size: 62.5%;
    }

    section#answer {
        display: flex;
        width: 100%;
        margin: 2.5rem 0;
        align-items: center;
        font-size: 1.8rem;
    }

    section input {
        width: 100%;
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

    .button-color:active {
        background-color: var(--primary-button-color);
        color: var(--white);
    }

    @media (hover: hover) {
        .button-color:hover {
            background-color: var(--secondary-button-color);
            color: var(--white);
        }
    }

    @media (max-width: 490px) {
        section#answer {
            margin: 1rem 0;
        }
    }
</style>
