<script>
    import Autoplay from "./Autoplay.svelte";
    import Draw from "./Draw.svelte";
    import NewRound from "./NewRound.svelte";

    export let card = {
        Autoplay: false,
        Card: 0,
        Checked: 0,
        LastNumber: 0,
        Round: 0,
    };
</script>

<div class="container d-flex align-items-center flex-column">
    <div class="d-flex flex-row align-items-baseline justify-content-around" id="auto_play">
        <Autoplay bind:card={card} />
        <NewRound bind:card={card} />
        <Draw bind:card />
        <p>{card.Checked} Bolas Sorteadas</p>
    </div>
    <div aria-live="polite">
        <p>Última bola sorteada: {card.LastNumber}</p>
    </div>
</div>

<style>
    #auto_play{
        width: 100%;
    }
    p{
        margin-top: 25px;
        padding: 0 10px 0 10px;
        font-size: 1.6em;
    }
    @media(max-width: 450px){
        p{
            margin: 0;
            margin: 5px 0 7px 0px;
            font-size: 1em;
        }
    }
</style>