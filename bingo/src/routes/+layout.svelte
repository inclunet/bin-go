<header class="d-flex flex-row">
    <a href="/"><img src="/img/logo.png" alt="Logotipo da Inclunet"></a>
    <h1 class="text-center ">Inclubingo</h1>
</header>
<main>
    <slot></slot>
</main>
<footer>
    <p class="text-center">
        <strong>
            Jogo de Bingo acessível e inclusivo para pessoas com deficiência.
        </strong>
    </p>
</footer>

<style>
   header{
       background-color: #2b7ef4;
       padding: 50px 50px 20px 50px;
       align-items: center;
   }
   img{
   }
    header h1{
        width: 100%;
        font-size: 4.4em;
        color: #000;
    }
    p{
        margin-top: 30px;
        padding: 0 10px 0 10px;
        font-size: 1.6em;
    }
    @media(max-width: 450px){
        header{
            padding: 0 10px 5px 10px;
            justify-content: space-between;
            /* align-items: baseline; */
        }
        header img{
            /* border: 1px solid black; */
            width: 80%;
        }
        header h1{
            /* border: 1px solid black; */
            height: auto;
            margin-top: 15px;
            /* padding: 15px 0 0 0; */
            font-size: 2.5em;
        }
        p{
            font-size: 1.3em;
        }
    }
</style>