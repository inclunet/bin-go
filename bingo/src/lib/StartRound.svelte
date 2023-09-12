<script>
    export let card;

    function startRound() {
        window.location.href = "/card/" + card.Round + "/1";
    }
</script>

<button on:click={startRound} class="btn button-color text-black"><strong>Começar a Rodada</strong></button>

<style>
    :root{
        --primary-button-color: #2b7ef4;
    }
    button{
        padding: 10px;
        font-size: 1.3em;
        background-color: var(--primary-button-color);
    }
    @media(max-width: 450px){
        button{
            height: 48px;
            padding: 7px;
            font-size: 1em;
        }
    }
</style>