<script>
    import { goto } from "$app/navigation";
    import { page } from "$app/stores";
    import { newGame } from "./bingo";
    </script>

<h1>Bem-Vindo ao Inclubingo</h1>
<p>Vamos Jogar Escolha a quantidade de bolinhas que serão sorteadas:</p>
<ul>
    <li>
        <button on:click={newGame} class="btn btn-primary"
            >Bingo com 75 bolinhas</button
        >
    </li>
</ul>

<style>
    h1,
    p {
        text-align: center;
    }

    h1 {
        font-size: 3em;
        margin-bottom: 50px;
    }
    ul {
        margin: 0;
        padding: 0;
    }
    button {
        display: block;
        margin: 0 auto;
        font-size: 1.8em;
        margin-bottom: 50px;
        padding: 20px;
    }
    p {
        font-size: 1.8em;
        line-height: 1.7em;
        padding: 0 20px 0 20px;
        margin-bottom: 50px;
    }
</style>
