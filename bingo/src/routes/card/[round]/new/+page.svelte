<script>
    import Adds from "$lib/Adds.svelte";
    import PageTitle from "$lib/PageTitle.svelte";
    import Play from "$lib/Play.svelte";
    import { onMount } from "svelte";
    import { card } from "$lib/bingo";

    export let data;

    const loadCard = async () => {
        $card.Card = Number(data.Card);
        $card.Round = Number(data.Round);
    };

    onMount(loadCard);
</script>

<PageTitle title="Nova Cartela de Bingo - Rodada #{$card.Round}" />

<div class="container-fluid d-flex align-items-center flex-column">
    <h2>Rodada #{$card.Round}</h2>
    <p class="text-center my-3">Para jogar clique no botão "Jogar" a baixo</p>
    <Adds />
    <Play />
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
