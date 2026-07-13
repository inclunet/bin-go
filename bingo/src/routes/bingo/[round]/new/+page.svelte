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

    const loadCard = async () => {
        const result = await callApiResult(
            $card,
            `/api/bingo/${data.Round}/${data.Card}`,
            "GET"
        );
        if (!result.ok || !result.data.RoundID || !result.data.ID) {
            loadError =
                "Não foi possível criar sua cartela. Verifique o link e tente novamente.";
            return;
        }
        $card = result.data;
        loaded = true;
    };

    onMount(loadCard);
</script>

<PageTitle title="Nova cartela, rodada {$card.Round}" game="Inclubingo" />

<div class="container-fluid d-flex align-items-center flex-column">
    {#if loadError}
        <p class="alert alert-danger text-center" role="alert">{loadError}</p>
    {:else if loaded}
        <h2>Rodada #{$card.Round}</h2>
        <p class="text-center my-3">Para jogar clique no botão "Jogar" a baixo</p>
        <Adds />
        <Play />
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
