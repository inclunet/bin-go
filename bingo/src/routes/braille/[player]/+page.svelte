<script>
    import { goto } from "$app/navigation";
    import PageTitle from "$lib/PageTitle.svelte";
    import { callApi } from "$lib/api.js";
    import { braille } from "$lib/braille.js";
    import YoutubePlayer from "$lib/YoutubePlayer.svelte";
    import { onMount } from "svelte";

    export let data;

    const LoadBrailleClass = async () => {
        $braille = await callApi(
            $braille,
            `/api/braille/${data.Player}`,
            "GET"
        );
    };

    const handlePlay = async () => {
        goto(`/braille/${data.Player}/${$braille.CurrentClass}`);
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

    <YoutubePlayer videoId="eGuK2wsBL48" autoplay={true} />

    <section>
        <p>{$braille.Description}</p>
    </section>

    <section>
        <button class="btn button-color" on:click={handlePlay}>
            <strong>Jogar</strong>
        </button>
    </section>
</div>

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

    section p,
    section button {
        font-size: 1.8rem;
    }

    p {
        line-height: 4rem;
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
