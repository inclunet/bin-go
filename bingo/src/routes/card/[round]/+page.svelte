<script>
    import { goto } from "$app/navigation";
    import { page } from "$app/stores";

    export let data = { Round: 0, Card: 0 };

    function startNewGame() {
        goto("/card/" + data.Round + "/1");
    }
</script>

<h1>Rodada #{data.Round}</h1>
<p>
    Aponte a câmera do seu celular aqui para pegar a sua cartela ou acesse o
    link:
    <strong>
        <a href="{$page.url}/0">{$page.url}/0</a>
    </strong>
</p>
<img src="http://localhost:8080/qr/1/1?url={$page.url}/0" alt="QR-Code" />
<button on:click={startNewGame} class="btn btn-primary">Começar a rodada</button>
<style>
    h1, p{
        text-align: center;
    }
    button{
        display: block;
        margin: 0 auto;
    }
</style>
