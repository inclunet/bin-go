<script>
    import Tooltip from "./Tooltip.svelte";
    import BrailleDot from "./BrailleDot.svelte";
    import BrailleWord from "./BrailleWord.svelte";

    export let brailleWord = "";
    let brailleCell = 0;
    let enableSpaceTip = false;
    let spacebar;

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

    const handleEnableSpaceTip = () => {
        if (brailleCell > 0) {
            spacebar.focus();
            enableSpaceTip = true;
        }
    };

    const handleDisableSpaceTip = () => {
        enableSpaceTip = false;
    };

    //
</script>

<div role="region" aria-label="Resposta" class="container">
    <div role="region" aria-label="Teclado Braille" class="container_keyboard">
        <div class="container_numbers">
            <div class="brailleDot_numbers">
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

            <div class="brailleDot_numbers">
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
        </div>
        <div class="container_buttons">
            <button on:click={handleClearKey} class="btn" id="limpar"
                >Limpar</button
            >

            <div class="container_button-tooltip">
                <button
                    bind:this={spacebar}
                    on:click={handleSpaceKey}
                    on:click={handleDisableSpaceTip}
                    class="btn"
                    id="espaco">Espaço</button
                >
                <Tooltip {enableSpaceTip} marginTop="4.5rem">
                    <p class="texto_tip">Ei, não se esqueça!</p>
                    <p class="texto_tip">
                        Pressione espaço para confirmar a letra que você quer
                        enviar.
                    </p>
                </Tooltip>
            </div>

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
            on:enableSpaceTip={handleEnableSpaceTip}
            bind:brailleWord
            bind:brailleCell
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
        --tooltip: #fcdd56ef;
        --white: #fff;
        --black: #000;

        font-size: 62.5%;
    }

    .container {
        display: flex;
        flex-direction: column;
    }
    .container_keyboard {
        width: 100%;
        display: flex;
        flex-direction: column;
        align-items: center;
        justify-content: center;
        margin-top: 1em;
        margin-bottom: 2em;
    }

    .brailleDot_numbers {
        display: flex;
    }
    .btn {
        color: var(--black);
        padding: 1rem 5rem;
        border: 1px solid var(--black);
        border-radius: 10rem;
        font-weight: 500;
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
    }
    #limpar {
        background-color: var(--primary-button-limpar-color);
    }

    #limpar:active {
        background-color: var(--secondary-button-limpar-color);
        color: var(--white);
    }
    #espaco {
        background-color: var(--primary-button-espaco-color);
        padding-left: 15rem;
        padding-right: 15rem;
    }

    #espaco:active {
        background-color: var(--secondary-button-espaco-color);
        color: var(--white);
    }
    .texto_tip {
        margin: 0;
        padding: 0;
        font-size: 1.2em;
    }
    #backspace {
        background-color: var(--primary-button-backspace-color);
    }

    #backspace:active {
        background-color: var(--secondary-button-backspace-color);
        color: var(--white);
    }

    button {
        font-size: 1.8rem;
    }

    .container_numbers {
        display: flex;
        gap: 5rem;
    }
    .container_buttons {
        width: 100%;
        margin-top: 2rem;
        display: flex;
        justify-content: space-around;
    }
    .container_button-tooltip {
        display: flex;
        flex-direction: column;
        align-items: center;
        font-size: 1.2rem;
    }

    @media (hover: hover) {
        #limpar:hover {
            background-color: var(--secondary-button-limpar-color);
            color: var(--white);
        }
        #espaco:hover {
            background-color: var(--secondary-button-espaco-color);
            color: var(--white);
        }
        #backspace:hover {
            background-color: var(--secondary-button-backspace-color);
            color: var(--white);
        }
    }
    @media (max-width: 767px) {
        .container_buttons {
            justify-content: space-between;
        }
        .btn {
            width: 15rem;
            padding: 1rem 0;
        }
        #espaco {
            padding-left: 5rem;
            padding-right: 5rem;
        }
    }

    @media (max-width: 547px) {
        .container_numbers {
            gap: 2rem;
        }
    }

    @media (max-width: 518px) {
        button {
            margin: 0;
        }
        .container_numbers {
            gap: 0.5rem;
        }
    }

    @media (max-width: 490px) {
        .container_numbers {
            background-color: #696b6d;
            border-radius: 8rem;
            padding: 2rem;
        }
        .brailleDot_numbers {
            flex-direction: column;
        }

        .brailleDot_numbers:nth-child(1) {
            flex-direction: column-reverse;
        }

        .sete-oito {
            margin-top: 1rem;
        }
        .tres-seis {
            margin-top: 1rem;
        }
        .dois-cinco {
            margin-top: 1rem;
        }
        .btn {
            width: 12rem;
            font-size: 2rem;
            padding: 0 1rem;
            font-weight: 500;
        }
        #espaco {
            padding: 1rem 0;
        }
    }
</style>
