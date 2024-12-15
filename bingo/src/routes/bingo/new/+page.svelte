<script>
    import { goto } from "$app/navigation";
    import NewRound from "$lib/NewRound.svelte";
    import PageTitle from "$lib/PageTitle.svelte";
    import { callApi } from "$lib/api";
    import { card } from "$lib/bingo";

    const handleNewRoundEvent = async () => {
        $card = await callApi($card, `/api/bingo/0/new/${$card.Type}`, "GET");
        goto(`/bingo/${$card.Round}`);
    };
</script>

<PageTitle title="Nova Rodada" game="Inclubingo" />

<div class="container-fluid d-flex align-items-center flex-column">
    <h2 class="text-center">Bem-vindo ao <strong>Inclubingo</strong></h2>
    <p class="text-center my-3">
        Vamos Jogar! Escolha a quantidade de bolinhas que serão sorteadas:
    </p>
    <NewRound on:click={handleNewRoundEvent} />
</div>

<style>
    :root {
        font-size: 62.5%;
    }
    h2 {
        margin-top: 4rem;
        font-size: 2.8rem;
    }
    p {
        padding: 0 1rem 0 1rem;
        font-size: 1.8rem;
    }
    @media (max-width: 450px) {
        h2 {
            margin: 0;
            padding: 1rem 0 1rem 0;
            font-size: 2.3rem;
        }
        p {
            margin: 0;
            margin-bottom: 2rem;
        }
    }
</style>
