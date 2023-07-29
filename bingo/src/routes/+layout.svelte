<header>
    <h1>Inclubingo</h1>
</header>
<main>
    <slot></slot>
</main>
<footer>
    <p>Jogo de Bingo acessível e inclusivo para pessoas com deficiência.</p>
</footer>

<style>
    h1, p{
        text-align: center;
    }
    h1{
        width: 100%;
        padding-top: 10px;
        padding-bottom: 5px;
        margin-bottom: 10px;
        font-size: 4.4em;
        background-color: #2b7ef4;
        color: #000;
    }
    p{
        margin-top: 30px;
        padding: 0 10px 0 10px;
        font-size: 1.8em;
        line-height: 1.7em;
    }
    @media(max-width: 450px){
        h1{
            font-size: 2.5em;
        }
    }
</style>