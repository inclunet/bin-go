<script>
    import { page } from "$app/stores";
    import PageTitle from "$lib/PageTitle.svelte";
    import StartRound from "$lib/StartRound.svelte";
    import { onMount } from "svelte";
    import { card } from "$lib/bingo";
    import { callApi } from "$lib/api.js";
    import { goto } from "$app/navigation";

    export let data;

    const loadCard = async () => {
        $card = await callApi($card, `/api/bingo/${data.Round}/1`, "GET");
    };

    const handleCallToAction = async () => {
        goto(`/bingo/${$card.Round}/1`);
    };

    onMount(loadCard);
</script>

<PageTitle
    title="Sorteio de Cartelas, rodada {$card.Round}"
    game="Inclubingo"
/>

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
        <img src="/qr/bingo/{$card.Round}" alt="QR-Code" />
    </div>
    <StartRound on:callToAction={handleCallToAction} />
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

    a#link_jogo {
        word-break: break-word;
    }

    #qr_code {
        margin: 1rem 0 1.5rem 0;
    }
    img {
        width: 24.8rem;
        height: 20.8rem;
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
