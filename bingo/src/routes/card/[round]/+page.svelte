<script>
    import { goto } from "$app/navigation";
    import { page } from "$app/stores";
    import Card from "$lib/Card.svelte";
    import PageTitle from "$lib/PageTitle.svelte";
    import StartRound from "$lib/StartRound.svelte";
    import { getCard } from "$lib/bingo.js";
    import { onMount } from "svelte";

    export let data;
    export let card = data;

    async function updateCard() {
        card = await getCard("/card/" + card.Round + "/1");
    }

    onMount(updateCard);
</script>

<PageTitle title="Sorteio de Cartelas" />

<div class="container-fluid d-flex align-items-center flex-column">
    <h2 class="text-center">Rodada #{card.Round}</h2>
    
    <p class="text-center">
        Aponte a câmera do seu celular aqui para pegar a sua cartela ou acesse o
        link:
        <strong>
            <a href="{$page.url}/new" id="link_jogo">{$page.url}/new</a>
        </strong>
    </p>
    <div id="qr_code" class="d-flex ">
        <img
            src="/qr/{card.Round}/{card.Card}?url={$page.url}/new"
            alt="QR-Code"
        />
    </div>
    <StartRound bind:card />
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

    a#link_jogo {
        word-break: break-word;
    }

    #qr_code {
        margin: 20px 0px 25px 0px;
    }
    img{
        width: 250px;
        height: 210px;
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
