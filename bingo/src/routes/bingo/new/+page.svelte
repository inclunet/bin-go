<script>
    import { goto } from "$app/navigation";
    import NewRound from "$lib/NewRound.svelte";
    import PageTitle from "$lib/PageTitle.svelte";
    import { callApiResult } from "$lib/api";
    import { card } from "$lib/bingo";

    let createError = "";

    const handleNewRoundEvent = async () => {
        const result = await callApiResult(
            $card,
            `/api/bingo/0/new/${$card.Type}`,
            "GET"
        );
        if (!result.ok || !result.data.RoundID || !result.data.ID) {
            createError =
                "Não foi possível criar a rodada. Tente novamente em alguns instantes.";
            return;
        }
        createError = "";
        $card = result.data;
        goto(`/bingo/${$card.RoundID}/${$card.ID}`);
    };
</script>

<PageTitle title="Nova Rodada" game="Inclubingo" />

<div class="container-fluid d-flex align-items-center flex-column">
    <h2 class="text-center">Bem-vindo ao <strong>Inclubingo</strong></h2>
    <p class="text-center my-3">
        Vamos Jogar! Escolha a quantidade de bolinhas que serão sorteadas:
    </p>
    <NewRound on:click={handleNewRoundEvent} />
    {#if createError}
        <p class="alert alert-danger" role="alert">{createError}</p>
    {/if}
</div>

<style>
    h2 {
        margin-top: 4rem;
    }
    p {
        padding: 0 1rem 0 1rem;
    }
    @media (max-width: 450px) {
        h2 {
            margin: 0;
            padding: 1rem 0 1rem 0;
        }
        p {
            margin: 0;
            margin-bottom: 2rem;
        }
    }
</style>
