<script>
    import { goto } from "$app/navigation";
    import { page } from "$app/stores";
    import { card, startNewGame } from "../../bingo.js";
</script>

<h1>Rodada #{card.Round}</h1>
<p>
    Aponte a câmera do seu celular aqui para pegar a sua cartela ou acesse o
    link:
    <strong>
        <a href="{$page.url}/0" id="link_jogo">{$page.url}/0</a>
    </strong>
</p>
<div id="qr_code">
    <img src="http://localhost:8080/qr/1/1?url={$page.url}/0" alt="QR-Code" />
</div>
<button on:click={startNewGame} class="btn btn-primary">Começar a rodada</button
>

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
