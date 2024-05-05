<script>
    import BrailleKeyboard from "./BrailleKeyboard.svelte";
    import BrailleWord from "./BrailleWord.svelte";

    export let brailleChallenge;
</script>

<div role="region" aria-label="Exercício">
    <div aria-live="polite">
        {#if brailleChallenge.Challenge == "word"}
            <h3>Escreva {brailleChallenge.Word} em Braille</h3>
        {:else}
            <h3>Escreva {brailleChallenge.Word} em tinta</h3>
        {/if}
    </div>

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
