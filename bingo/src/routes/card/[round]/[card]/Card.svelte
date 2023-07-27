<script>
    import { onMount } from "svelte";
    import CardNumber from "./CardNumber.svelte";
    import { goto } from "$app/navigation";
    import { page } from "$app/stores";

    export let card = { Card: 0, Round: 0, Checked: 0, Numbers: [[]] };

    function getEndpointUrl(call = "") {
        if ($page.url.port == "5173") {
            return (
                $page.url.protocol +
                "//" +
                $page.url.hostname +
                ":8080/api" +
                call
            );
        } else {
            return $page.url.protocol + "//" + $page.url.host + "/api" + call;
        }
    }

    async function getCard() {
        const url = getEndpointUrl("/card/" + card.Round + "/" + card.Card);
        console.log(url);
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
        const response = await fetch(
            getEndpointUrl(
                "/card/" +
                    card.Round +
                    "/" +
                    card.Card +
                    "/" +
                    event.detail.Number
            )
        );
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
                    <td
                        ><CardNumber
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
    table {
        margin: 0 auto;
        display: inline-block;
        border: 3px solid black;
        /* padding: 0.2em; */
        border-radius: 10px;
    }
    th {
        text-align: center;
    }
    table,
    li {
        margin-top: 20px;
    }
    li {
        font-size: 1.5em;
    }
</style>
