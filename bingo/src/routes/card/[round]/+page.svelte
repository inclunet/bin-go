<script>
    import { goto } from "$app/navigation";

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
        http://localhost/card/{data.Round}/0
    </strong>
</p>
<button on:click={startNewGame}>Começar a rodada</button>
