<script>
    import NewRound from "$lib/NewRound.svelte";
    import PageTitle from "$lib/PageTitle.svelte";
</script>

<PageTitle title="Nova Rodada" game="Inclubingo" />

<div class="container-fluid d-flex align-items-center flex-column">
    <h2 class="text-center">Bem-vindo ao <strong>Inclubingo</strong></h2>
    <p class="text-center my-3">
        Vamos Jogar! Escolha a quantidade de bolinhas que serão sorteadas:
    </p>
    <NewRound />
</div>

<style>
    h2 {
        margin-top: 40px;
        font-size: 2em;
    }
    p {
        padding: 0 10px 0 10px;
        font-size: 1.2em;
    }
    @media (max-width: 450px) {
        h2 {
            margin: 0;
            padding: 10px 0 10px 0;
            font-size: 1.5em;
        }
        p {
            margin: 0;
            margin-bottom: 20px;
        }
    }
</style>
