<script>
    import BrailleChallenge from "$lib/BrailleChallenge.svelte";
    import PageTitle from "$lib/PageTitle.svelte";
    import { callApi } from "$lib/api.js";
    import { braille } from "$lib/braille.js";
    import { onMount } from "svelte";

    export let data;

    const LoadBrailleClass = async () => {
        $braille = await callApi(
            $braille,
            `/api/braille/${data.Player}`,
            "GET",
        );
    };

    const handleChallengeSubmit = async () => {
        console.log("calling api");
        $braille = await callApi(
            $braille,
            `/api/braille/${data.Player}`,
            "POST",
            $braille,
        );

        console.log($braille);
    };

    onMount(LoadBrailleClass);
</script>

<PageTitle title="Braille Personal Trainer" />

<h2>Braille Personal Trainer</h2>
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
