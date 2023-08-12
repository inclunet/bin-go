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
    <button on:click={newRound}>Nova Rodada</button>
{/if}
