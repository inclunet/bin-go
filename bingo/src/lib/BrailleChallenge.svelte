<script>
    import BrailleKeyboard from "./BrailleKeyboard.svelte";
    import BrailleWord from "./BrailleWord.svelte";

    export let brailleChallenge;
</script>

<div role="region" aria-label="Exercício">
    <header aria-live="polite">
        {#if brailleChallenge.Challenge == "word"}
            <h3>
                Escreva em Braille: <div class="desafio-texto">
                    <span class="texto-tinta">
                        {brailleChallenge.Word}
                    </span>
                </div>
            </h3>
        {:else}
            <h3>
                Escreva em tinta: <div class="desafio-texto">
                    <span class="texto-braille">
                        {brailleChallenge.Word}
                    </span>
                </div>
            </h3>
        {/if}
    </header>

    {#if brailleChallenge.Challenge == "braille"}
        <BrailleWord
            bind:brailleWord={brailleChallenge.Repply}
            brailleKeyboard={false}
            on:submitChallenge
        />
    {:else}
        <BrailleKeyboard
            bind:brailleWord={brailleChallenge.Repply}
            on:submitChallenge
        />
    {/if}
</div>

<style>
    header {
        display: flex;
    }
    h3 {
        font-size: 1.4em;
    }
    header h3 div {
        display: inline-block;
    }
    .texto-tinta {
        display: block;
        margin-top: -0.2em;
    }
    .texto-braille {
        display: block;
        margin: 0;
    }

    .desafio-texto {
        width: 1.3em;
        height: 1.5em;
        font-size: 1.8em;
        align-content: center;
        text-align: center;
        background: #00f279;
        color: #000;
        font-weight: 500;
        border-radius: 0.7em;
    }
</style>
