<script>
    export let card;

    function startRound() {
        window.location.href = "/card/" + card.Round + "/1";
    }
</script>

<button on:click={startRound} class="btn btn-primary"><strong>Começar a Rodada</strong></button>

<style>
    button{
        padding: 10px;
        font-size: 1.3em;
        color: #000;
    }
    @media(max-width: 450px){
        button{
            padding: 7px;
            font-size: 1em;
        }
    }
</style>