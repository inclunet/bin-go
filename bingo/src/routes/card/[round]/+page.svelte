<script>
    import { goto } from "$app/navigation";
    import { page } from "$app/stores";
    import Card from "$lib/Card.svelte";
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

<h1>Rodada #{card.Round}</h1>
<p>
    Aponte a câmera do seu celular aqui para pegar a sua cartela ou acesse o
    link:
    <strong>
        <a href="{$page.url}/new" id="link_jogo">{$page.url}/new</a>
    </strong>
</p>
<div id="qr_code">
    <img src="/qr/{card.Round}/{card.Card}?url={$page.url}/new" alt="QR-Code" />
</div>
<StartRound bind:card />

<style>
    h1,
    p {
        text-align: center;
    }
    h1 {
        font-size: 3em;
    }
    p {
        font-size: 1.8em;
        line-height: 1.7em;
        padding: 0 20px 0 20px;
    }
    a#link_jogo {
        word-break: break-word;
    }
    button {
        display: block;
        margin: 0 auto;
        font-size: 1.8em;
        padding: 20px;
    }
    #qr_code {
        display: flex;
        justify-content: center;
        margin: 0 auto;
        margin-top: 40px;
        margin-bottom: 40px;
    }
</style>
