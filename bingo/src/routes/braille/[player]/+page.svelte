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
    .container {
        min-height: auto;
    }
</style>
