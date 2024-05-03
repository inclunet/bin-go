<script>
    import { createEventDispatcher } from "svelte";

    export let brailleWord = "";
    export let brailleKeyboard;

    const dispatch = createEventDispatcher();

    const handleKeyDown = (
        /** @type {{ key: string; preventDefault: () => void; }} */ event,
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

            if (event.key === "Enter") {
                handleSubmit();
                event.preventDefault();
            }
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
                    break;
                case " ":
                    dispatch("brailleEnter", "space");
                    break;
            }
        }
    };

    const handleSubmit = () => {
        dispatch("submitChallenge");
    };
</script>

<div>
    <label for="word">Resposta:</label>
</div>

<div>
    <input
        type="text"
        id="word"
        bind:value={brailleWord}
        on:keydown={handleKeyDown}
        on:keyup={handleKeyUp}
        autocomplete="off"
        autofocus
    />
</div>

<div>
    <button on:click={handleSubmit}>Enviar</button>
</div>
