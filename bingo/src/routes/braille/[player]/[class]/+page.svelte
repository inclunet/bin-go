<script>
    import { goto } from "$app/navigation";
    import BrailleChallenge from "$lib/BrailleChallenge.svelte";
    import PageTitle from "$lib/PageTitle.svelte";
    import { callApi } from "$lib/api.js";
    import { braille } from "$lib/braille.js";
    import { onMount } from "svelte";

    export let data;

    /**
     * @type {HTMLAudioElement}
     */
    let correctAudio;
    /**
     * @type {HTMLAudioElement}
     */
    let wrongAudio;
    /**
     * @type {HTMLAudioElement}
     */
    let endClassAudio;
/**
     * @type {HTMLAudioElement}
     */
    let startClassAudio;

    const LoadBrailleClass = async () => {
        $braille = await callApi(
            $braille,
            `/api/braille/${data.Player}`,
            "GET",
        );

        startClassAudio.play();
    };

    const handleChallengeSubmit = async () => {
        let currentPunctuation = $braille.CurrentPunctuation;

        $braille = await callApi(
            $braille,
            `/api/braille/${data.Player}`,
            "POST",
            $braille,
        );

        if (data.Class < $braille.CurrentClass) {
            endClassAudio.play();
            await new Promise((r) => setTimeout(r, 2000));
            goto(`/braille/${data.Player}`);
        }

        if (currentPunctuation < $braille.CurrentPunctuation) {
            correctAudio.play();
        } else {
            wrongAudio.play();
        }
    };

    onMount(LoadBrailleClass);
</script>

<PageTitle title="Aula #{$braille.CurrentClass} - Braille Personal Trainer" />

<h2>Aula #{$braille.CurrentClass} - Braille Personal Trainer</h2>
<div>
    {$braille.Description}
</div>
<div>
    <ul>
        <li>Exercício: {$braille.CurrentClass + 1}</li>
        <li>Rodada: {$braille.CurrentRound}</li>
        <li>Pontuação do exercício: {$braille.CurrentPunctuation}</li>
        <li>Pontuação total: {$braille.TotalPunctuation}</li>
    </ul>
</div>

<BrailleChallenge
    bind:brailleChallenge={$braille.Challenge}
    on:submitChallenge={handleChallengeSubmit}
/>

<audio bind:this={correctAudio} src="/correct.wav"></audio>
<audio bind:this={wrongAudio} src="/wrong.wav"></audio>
<audio bind:this={endClassAudio} src="/endclass.wav"></audio>
<audio bind:this={startClassAudio} src="/startclass.wav"></audio>
