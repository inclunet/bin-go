<script>
    import { sendCheckNumber } from "../../../bingo";
    export let card = { Card: 0, Round: 0 };
    export let number = { Checked: false, Number: 0 };

    async function checkNumber() {
        if (!number.Checked || card.Card == 1) {
            if (number.Checked) {
                number.Checked = false;
            } else {
                number.Checked = true;
            }
            card = await sendCheckNumber(number.Number);
        }
    }
</script>

{#if number.Number != 0}
    <button
        aria-pressed={number.Checked}
        on:click={checkNumber}
        class="btn active"
    >
        {number.Number}
    </button>
{:else}
    <img src="/favicon.png" alt="Inclunet" />
{/if}

<style>
    img {
        width: 100px;
        height: 100px;
    }
    button[aria-pressed="true"] {
        background-color: #008000;
        color: white;
        border-color: black;
    }
    button {
        margin: 1em;
        width: 70px;
        height: 70px;
        color: black;
        font-weight: 900;
        font-size: 1.8em;
    }
    @media (max-width: 680px) {
        img {
            width: 80px;
            height: 80px;
        }
        button {
            margin: 0.6em;
            width: 55px;
            height: 55px;
            font-weight: 900;
            font-size: 1.6em;
        }
    }
    @media (max-width: 470px) {
        img {
            width: 75px;
            height: 75px;
        }
        button {
            margin: 0.5em;
            width: 50px;
            height: 50px;
            font-size: 1.5em;
        }
    }
    @media (max-width: 400px) {
        button {
            margin: 0.3em;
            font-size: 1.3em;
        }
    }
    @media (max-width: 350px) {
        img {
            width: 70px;
            height: 70px;
        }
        button {
            margin: 0.2em;
            width: 45px;
            height: 45px;
            font-size: 1.1em;
        }
    }
</style>
