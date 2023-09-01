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
    <div id="qr_code" class="d-flex justify-content-center">
        <img
            src="/qr/{card.Round}/{card.Card}?url={$page.url}/new"
            alt="QR-Code"
        />
    </div>
    <StartRound bind:card />
</div>

<style>
    * {
        margin: 0;
        padding: 0;
    }

    h2 {
        margin-top: 50px;
        font-size: 2.5em;
    }
    p {
        margin-top: 25px;
        padding: 0 10px 0 10px;
        font-size: 1.6em;
    }

    a#link_jogo {
        word-break: break-word;
    }

    #qr_code {
        margin: 0 auto;
        margin-top: 40px;
        margin-bottom: 40px;
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
            font-size: 1em;
        }
    }
</style>
