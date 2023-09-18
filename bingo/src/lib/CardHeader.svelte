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
    <div class="container-header">
        <div class="container-header-inner">
            <Autoplay bind:card={card} />
            <Bingo bind:card={card} />
            <NewRound bind:card={card} />
            <Draw bind:card />
        </div>
        <div>
            
        </div>
        <div aria-live="polite">
            <p>{card.Checked} Bolas Sorteadas</p>
            <p>{card.LastNumber} Última bola sorteada</p>
        </div>
    </div>
</div>

<style>
    .container{
        padding: 0;
    }
    .container-header{
        display: flex;
        flex-direction: column; 
    }
    .container-header-inner{
        display: flex;
        margin-top: 50px;
        margin-bottom: 50px;
        justify-content: center;
    }
    p{
        font-size: 1.6em;
    }
    @media(max-width: 767px){
        p{
            margin: 0;
            padding: 0;
            font-size: 1.2em;
        }
        .container-header{
            flex-direction: row;
            justify-content: space-between;
        }
        .container-header-inner{
            margin: 0;
        }
    }
    @media(max-width: 450px){
        p{
            margin: 0;
            margin: 5px 0 7px 0px;
            font-size: 1em;
        }
    }
</style>