<script>
    import Adds from "$lib/Adds.svelte";
    import Play from "$lib/Play.svelte";

export let data;
export let card = data;
</script>
<div class="container-fluid d-flex align-items-center flex-column">
    <h2>Rodada #{card.Round}</h2>
    <p>Para jogar clique no botão "Jogar" a baixo</p>
    <Adds />
    <Play bind:card />
</div>


<style>
    h2{
        margin-top: 50px;
        font-size: 2.5em;
    }
    p{
        margin-top: 25px;
        padding: 0 10px 0 10px;
        font-size: 1.6em;
    }
    @media(max-width: 450px){
        h2{
            margin: 0;
            padding: 10px 0 10px 0;
            font-size: 1.5em;
        }
        p{
            margin: 0;
            margin-bottom: 20px;
            font-size: 1em;
        }
    }
</style>