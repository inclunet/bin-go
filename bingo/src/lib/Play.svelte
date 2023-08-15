<script>
    import { getCard } from "./bingo";

    export let card = {Card: 0, Round: 0};

    async function startGame() {
        card = await getCard("/card/" + card.Round + "/0");
        window.location.href = "/card/" + card.Round + "/" + card.Card
    }
</script>

<button class="btn btn-primary" on:click={startGame}><strong>Jogar!</strong></button>

<style>
    button{
        padding: 10px;
        font-size: 1.3em;
        color: #000;
    }
    @media(max-width: 450px){
        button{
            height: 48px;
            padding: 7px;
            font-size: 1em;
        }
    }
</style>