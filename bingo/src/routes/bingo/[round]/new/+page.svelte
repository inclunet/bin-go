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
        <h2>Sorteio #{roundNumber}</h2>
        <p class="text-center my-3">
            Para gerar sua cartela, clique no botão “Jogar” abaixo.
        </p>
        <Adds />
        <Play {roundID} />
    {:else}
        <div class="loading-invite" role="status" aria-live="polite">
            <span class="loading-indicator" aria-hidden="true"></span>
            <div>
                <h2>Preparando o convite</h2>
                <p>Aguarde enquanto carregamos os dados desta rodada.</p>
            </div>
        </div>
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

    .loading-invite {
        display: flex;
        width: min(60rem, calc(100% - 3rem));
        min-height: 14rem;
        align-items: center;
        gap: 2rem;
        margin: 4rem auto;
        padding: 2.4rem;
        border-left: 0.6rem solid var(--primary-color, #2b7ef4);
        border-radius: 0.8rem;
        background: var(--white, #fff);
        box-shadow: 0 0.6rem 1.8rem rgba(29, 29, 29, 0.14);
    }

    .loading-invite h2,
    .loading-invite p {
        margin: 0;
        padding: 0;
    }

    .loading-invite p {
        line-height: 1.6;
    }

    .loading-indicator {
        width: 4rem;
        height: 4rem;
        flex: 0 0 auto;
        border: 0.5rem solid #d8e5f8;
        border-top-color: var(--primary-color, #2b7ef4);
        border-radius: 50%;
        animation: loading-spin 0.9s linear infinite;
    }

    @keyframes loading-spin {
        to {
            transform: rotate(360deg);
        }
    }

    @media (prefers-reduced-motion: reduce) {
        .loading-indicator {
            animation: none;
            border-color: var(--primary-color, #2b7ef4);
        }
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
