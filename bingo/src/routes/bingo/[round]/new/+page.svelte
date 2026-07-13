<script>
    import Adds from "$lib/Adds.svelte";
    import PageTitle from "$lib/PageTitle.svelte";
    import Play from "$lib/Play.svelte";
    import { onMount } from "svelte";
    import { card } from "$lib/bingo";
    import { callApiResult } from "$lib/api.js";

    export let data;
    let loadError = "";
    let loaded = false;
    let roundID = "";
    let roundNumber = 0;

    const loadRound = async () => {
        roundID = String(data.Round);
        const result = await callApiResult(
            $card,
            `/api/bingo/${roundID}`,
            "GET"
        );
        if (!result.ok || !result.data.ID || !result.data.Round) {
            loadError =
                "Não foi possível abrir esta rodada. Verifique o link e tente novamente.";
            return;
        }
        roundNumber = result.data.Round;
        loaded = true;
    };

    onMount(loadRound);
</script>

<PageTitle title="Nova cartela de Inclubingo" game="Inclubingo" />

<div class="container-fluid d-flex align-items-center flex-column">
    {#if loadError}
        <p class="alert alert-danger text-center" role="alert">{loadError}</p>
    {:else if loaded}
        <h2>Rodada #{roundNumber}</h2>
        <p class="text-center my-3">
            Para gerar sua cartela, clique no botão “Jogar” abaixo.
        </p>
        <Adds />
        <Play {roundID} />
    {/if}
</div>

<style>
    :root {
        font-size: 62.5%;
    }

    h2 {
        margin-top: 40px;
        font-size: 2.8rem;
    }
    p {
        padding: 0 10px 0 10px;
        font-size: 1.8rem;
    }
    @media (max-width: 450px) {
        h2 {
            margin: 0;
            padding: 10px 0 10px 0;
            font-size: 2.3rem;
        }
        p {
            margin: 0;
            margin-bottom: 20px;
        }
    }
</style>
