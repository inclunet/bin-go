<script>
    import { goto } from "$app/navigation";
    import { getCard } from "./bingo";

    export let card = {
        Card: 0,
        Checked: 0,
        Round: 0,
        Type: 0,
    };

    async function newRound() {
        card = await getCard("/card/" + card.Round + "/new/" + card.Type);
        window.location.href = "/card/" + card.Round;
    }
</script>

{#if (card.Card == 1 && card.Checked == card.Type) || card.Round == 0}
    <button class="btn btn-primary" on:click={newRound}><strong>Nova Rodada</strong></button>
{/if}

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