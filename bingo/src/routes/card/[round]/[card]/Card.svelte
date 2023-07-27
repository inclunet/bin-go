<script>
    import { onMount } from "svelte";
    import CardNumber from "./CardNumber.svelte";
    import { goto } from "$app/navigation";

    export let card = { Card: 0, Round: 0, Checked: 0, Numbers: [[]] };

    async function getCard() {
        const url = "http://localhost:8080/api/card/" + card.Round + "/" + card.Card;
        const response = await fetch(url);
        const loadingCard = await response.json();
        if (card.Card > 0) {
            card = loadingCard;
        } else {
            goto("/card/" + card.Round + "/" + loadingCard.Card);
            getCard();
        }
    }

    async function sendNumber(event = {}) {
        const url =
            "http://localhost:8080/api/card/" +
            card.Round +
            "/" +
            card.Card +
            "/" +
            event.detail.Number;
        const response = await fetch(url);
        const result = await response.json();
        getCard();
    }

    onMount(getCard);
</script>

<div>
    <ul>
        <li>Bolas sorteadas: {card.Checked}</li>
    </ul>
    <table summary="Cartela">
        <tr id="tr_first">
            <th scope="col">B</th>
            <th scope="col">I</th>
            <th scope="col">N</th>
            <th scope="col">G</th>
            <th scope="col">O</th>
        </tr>
        {#each card.Numbers as row}
            <tr>
                {#each row as number}
                    <td><CardNumber
                            bind:number
                            on:checkNumber={sendNumber}
                        /></td
                    >
                {/each}
            </tr>
        {/each}
    </table>
</div>
<style>
    table{
        margin: 0 auto;
        display: inline-block;
        border: 3px solid black;
        /* padding: 0.2em; */
        border-radius: 10px;
    }
    th{
        text-align: center;
    }
    table, li{
        margin-top: 40px;
    }
    li{
        font-size: 1.5em;
    }

</style>