<script>
    import { card } from "./bingo.js";
    import Autoplay from "./Autoplay.svelte";
    import Draw from "./Draw.svelte";
    import NewRound from "./NewRound.svelte";
    import Bingo from "./Bingo.svelte";
</script>

<div class="container d-flex flex-column">
    <div class="container-header">
        <div class="container-header-inner">
            <Autoplay />
            <Bingo />
            
            {#if $card.Card == 1 && $card.Checked == $card.Type}
                <NewRound on:newRound />
            {/if}

            {#if $card.Card == 1 && $card.Checked < $card.Type}
                <Draw on:draw />
            {/if}
        </div>
        <div></div>
        <div aria-live="polite">
            <p><strong>{$card.Checked}</strong> Bolas Sorteadas</p>
            <p><strong>{$card.LastNumber}</strong> Última bola sorteada</p>
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
