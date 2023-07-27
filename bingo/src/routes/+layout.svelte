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
        padding-top: 30px;
        padding-bottom: 30px;
        margin-bottom: 30px;
        font-size: 4.4em;
        background-color: #6aa5ff;
        color: white;
    }
    p{
        margin-top: 30px;
        padding: 0 10px 0 10px;
        font-size: 1.8em;
        line-height: 1.7em;
    }
    @media(max-width: 370px){
        h1{
            font-size: 3.5em;
        }
    }
</style>