<script>
    import { createEventDispatcher } from "svelte";
    import { card } from "./bingo.js";
    import Autoplay from "./Autoplay.svelte";
    import Draw from "./Draw.svelte";
    import NewRound from "./NewRound.svelte";
    import Bingo from "./Bingo.svelte";
    import Botao from "./Botao.svelte";

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
            <div class="container-botao">
                {#if $card.Card > 1 && !$card.Bingo}
                    <Autoplay
                        autoplay={$card.Autoplay}
                        on:click={handleDispatchAutoplayEvent}
                    />
                {/if}

                <Botao color="btn-dark" on:click={handleDispatchOpenConfig}
                    >Config</Botao
                >
            </div>

            {#if $card.Card > 1 && $card.Bingo}
                <Bingo on:click={handleDispatchStopBingoAlert} />
            {/if}

            {#if $card.Card == 1 && $card.Checked == $card.Type}
                <NewRound on:click={handleDispatchNewRound} />
            {/if}

            {#if $card.Card == 1 && $card.Checked < $card.Type}
                <Draw on:click={handleDispatchDraw} />
            {/if}
        </div>
        <div></div>
        <div aria-live="polite">
            <p><strong>{$card.Checked}</strong> Bolas Sorteadas</p>
            <p><strong>{$card.LastNumber}</strong> Última sorteada</p>
        </div>
    </div>
</div>

<style>
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
        justify-content: center;
    }
    .container-botao {
        margin: 0;
        padding: 0;
        width: 12em;
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
    }
</style>
