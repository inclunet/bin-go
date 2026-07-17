<script>
    import NewRound from "$lib/NewRound.svelte";
    import PageTitle from "$lib/PageTitle.svelte";
    import { callApiResult } from "$lib/api";
    import { card } from "$lib/bingo";
    import {
        clearRoundCreationID,
        getRoundCreationID,
    } from "$lib/bingo/playerCards";

    let createError = "";
    let creating = false;
    let createdLobbyURL = "";

    const openCreatedLobby = () => {
        window.location.assign(createdLobbyURL);
        window.setTimeout(() => {
            creating = false;
        }, 2000);
    };

    const handleNewRoundEvent = async () => {
        if (creating) {
            return;
        }
        creating = true;
        createError = "";
        if (createdLobbyURL) {
            openCreatedLobby();
            return;
        }
        const result = await callApiResult(
            $card,
            `/api/bingo/0/new/${$card.Type}`,
            "GET",
            null,
            false,
            { "X-Bingo-Creation-ID": getRoundCreationID() }
        );
        if (!result.ok || !result.data.RoundID || !result.data.ID) {
            createError =
                "Não foi possível criar a rodada. Tente novamente em alguns instantes.";
            creating = false;
            return;
        }
        $card = result.data;
        createdLobbyURL = `/bingo/${$card.RoundID}/${$card.ID}/lobby`;
        clearRoundCreationID();
        openCreatedLobby();
    };
</script>

<PageTitle title="Nova Rodada" game="Inclubingo" />

<div class="container-fluid d-flex align-items-center flex-column">
    <h2 class="text-center">Bem-vindo ao <strong>Inclubingo</strong></h2>
    <p class="text-center my-3">
        Vamos Jogar! Escolha a quantidade de bolinhas que serão sorteadas:
    </p>
    <NewRound
        on:click={handleNewRoundEvent}
        disabled={creating}
        ariaBusy={creating}
    />
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
