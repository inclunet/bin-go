<script>
    import { goto } from "$app/navigation";
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

    const handlePlay = async() => {
        goto(`/braille/${data.Player}/${$braille.CurrentClass}`);
    };

    onMount(LoadBrailleClass);
</script>

<PageTitle title="Aula #{$braille.CurrentClass} - Braille Personal Trainer" />

<div>
<h2>Aula #{$braille.CurrentClass} - Braille Personal Trainer</h2>

<p>{$braille.Description}</p>



<button on:click={handlePlay}>Jogar</button>
</div>