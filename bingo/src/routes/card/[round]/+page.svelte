<script>
    import { page } from "$app/stores";
    import PageTitle from "$lib/PageTitle.svelte";
    import StartRound from "$lib/StartRound.svelte";
    import { onMount } from "svelte";
    import { card, getCard } from "$lib/bingo";

    export let data;

    const loadCard = async () => {
        $card.Card = 1;
        $card.Round = Number(data.Round);
        $card = await getCard("/card/" + $card.Round + "/" + $card.Card);
    };

    onMount(loadCard);
</script>

<PageTitle title="Inclubingo - Sorteio de Cartelas, rodada {$card.Round}" />

<div class="container-fluid d-flex align-items-center flex-column">
    <h2 class="text-center">Rodada #{$card.Round}</h2>

    <p class="text-center">
        Aponte a câmera do seu celular aqui para pegar a sua cartela ou acesse o
        link:
        <strong>
            <a href="{$page.url}/new" id="link_jogo">{$page.url}/new</a>
        </strong>
    </p>
    <div id="qr_code" class="d-flex">
        <img
            src="/qr/{$card.Round}/{$card.Card}?url={$page.url}/new"
            alt="QR-Code"
        />
    </div>
    <StartRound />
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
        margin: 10px 0px 15px 0px;
    }
    img {
        width: 248px;
        height: 208px;
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
