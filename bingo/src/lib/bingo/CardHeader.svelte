<script>
    import { createEventDispatcher } from "svelte";
    import { card } from "$lib/bingo.js";
    import Autoplay from "$lib/Autoplay.svelte";
    import NewRound from "$lib/NewRound.svelte";
    import Bingo from "$lib/Bingo.svelte";
    import Button from "$lib/Button.svelte";

    const dispatch = createEventDispatcher();

    const handleDispatchAutoplayEvent = () => {
        dispatch("autoplay");
    };

    const handleDispatchOpenConfig = () => {
        dispatch("openConfig");
    };

    const handleDispatchDraw = () => {
        dispatch("draw");
    };

    const handleDispatchNewRound = async () => {
        dispatch("newRound");
    };

    const handleDispatchStopBingoAlert = () => {
        dispatch("stopBingoAlert");
    };
</script>

<div class="container d-flex flex-column">
    <div class="container-header">
        <div class="container-header-inner">
            {#if $card.Card > 1 && !$card.Bingo}
                <div class="container-botao">
                    <Autoplay
                        autoplay={$card.Autoplay}
                        on:click={handleDispatchAutoplayEvent}
                    />

                    <Button
                        color="dark"
                        on:click={handleDispatchOpenConfig}
                        data_bs_toggle="modal"
                        data_bs_target="#exampleModal">Config</Button
                    >
                </div>
            {/if}

            {#if $card.Card > 1 && $card.Bingo}
                <Bingo on:click={handleDispatchStopBingoAlert} />
            {/if}

            {#if $card.Card == 1 && $card.Checked == $card.Type}
                <NewRound on:click={handleDispatchNewRound} />
            {/if}

            {#if $card.Card == 1 && $card.Checked < $card.Type}
                <div class="container-botao">
                    <Button on:click={handleDispatchDraw}>Sortear</Button>
                    <Button
                        color="dark"
                        on:click={handleDispatchOpenConfig}
                        data_bs_toggle="modal"
                        data_bs_target="#exampleModal">Config</Button
                    >
                </div>
            {/if}
        </div>
        <div></div>
        <div aria-live="polite" class="container-status">
            <p><strong>{$card.Checked}</strong> Bolas Sorteadas</p>
            <p><strong>{$card.LastNumber}</strong> Última sorteada</p>
        </div>
    </div>
</div>

<style>
    :root {
        font-size: 62.5%;
    }
    .container {
        padding: 0;
    }
    .container-header {
        display: flex;
        flex-direction: column;
    }
    .container-header-inner {
        display: flex;
        margin-top: 50px;
        margin-bottom: 50px;
    }
    .container-botao {
        width: 25rem;
        margin: 0;
        padding: 0;
        display: flex;
        justify-content: space-between;
        align-items: center;
    }
    p {
        font-size: 1.6em;
    }
    @media (max-width: 767px) {
        p {
            margin: 0;
            padding: 0;
            font-size: 1.2em;
        }
        .container-header {
            flex-direction: row;
            justify-content: space-between;
        }
        .container-header-inner {
            margin: 0;
        }
    }
    @media (max-width: 450px) {
        p {
            margin: 0;
            margin: 5px 0 7px 0px;
            font-size: 1em;
        }
        .container-botao {
            width: 17.5rem;
        }
        .container-status p {
            font-size: 1.5rem;
        }
    }
</style>
