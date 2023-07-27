<script>
    import { goto } from "$app/navigation";

async function new75Card() {
    let card = {Round: 0};
    const url = "http://localhost:8080/api/card/0/75";
        const response = await fetch(url);
        card = await response.json();
        goto("/card/" + card.Round);
}

    
</script>
<h1>Bem-Vindo ao Inclubingo</h1>
<p>Vamos Jogar Escolha a quantidade de bolinhas que serão sorteadas:</p>
<ul>
    <li><button on:click={new75Card} class="btn btn-primary">Bingo com 75 bolinhas</button></li>
</ul>
<style>
    h1, p{
        text-align: center;
    }
    h1{
        font-size: 3em;
        margin-bottom: 50px;
    }
    ul{
        margin: 0;
        padding: 0;
    }
    button{
        display: block;
        margin: 0 auto;
        font-size: 1.8em;
        margin-bottom: 50px;
        padding: 20px;
    }
    p{
        font-size: 1.8em;
        line-height: 1.7em;
        padding: 0 20px 0 20px;
        margin-bottom: 50px;
    }
</style>