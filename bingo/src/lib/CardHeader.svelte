<script>
    import Autoplay from "./Autoplay.svelte";
    import Draw from "./Draw.svelte";
    import NewRound from "./NewRound.svelte";
    import Bingo from "./Bingo.svelte";

    export let card = {
        Autoplay: false,
        Card: 0,
        Checked: 0,
        LastNumber: 0,
        Round: 0,
    };
</script>

<div class="container d-flex flex-column">
    <div class="d-flex flex-column">
        <div class="d-flex my-5 justify-content-center">
            <Autoplay bind:card={card} />
            <Bingo bind:card={card} />
            <NewRound bind:card={card} />
            <Draw bind:card />
        </div>
        <div>
            <p>{card.Checked} Bolas Sorteadas</p>
            
        </div>
    </div>
    <div aria-live="polite">
        <p>Última bola sorteada: {card.LastNumber}</p>
    </div>
</div>

<style>
    .container{
        padding: 0;
    }
    p{
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