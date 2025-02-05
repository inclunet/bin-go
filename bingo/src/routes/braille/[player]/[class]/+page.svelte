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

    let items_points = [
        {
            name: "Exercício",
            value: $braille.CurrentClass + 1,
            color: "#1E90FF",
        },
        {
            name: "Rodada",
            value: $braille.CurrentRound,
            color: "#FF4500",
        },
        {
            name: "Pontuação do Exercício",
            value: $braille.CurrentPunctuation,
            color: "#42CD42",
        },
        {
            name: "Pontuação Total",
            value: $braille.TotalPunctuation,
            color: "#8A2BE2",
        },
    ];

    const upCardsPoints = () => {
        items_points[0].value = $braille.CurrentClass + 1;
        items_points[1].value = $braille.CurrentRound;
        items_points[2].value = $braille.CurrentPunctuation;
        items_points[3].value = $braille.TotalPunctuation;
        items_points = items_points;
    };

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

        upCardsPoints();
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
        <p class="paragraph-description">{$braille.Description}</p>
    </section>

    <section>
        <div class="container_card">
            {#each items_points as item, index}
                <div class="card text-center" style="width: 18rem;">
                    <div class="card-body">
                        <div
                            class="card-points"
                            style="background-color: {item.color};"
                        >
                            <p>{item.value}</p>
                        </div>
                        <h4 class="card-title">{item.name}</h4>
                    </div>
                </div>
            {/each}
        </div>
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

        font-size: 62.5%;
    }

    h2 {
        font-size: 2.8rem;
    }

    h4 {
        font-size: 2rem;
    }

    .container_card {
        margin: 50px 0;
        display: flex;
        justify-content: space-between;
    }

    .card {
        border: 1px solid rgb(102, 102, 102);
    }
    .card-body {
        padding: 0;
    }

    .card-points {
        font-weight: bold;
    }
    .card-points p {
        font-size: 4rem;
        color: var(--white);
        padding: 25px 0;
    }

    .card-title {
        font-weight: 600;
        padding: 8px;
    }
    .paragraph-description {
        font-size: 1.8rem;
    }

    .paragraph-description {
        line-height: 4rem;
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
    }
</style>
