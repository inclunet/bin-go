<header>
    <h1 class="text-center">Inclubingo</h1>
</header>
<main>
    <slot></slot>
</main>
<footer>
    <p class="text-center">Jogo de Bingo acessível e inclusivo para pessoas com deficiência.</p>
</footer>

<style>
   
    h1{
        padding-top: 50px;
        padding-bottom: 20px;
        font-size: 4.4em;
        background-color: #2b7ef4;
        color: #000;
    }
    p{
        margin-top: 30px;
        padding: 0 10px 0 10px;
        font-size: 1.6em;
    }
    @media(max-width: 450px){
        h1{
            padding: 0;
            font-size: 2.5em;
        }
        p{
            font-size: 1em;
        }
    }
</style>