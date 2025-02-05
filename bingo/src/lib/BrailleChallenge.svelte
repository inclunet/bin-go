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

    <div class="container_write">
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
</div>

<style>
    header {
        display: flex;
    }
    h3 {
        font-size: 2.2rem;
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
        height: 5rem;
        padding: 0 3rem;
        font-size: 3rem;
        align-content: center;
        text-align: center;
        background: #00f279;
        color: #000;
        font-weight: 500;
        border-radius: 1.2rem;
    }

    .container_write {
        display: flex;
    }
</style>
