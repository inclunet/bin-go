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

<div class="container container-header">
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

<div class="container-mobile">
    <div class="d-flex flex-row container-mobile-header">
        <div class="">
            <Autoplay bind:card={card} />
            <Bingo bind:card={card} />
            <NewRound bind:card={card} />
            <Draw bind:card />
        </div>
        <div aria-live="polite">
            <p><strong>{card.Checked}</strong> Bolas Sorteadas</p>
            <p>Última bola sorteada: <strong>{card.LastNumber}</strong></p>
        </div>
    </div>
</div>

<style>
    .container-header{
        padding: 0;
    }
    .container-header{
        display: flex;
        flex-direction: column;
    }
    p{
        margin: 0;
        font-size: 1.6em;
    }
    .container-mobile{
        display: none;
    }
    @media(max-width: 762px){
        .container-header{
            display: none;
        }
        .container-mobile{
            display: flex;
            flex-direction: row;
            justify-content: center;
        }
        .container-mobile-header{
            width: 80%;
            justify-content: space-between;
            margin-bottom: 10px;
        }

    }
    @media(max-width: 450px){
        p{
            font-size: 1.2em;
        }
    }
</style>