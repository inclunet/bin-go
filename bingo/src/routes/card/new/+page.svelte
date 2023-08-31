<script>
    import NewRound from "$lib/NewRound.svelte";
    export let data;
    export let card = data;
</script>

<div class="container-fluid d-flex align-items-center flex-column">
    
    <h2 class="text-center">Bem-Vindo ao Inclubingo</h2>
    <p class="text-center">Vamos Jogar! Escolha a quantidade de bolinhas que serão sorteadas:</p>
    <NewRound bind:card/>
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