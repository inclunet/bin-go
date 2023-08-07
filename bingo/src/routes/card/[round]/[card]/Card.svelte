<script>
    import { onMount } from "svelte";
    import CardNumber from "./CardNumber.svelte";
    import { goto } from "$app/navigation";
    import { page } from "$app/stores";
    let audio;
    export let card = {
        AutoPlay: false,
        Bingo: false,
        Card: 0,
        Round: 0,
        Checked: 0,
        NextRound: 0,
        Numbers: [[]],
    };

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
        const response = await fetch(url);
        const loadingCard = await response.json();
        if (card.Card > 0) {
            card = loadingCard;
            if (card.NextRound > 0) {
                goto("/card/" + card.NextRound + "/0");
            }
            if (card.Bingo) {
                audio.play();
            }
            setTimeout(getCard, 2000);
        } else {
            goto("/card/" + card.Round + "/" + loadingCard.Card);
            getCard();
        }
    }

    async function new75Card() {
        const url = getEndpointUrl("/card/" + card.Round + "/new/75");
        const response = await fetch(url);
        card = await response.json();
        goto("/card/" + card.Round);
    }

    async function sendNumber(
        event = { detail: { Checked: false, Number: 0 } }
    ) {
        let url = getEndpointUrl(
            "/card/" + card.Round + "/" + card.Card + "/" + event.detail.Number
        );
        const response = await fetch(url);
        const result = await response.json();
        getCard();
    }

    async function toggleAutoplay() {
        let url = getEndpointUrl(
            "/card/" + card.Round + "/" + card.Card + "/autoplay"
        );
        const response = await fetch(url);
        const result = await response.json();
        getCard();
    }

    async function drawNumber(
        event = { detail: { Checked: false, Number: 0 } }
    ) {
        let url = getEndpointUrl("/card/" + card.Round + "/1/0");
        const response = await fetch(url);
        const result = await response.json();
        getCard();
    }

    onMount(getCard);
</script>

<div>
    <ul>
        <div id="container_botao_jogo_utomatico">
            {#if card.Card == 1}
                <button on:click={new75Card}>Nova Rodada</button>
                <button class="auto_play btn" on:click={drawNumber}
                    >Sortear</button
                >
            {:else}
                <button
                    class="auto_play btn"
                    aria-pressed={card.AutoPlay}
                    on:click={toggleAutoplay}
                >
                    {#if card.AutoPlay}
                        Automático
                    {:else}
                        Manual
                    {/if}
                </button>
            {/if}
            <li>Bolas sorteadas: {card.Checked}</li>
        </div>
        <div id="container_ultima_bola">
            <li id="texto_ultima_bola">Última bola</li>
            <li id="ultima_bola" aria-live="polite">{card.LastNumber}</li>
        </div>
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
<audio src="/sms.mp3" bind:this={audio} />

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
    /* li {
        margin-top: 10px;
    } */
    ul {
        margin: 0;
        padding: 0;
    }
    li {
        font-size: 1.5em;
    }
    #container_botao_jogo_utomatico {
        width: 320px;
        display: flex;
        margin: 0 auto;
        justify-content: space-between;
    }
    .auto_play {
        color: black;
        background-color: #2b7ef4;
        font-weight: bold;
        height: 45px;
        padding: 0 10px 0 10px;

        font-size: 1.3em;
    }
    #container_ultima_bola {
        width: 200px;
        display: flex;
        margin: 0 auto;
        margin-bottom: 10px;
        padding: 0;
        justify-content: space-between;
    }
    #texto_ultima_bola {
        line-height: 50px;
    }
    #ultima_bola {
        display: block;
        background-color: #008000;
        width: 55px;
        height: 55px;
        /* line-height: 50px; */
        font-size: 2.3em;
        text-align: center;
        font-weight: bold;
        border-radius: 55px;
    }
</style>
