<script>
    import { createEventDispatcher } from "svelte";
    import { card } from "$lib/bingo.js";
    import Autoplay from "$lib/Autoplay.svelte";
    import NewRound from "$lib/NewRound.svelte";
    import Bingo from "$lib/Bingo.svelte";
    import Button from "$lib/Button.svelte";
    import MediaQuery from "$lib/MediaQuery.svelte";

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

    const handleToggleFullScreen = () => {
        if (
            !document.fullscreenElement && // alternative standard method
            !document.mozFullScreenElement &&
            !document.webkitFullscreenElement &&
            !document.msFullscreenElement
        ) {
            // current working methods
            if (document.documentElement.requestFullscreen) {
                document.documentElement.requestFullscreen();
            } else if (document.documentElement.msRequestFullscreen) {
                document.documentElement.msRequestFullscreen();
            } else if (document.documentElement.mozRequestFullScreen) {
                document.documentElement.mozRequestFullScreen();
            } else if (document.documentElement.webkitRequestFullscreen) {
                document.documentElement.webkitRequestFullscreen(
                    Element.ALLOW_KEYBOARD_INPUT
                );
            }
        } else {
            if (document.exitFullscreen) {
                document.exitFullscreen();
            } else if (document.msExitFullscreen) {
                document.msExitFullscreen();
            } else if (document.mozCancelFullScreen) {
                document.mozCancelFullScreen();
            } else if (document.webkitExitFullscreen) {
                document.webkitExitFullscreen();
            }
        }
    };
</script>

<div class="container d-flex flex-column">
    {#if $card.Card == 1}
        <MediaQuery query="(min-width: 1150px)" let:matches>
            {#if matches}
                <div class="container-header container-header_horizontal">
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
                                    data_bs_target="#exampleModal"
                                    >Config</Button
                                >
                            </div>
                        {/if}

                        {#if $card.Card > 1 && $card.Bingo}
                            <Bingo on:click={handleDispatchStopBingoAlert} />
                        {/if}

                        {#if $card.Card == 1 && $card.Finished}
                            <NewRound on:click={handleDispatchNewRound} />
                        {/if}

                        {#if $card.Card == 1 && !$card.Finished}
                            <div class="container-botao card-botao_horizontal">
                                <Button on:click={handleDispatchDraw}
                                    >Sortear</Button
                                >
                                <Button
                                    color="dark"
                                    on:click={handleDispatchOpenConfig}
                                    data_bs_toggle="modal"
                                    data_bs_target="#exampleModal"
                                    >Config</Button
                                >
                                <Button
                                    color="info"
                                    on:click={handleToggleFullScreen}
                                    >Tela Cheia</Button
                                >
                            </div>
                        {/if}
                    </div>
                    <div aria-live="polite" class="container-status">
                        <p>Bolas Sorteadas: <strong>{$card.Checked}</strong></p>
                        <p>
                            Última sorteada: <strong>{$card.LastNumber}</strong>
                        </p>
                    </div>
                </div>
            {/if}
        </MediaQuery>

        <MediaQuery query="(max-width: 1149px)" let:matches>
            {#if matches}
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
                                    data_bs_target="#exampleModal"
                                    >Config</Button
                                >
                            </div>
                        {/if}

                        {#if $card.Card > 1 && $card.Bingo}
                            <Bingo on:click={handleDispatchStopBingoAlert} />
                        {/if}

                        {#if $card.Card == 1 && $card.Finished}
                            <NewRound on:click={handleDispatchNewRound} />
                        {/if}

                        {#if $card.Card == 1 && !$card.Finished}
                            <div class="container-botao">
                                <Button on:click={handleDispatchDraw}
                                    >Sortear</Button
                                >
                                <Button
                                    color="dark"
                                    on:click={handleDispatchOpenConfig}
                                    data_bs_toggle="modal"
                                    data_bs_target="#exampleModal"
                                    >Config</Button
                                >
                            </div>
                        {/if}
                    </div>
                    <div aria-live="polite" class="container-status">
                        <p>Bolas Sorteadas: <strong>{$card.Checked}</strong></p>
                        <p>
                            Última sorteada: <strong>{$card.LastNumber}</strong>
                        </p>
                    </div>
                </div>
            {/if}
        </MediaQuery>
    {/if}
    {#if $card.Card > 1}
        <div class="container-header-client">
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

                {#if $card.Card == 1 && $card.Finished}
                    <NewRound on:click={handleDispatchNewRound} />
                {/if}

                {#if $card.Card == 1 && !$card.Finished}
                    <div class="container-botao">
                        <Button on:click={handleDispatchDraw}>Sortear</Button>
                        <Button
                            color="dark"
                            on:click={handleDispatchOpenConfig}
                            data_bs_toggle="modal"
                            data_bs_target="#exampleModal">Config</Button
                        >
                        <Button color="info" on:click={handleToggleFullScreen}
                            >Tela Cheia</Button
                        >
                    </div>
                {/if}
            </div>
            <div aria-live="polite" class="container-status">
                <p>Bolas Sorteadas: <strong>{$card.Checked}</strong></p>
                <p>
                    Última sorteada: <strong>{$card.LastNumber}</strong>
                </p>
            </div>
        </div>
    {/if}
</div>

<style>
    :root {
        font-size: 62.5%;
    }
    .container {
        width: 100%;
        min-height: 0vh;
        padding: 0;
        margin: 0;
    }
    .container-header,
    .container-header-client {
        display: flex;
        flex-direction: column;
    }
    .container-header-client {
        width: 100%;
        display: flex;
        flex-direction: column;
    }
    .container-header-inner {
        display: flex;
        margin-top: 50px;
        margin-bottom: 50px;
    }
    .container-botao {
        width: 100%;
        margin: 0;
        padding: 0;
        display: flex;
        justify-content: space-between;
        align-items: center;
    }
    p {
        font-size: 1.6em;
    }
    .container-status {
        width: 15rem;
    }

    @media (min-width: 1150px) {
        .container-header_horizontal {
            display: grid;
            grid-template-areas: "container-status card-botao_horizontal ";
            gap: 0 2rem;
        }
        .container-header_horizontal .card-botao_horizontal {
            grid-area: card-botao_horizontal;
        }
        .container-header_horizontal .container-status {
            grid-area: container-status;
            display: flex;
            flex-direction: column;
            width: 18rem;
            height: 7rem;
            font-size: 1.2rem;
            justify-content: space-between;
        }
        .container-header_horizontal .container-status p {
            margin: 0;
        }
        .card-botao_horizontal {
            align-items: start;
        }
        .container-header_horizontal .container-header-inner {
            margin: 0;
        }
    }
    @media (max-width: 1149px) {
        .container-header {
            width: 25rem;
        }
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
            padding: 0;
            font-size: 1.5rem;
        }
        .container-botao {
            width: auto;
        }
        .container-status {
            width: auto;
            display: flex;
            flex-direction: column;
            align-items: center;
            margin: 0;
            padding: 0;
        }
        .container-header-client {
            width: auto;
            flex-direction: row;
            justify-content: space-between;
            margin: 0 1rem;
        }

        .container-header {
            width: 100%;
        }
    }

    @media (max-width: 410px) {
        .container-header-client,
        .container-header {
            flex-direction: column;
        }
        .container-status {
            display: flex;
            flex-direction: row;
            justify-content: space-between;
        }

        .container-botao {
            width: 100%;
        }
    }
</style>
