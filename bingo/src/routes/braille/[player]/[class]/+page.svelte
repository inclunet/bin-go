<script>
    import { goto } from "$app/navigation";
    import BrailleChallenge from "$lib/BrailleChallenge.svelte";
    import PageTitle from "$lib/PageTitle.svelte";
    import { callApi } from "$lib/api.js";
    import { braille } from "$lib/braille.js";
    import { StarIcon, TrophyIcon } from "$lib/braille/icons.js";
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

    let progress = 0;

    function progressClass() {
        return Math.round(
            (($braille.CurrentRound - 0) / ($braille.TotalRounds - 0)) * 100
        );
    }

    $: if ($braille.CurrentRound > 0) {
        progress = progressClass();
    }

    onMount(LoadBrailleClass);
</script>

<PageTitle
    title="Aula #{$braille.CurrentClass} - Braille Personal Trainer"
    game="Inclubraille"
/>

<div class="container">
    <div class="container_progress">
        <div
            class="progress"
            role="progressbar"
            aria-label="Progresso da rodada em {progress}%"
            aria-valuenow={$braille.CurrentRound}
            aria-valuemin="0"
            aria-valuemax={$braille.TotalRounds}
        >
            <div
                class="progress-bar progress-bar-striped bg-info text-dark"
                style="width: {progress}%;"
            >
                {progress}%
            </div>
        </div>
        {@html TrophyIcon(30, "gold")}
    </div>

    <section class="container_play">
        <BrailleChallenge
            bind:brailleChallenge={$braille.Challenge}
            on:submitChallenge={handleChallengeSubmit}
        />
    </section>
</div>

<audio bind:this={correctAudio} src="/correct.wav"></audio>
<audio bind:this={wrongAudio} src="/wrong.wav"></audio>
<audio bind:this={endClassAudio} src="/endclass.wav"></audio>
<audio bind:this={startClassAudio} src="/startclass.wav"></audio>

<style>
    .container_progress {
        width: 100%;
        display: flex;
        align-items: center;
        justify-content: space-between;
        margin: 1.5rem;
    }

    .progress {
        width: 96%;
        height: 2rem;
        border-radius: 10px;
    }

    .progress-bar {
        font-size: 1.8rem;
        font-weight: 500;
    }
    .container {
        min-height: 100vh;
    }

    @media (max-width: 490px) {
        .container_play {
            min-height: 110vw;
        }
    }

    @media (max-width: 470px) {
        .container {
            padding: 0 20px 0 20px;
        }
    }
</style>
