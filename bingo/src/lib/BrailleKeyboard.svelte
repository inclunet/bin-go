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
    <div role="region" aria-label="Teclado Braille" id="container-keyboard">
        <div>
            <button on:click={handleClearKey} class="btn" id="limpar"
                >Limpar</button
            >
        </div>

        <div class="brailleDot-numbers">
            <div class="sete-oito">
                <BrailleDot
                    on:brailleKey={handleBrailleKey}
                    bind:brailleCell
                    brailleDot={7}
                />
            </div>
            <div class="tres-seis">
                <BrailleDot
                    on:brailleKey={handleBrailleKey}
                    bind:brailleCell
                    brailleDot={3}
                />
            </div>
            <div class="dois-cinco">
                <BrailleDot
                    on:brailleKey={handleBrailleKey}
                    bind:brailleCell
                    brailleDot={2}
                />
            </div>
            <div class="um-quatro">
                <BrailleDot
                    on:brailleKey={handleBrailleKey}
                    bind:brailleCell
                    brailleDot={1}
                />
            </div>
        </div>

        <div>
            <button on:click={handleSpaceKey} class="btn" id="espaco"
                >Espaço</button
            >
        </div>

        <div class="brailleDot-numbers">
            <div class="um-quatro">
                <BrailleDot
                    on:brailleKey={handleBrailleKey}
                    bind:brailleCell
                    brailleDot={4}
                />
            </div>
            <div class="dois-cinco">
                <BrailleDot
                    on:brailleKey={handleBrailleKey}
                    bind:brailleCell
                    brailleDot={5}
                />
            </div>
            <div class="tres-seis">
                <BrailleDot
                    on:brailleKey={handleBrailleKey}
                    bind:brailleCell
                    brailleDot={6}
                />
            </div>
            <div class="sete-oito">
                <BrailleDot
                    on:brailleKey={handleBrailleKey}
                    bind:brailleCell
                    brailleDot={8}
                />
            </div>
        </div>

        <div>
            <button on:click={handleBackspaceKey} class="btn" id="backspace"
                >Backspace</button
            >
        </div>
    </div>
    <div>
        <BrailleWord
            on:brailleKey={handleBrailleKey}
            on:brailleEnter={handleBrailleTypping}
            on:submitChallenge
            bind:brailleWord
            brailleKeyboard="true"
        />
    </div>
</div>

<style>
    :root {
        --primary-button-limpar-color: #ffff00;
        --secondary-button-limpar-color: #ffa500;
        --primary-button-espaco-color: #1e90ff;
        --secondary-button-espaco-color: #246bb3;
        --primary-button-backspace-color: #9c9c9c;
        --secondary-button-backspace-color: #5a5656;
        --white: #fff;
        --black: #000;
    }
    #container-keyboard {
        width: 55vw;
        display: flex;
        align-items: center;
        justify-content: space-around;
        margin-top: 2em;
        margin-bottom: 4em;
    }

    .brailleDot-numbers {
        display: flex;
    }
    .btn {
        color: var(--black);
        margin-top: 4.5em;
    }
    .sete-oito {
        display: block;
        margin-top: 2.2em;
    }
    .tres-seis {
        display: block;
        margin-top: 1em;
    }
    .dois-cinco {
        display: block;
        margin-top: 0.3em;
    }
    .um-quatro {
        display: block;
        /* margin-top: 2.7em; */
    }

    #limpar {
        background-color: var(--primary-button-limpar-color);
    }
    #limpar:hover {
        background-color: var(--secondary-button-limpar-color);
        color: var(--white);
    }
    #limpar:active {
        background-color: var(--primary-button-limpar-color);
    }
    #espaco {
        background-color: var(--primary-button-espaco-color);
    }
    #espaco:hover {
        background-color: var(--secondary-button-espaco-color);
        color: var(--white);
    }
    #espaco:active {
        background-color: var(--primary-button-espaco-color);
    }
    #backspace {
        background-color: var(--primary-button-backspace-color);
    }
    #backspace:hover {
        background-color: var(--secondary-button-backspace-color);
        color: var(--white);
    }
    #backspace:active {
        background-color: var(--primary-button-backspace-color);
    }
</style>
