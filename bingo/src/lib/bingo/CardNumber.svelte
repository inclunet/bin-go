<script>
    import { afterUpdate, beforeUpdate, createEventDispatcher } from "svelte";

    export let number = { Checked: false, Number: 0 };
    export let mainCard = false;
    const dispatch = createEventDispatcher();
    let muted = false;

    const handleClick = () => {
        dispatch("checkNumber", number);
    };

    beforeUpdate(() => {
        if (!number.Checked && number.Number > 0) {
            muted = false;
        }
    });

    afterUpdate(() => {
        if (number.Checked && !muted && number.Number > 0) {
            muted = true;
            dispatch("playCheckSound", number);
        }
    });
</script>

{#if number.Number != 0}
    {#if mainCard}
        <button
            aria-pressed={number.Checked}
            on:click={handleClick}
            class="btn active button-card-draw"
        >
            {number.Number}
        </button>
    {:else}
        <button
            aria-pressed={number.Checked}
            on:click={handleClick}
            class="btn active button-card-client"
        >
            {number.Number}
        </button>
    {/if}
{:else}
    <img src="/img/favicon.png" alt="Logotipo Inclunet" />
{/if}

<style>
    :root {
        font-size: 62.5%;
    }
    img {
        width: 72px;
        height: 72px;
    }

    button {
        padding: 1rem 0 !important;
        text-align: center;
    }
    button[aria-pressed="true"] {
        background-color: #008000;
        color: white;
        border-color: black;
    }
    .button-card-draw {
        margin: 6px;
        width: 52px;
        color: black;
        font-weight: 900;
        font-size: 2.5rem;
    }
    .button-card-client {
        margin: 6px;
        width: 54px;
        height: 54px;
        color: black;
        font-weight: 900;
        font-size: 2.5em;
    }

    @media (max-width: 1132px) {
        .button-card-draw {
            margin: 6px;
            width: 47px;
            color: black;
            font-weight: 900;
            font-size: 1.2em;
        }
    }

    @media (max-width: 1059px) {
        .button-card-draw {
            margin: 4px;
            width: 47px;
            color: black;
            font-weight: 900;
            font-size: 1em;
        }
    }

    @media (max-width: 999px) {
        .button-card-draw {
            margin: 3px;
            width: 45px;
            color: black;
            font-weight: 900;
            font-size: 1em;
        }
    }
    @media (max-width: 939px) {
        .button-card-draw {
            margin: 6px;
            width: 52px;
            color: black;
            font-weight: 900;
            font-size: 1.4em;
        }
    }

    @media (max-width: 680px) {
        img {
            width: 80px;
            height: 80px;
        }
        .button-card-client {
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
        .button-card-client {
            margin: 0.5em;
            width: 50px;
            height: 50px;
            font-size: 2rem;
        }
    }

    @media (max-width: 400px) {
        .button-card-client {
            margin: 0.3rem;
        }
    }

    @media (max-width: 350px) {
        img {
            width: 70px;
            height: 70px;
        }
        .button-card-client {
            margin: 0.2em;
            width: 45px;
            height: 45px;
            font-size: 2rem;
        }
    }
</style>
