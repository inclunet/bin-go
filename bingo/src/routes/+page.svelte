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
    <li><button on:click={new75Card}>Bingo com 75 bolinhas</button></li>
</ul>
