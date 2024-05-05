<script>
    import BrailleDot from "./BrailleDot.svelte";
    import BrailleWord from "./BrailleWord.svelte";

    export let brailleWord = "";
    let brailleCell = 0;

    const handleBrailleKey = (event = { detail: { key: 0 } }) => {
        if (event.detail.key == 0) {
            brailleCell = 0;
        }

        if (event.detail.key > 0 && event.detail.key < 9) {
            brailleCell ^= 1 << (event.detail.key - 1);
        }
    };

    const handleBrailleTypping = (event = { detail: "" }) => {
        if (brailleCell > 0 || event.detail == "space") {
            brailleWord += String.fromCharCode(0x2800 + brailleCell);
        }
        brailleCell = 0;
    };

    const handleSpaceKey = () => {
        handleBrailleTypping({ detail: "space" });
    };

    const handleBackspaceKey = () => {
        brailleWord = brailleWord.slice(0, -1);
    };

    const handleClearKey = () => {
brailleWord = "";
    };
</script>

<div role="region" aria-label="Resposta">
    <div>
        <BrailleWord
            on:brailleKey={handleBrailleKey}
            on:brailleEnter={handleBrailleTypping}
            on:submitChallenge
            bind:brailleWord
            brailleKeyboard="true"
        />
    </div>

    <div role="region" aria-label="Teclado Braille">
        <div>
            <button on:click={handleClearKey}>Limpar</button>
        </div>

        <div>
            <BrailleDot
                on:brailleKey={handleBrailleKey}
                bind:brailleCell
                brailleDot={7}
            />
            <BrailleDot
                on:brailleKey={handleBrailleKey}
                bind:brailleCell
                brailleDot={3}
            />
            <BrailleDot
                on:brailleKey={handleBrailleKey}
                bind:brailleCell
                brailleDot={2}
            />
            <BrailleDot
                on:brailleKey={handleBrailleKey}
                bind:brailleCell
                brailleDot={1}
            />
        </div>

        <div>
            <button on:click={handleSpaceKey}>Espaço</button>
        </div>

        <div>
            <BrailleDot
                on:brailleKey={handleBrailleKey}
                bind:brailleCell
                brailleDot={4}
            />
            <BrailleDot
                on:brailleKey={handleBrailleKey}
                bind:brailleCell
                brailleDot={5}
            />
            <BrailleDot
                on:brailleKey={handleBrailleKey}
                bind:brailleCell
                brailleDot={6}
            />
            <BrailleDot
                on:brailleKey={handleBrailleKey}
                bind:brailleCell
                brailleDot={8}
            />
        </div>

        <div>
            <button on:click={handleBackspaceKey}>Backspace</button>
        </div>
    </div>
</div>
