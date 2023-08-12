<script>
    import { getCard } from "./bingo";

    export let card = {Card: 0, Round: 0};

    async function startGame() {
        card = await getCard("/card/" + card.Round + "/0");
        window.location.href = "/card/" + card.Round + "/" + card.Card
    }
</script>

<button on:click={startGame}>Jogar!</button>