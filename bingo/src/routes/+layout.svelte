<header class="d-flex flex-row">
    <a href="/"><img src="/img/logo.png" alt="Logotipo da Inclunet"></a>
    <h1 class="text-center ">Inclubingo</h1>
</header>
<main class="content">
    <slot></slot>
</main>
<footer>
    <p class="text-center">
        <strong>
            Jogo de Bingo acessível e inclusivo para pessoas com deficiência.
        </strong>
    </p>
    <section>
        <section>
            <a href="/"><img src="/img/logo.png" alt="Logotipo da Inclunet"></a>
        </section>
        <section>
            <ul>
                <li>SOBRE A INCLUNET</li>
                <li>DÚVIDAS FREQUENTES</li>
                <li>SUGIRA UMA FUNCIONALIDADE</li>
            </ul>
        </section>
        <section>BLOG DA INCLUNET</section>
        <section class="">
            <a href="https://"><img src="/img/facebook.png" alt=""></a>
            <a href="https://"><img src="" alt=""></a>
            <a href="https://"><img src="/img/linkedin.png" alt="" class="social-media"></a>
        </section>
    </section>
</footer>

<style>
   header{
       background-color: #2b7ef4;
       padding: 50px 50px 20px 50px;
       align-items: center;
   }
    header h1{
        width: 100%;
        font-size: 4.4em;
        color: #000;
    }
    .content{
        flex: 1 0 auto;
    }
    p{
        margin-top: 30px;
        padding: 0 10px 0 10px;
        font-size: 1.6em;
    }
    footer{
        flex-shrink: 0;
        width: 100%;
        min-height: 30vh;
        background-color: #982a35;
        margin-top: 45px;
    }
    .social-media{
        width: 36px;
        height: 36px;
        background-color: #fff;
    }
    
    @media(max-width: 450px){
        header{
            padding: 0 10px 5px 10px;
            justify-content: space-between;

        }
        header img{
            width: 80%;
        }
        header h1{
            height: auto;
            margin-top: 15px;
            font-size: 2.5em;
        }
        p{
            font-size: 1.3em;
        }
        footer{
            min-height: 26vh;
        }
    }
</style>