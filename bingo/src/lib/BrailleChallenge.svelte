<script>
    import BrailleKeyboard from "./BrailleKeyboard.svelte";
    import BrailleWord from "./BrailleWord.svelte";

    export let brailleChallenge;
</script>

<div role="region" aria-label="Exercício">
    <header aria-live="polite">
        {#if brailleChallenge.Challenge == "word"}
            <h3 class="mt-4">
                Escreva em Braille: <span>{brailleChallenge.Word}</span>
            </h3>
        {:else}
            <h3 class="mt-4">
                Escreva em tinta: <span>{brailleChallenge.Word}</span>
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
    h3 {
        font-size: 1.4em;
    }
    header span {
        font-size: 1.4em;
    }
</style>
