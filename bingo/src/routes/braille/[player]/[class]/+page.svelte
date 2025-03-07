<script>
    import { goto } from "$app/navigation";
    import BrailleChallenge from "$lib/BrailleChallenge.svelte";
    import PageTitle from "$lib/PageTitle.svelte";
    import { callApi } from "$lib/api.js";
    import Card from "$lib/bingo/Card.svelte";
    import { braille } from "$lib/braille.js";
    import CardPoints from "$lib/braille/CardPoints.svelte";
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

    <!-- <section>
        <p class="paragraph-description">{$braille.Description}</p>
    </section> -->

    <section>
        <h2 class="hidden-area">
            Nesta área contém a série atual, rodada e pontos
        </h2>
        <div class="container_card">
            <CardPoints
                pointBackgroundColor="blue"
                pointValue={$braille.CurrentClass + 1}
                pointTitle="Série"
                pointTitleHidden="Série"
            />
            <CardPoints
                pointBackgroundColor="orange"
                pointValue={$braille.CurrentRound}
                pointTitle="Rodada"
                pointTitleHidden="Pontos da Rodada"
            />
            <CardPoints
                pointBackgroundColor="green"
                pointValue={$braille.CurrentPunctuation}
                pointTitle="Exercício"
                pointTitleHidden="Pontos do exercício"
            />
            <CardPoints
                pointBackgroundColor="purple"
                pointValue={$braille.TotalPunctuation}
                pointTitle="Total"
                pointTitleHidden="Total de Pontos"
            />
        </div>
    </section>

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
    :root {
        --primary-button-color: #2b7ef4;
        --secondary-button-color: #2868c2;
        --white: #fff;
        --black: #000;

        font-size: 62.5%;
    }

    .container {
        min-height: 100vh;
    }

    h2 {
        font-size: 2.8rem;
    }

    .container_card {
        margin: 5rem 0;
        display: flex;
        justify-content: space-between;
    }
    .paragraph-description {
        font-size: 1.8rem;
    }

    .paragraph-description {
        line-height: 4rem;
    }

    @media (max-width: 490px) {
        .container_card {
            margin: 0;
        }
        .container_play {
            min-height: 110vw;
        }
    }

    @media (max-width: 470px) {
        .container {
            padding: 0 20px 0 20px;
        }
        header h2 {
            margin-top: -20px;
            line-height: 4rem;
            text-align: center;
        }

        p {
            line-height: 4.7rem;
        }
        section p {
            text-align: justify;
        }

        .container_card {
            /* flex-wrap: wrap; */
            /* agrupa em dois cards por linha no total de duas linhas */
        }
    }
</style>
