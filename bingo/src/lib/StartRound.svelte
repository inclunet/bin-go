<script>
    export let card;

    function startRound() {
        window.location.href = "/card/" + card.Round + "/1";
    }
</script>

<button on:click={startRound} class="btn btn-primary">Começar a rodada</button>
