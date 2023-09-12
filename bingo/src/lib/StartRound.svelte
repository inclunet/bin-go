<script>
    export let card;

    function startRound() {
        window.location.href = "/card/" + card.Round + "/1";
    }
</script>

<button on:click={startRound} class="btn text-black"><strong>Começar a Rodada</strong></button>

<style>
    :root{
        --primary-button-color: #2b7ef4;
        --secondary-button-color: #2868c2;
        --tertiary-button-color: #2157a1;
    }
    button{
        padding: 10px;
        font-size: 1.3em;
        background-color: var(--primary-button-color);
    }
    button:hover{
        color: #fff;
        background-color: var(--secondary-button-color);
    }
    .btn:active{
        color: #fff;
        background-color: var(--tertiary-button-color);
    }
    @media(max-width: 450px){
        button{
            height: 48px;
            padding: 7px;
            font-size: 1em;
        }
    }
</style>