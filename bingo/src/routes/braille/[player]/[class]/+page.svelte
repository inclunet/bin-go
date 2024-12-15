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
            "GET"
        );

        startClassAudio.play();
    };

    const handleChallengeSubmit = async () => {
        let currentPunctuation = $braille.CurrentPunctuation;

        $braille = await callApi(
            $braille,
            `/api/braille/${data.Player}`,
            "POST",
            $braille
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

<PageTitle
    title="Aula #{$braille.CurrentClass} - Braille Personal Trainer"
    game="Inclubraille"
/>

<div class="container">
    <header class="my-5">
        <h2>Aula #{$braille.CurrentClass} - Braille Personal Trainer</h2>
    </header>

    <section>
        <p>{$braille.Description}</p>
    </section>

    <section>
        <ul>
            <li>Exercício: {$braille.CurrentClass + 1}</li>
            <li>Rodada: {$braille.CurrentRound}</li>
            <li>Pontuação do exercício: {$braille.CurrentPunctuation}</li>
            <li>Pontuação total: {$braille.TotalPunctuation}</li>
        </ul>
    </section>

    <BrailleChallenge
        bind:brailleChallenge={$braille.Challenge}
        on:submitChallenge={handleChallengeSubmit}
    />
</div>

<audio bind:this={correctAudio} src="/correct.wav"></audio>
<audio bind:this={wrongAudio} src="/wrong.wav"></audio>
<audio bind:this={endClassAudio} src="/endclass.wav"></audio>
<audio bind:this={startClassAudio} src="/startclass.wav"></audio>

<style>
    :root {
        --primary-button-color: #2b7ef4;
        --secondary-button-color: #2868c2;
        --white: #fff;
        --black: #000;
    }

    h2:nth-child(2n) {
        font-size: 1.6em;
    }

    section p,
    section button {
        font-size: 1.2em;
    }

    p {
        line-height: 32px;
    }

    .button-color {
        background-color: var(--primary-button-color);
        color: var(--black);
    }

    .button-color:hover {
        background-color: var(--secondary-button-color);
        color: var(--white);
    }

    .term {
        padding: 3px 5px 3px 5px;
        margin: 0;
        font-size: 1em;
    }

    @media (max-width: 573px) {
        .container {
            padding: 0 20px 0 20px;
        }
        header h2 {
            margin-top: -20px;
            line-height: 40px;
            text-align: center;
        }

        p {
            line-height: 37px;
        }
        section p {
            text-align: justify;
        }
    }
</style>
