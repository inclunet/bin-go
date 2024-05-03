<script>
    import BrailleKeyboard from "./BrailleKeyboard.svelte";
    import BrailleWord from "./BrailleWord.svelte";

    export let brailleChallenge;
</script>

<BrailleKeyboard bind:brailleWord={brailleChallenge.Word} />
