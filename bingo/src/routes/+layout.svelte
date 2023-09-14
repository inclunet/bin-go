<header class="d-flex flex-row">
    <a href="/"><img src="/img/logo.png" alt="Inclunet"></a>
    <h1 class="text-center ">Inclubingo</h1>
    
    <nav class="navbar">
        <button class="navbar-toggler button-menu navbar-dark" type="button" data-bs-toggle="offcanvas" data-bs-target="#offcanvasDarkNavbar" aria-controls="offcanvasDarkNavbar" aria-label="Toggle navigation">
          <span class="navbar-toggler-icon"></span>
        </button>
        <div class="container-fluid">
          <div class="offcanvas offcanvas-end text-bg-dark" tabindex="-1" id="offcanvasDarkNavbar" aria-labelledby="offcanvasDarkNavbarLabel">
            <div class="offcanvas-header">
              <h5 class="offcanvas-title" id="offcanvasDarkNavbarLabel">Inclubingo</h5>
              <button type="button" class="btn-close btn-close-white" data-bs-dismiss="offcanvas" aria-label="Close"></button>
            </div>
            <div class="offcanvas-body">
              <ul class="navbar-nav justify-content-end flex-grow-1 pe-3">
                <li class="nav-item" data-bs-dismiss="offcanvas" aria-label="Close">
                  <a class="nav-link active" aria-current="page" href="/">Home</a>
                </li>
                <li class="nav-item" data-bs-dismiss="offcanvas" aria-label="Close">
                  <a class="nav-link" href="/terms">Termo de uso</a>
                </li>
              </ul>
              <hr>
              <section class="section-social-medias-footer mt-4">
                <ul class="list-group flex-row">
                    <li class="mx-3"><a href="https://www.facebook.com/CanalInclunet" target="_blank"><img src="/img/facebook.png" alt="Facebook" class="social-media-img list-group-item list-group-item-action"></a></li>
                    <li class="mx-3"><a href="https://www.youtube.com/@Inclunet" target="_blank"><img src="/img/youtube.png" alt="YouTube" class="social-media-img list-group-item list-group-item-action"></a></li>
                    <li class="mx-3"><a href="https://www.instagram.com/inclunet/" target="_blank"><img src="/img/instagram.png" alt="Instagram" class="social-media-img list-group-item list-group-item-action"></a></li>
                    <li class="mx-3"><a href="https://www.linkedin.com/company/inclunet/" target="_blank"><img src="/img/linkedin.png" alt="Linkedin" class="social-media-img list-group-item list-group-item-action"></a></li>
                    <li class="mx-3"><a href="https://www.twitter.com/inclunet" target="_blank"><img src="/img/twitter.png" alt="Twitter" class="social-media-img list-group-item list-group-item-action"></a></li>
                </ul>
            </section>
            </div>
          </div>
        </div>
    </nav>
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
    <hr>
    <section class="content-footer">
        <section class="section-logo-footer">
            <a href="/"><img src="/img/logo.png" alt="Inclunet" class="logo-footer"></a>
        </section>
        <section class="section-useful-links-footer">
            <ul class="list-group">
                <li><a href="https://inclunet.com.br/quem-somos/" class="useful-links list-group-item list-group-item-action">SOBRE A INCLUNET</a></li>
                <li><a href="#" class="useful-links list-group-item list-group-item-action">DÚVIDAS FREQUENTES</a></li>
                <li><a href="#" class="useful-links list-group-item list-group-item-action">SUGIRA UMA FUNCIONALIDADE</a></li>
            </ul>
        </section>
        <section class="section-blog-footer">
            <ul class="list-group">
                <li><a href="https://inclunet.com.br/" class="useful-links list-group-item list-group-item-action" target="_blank">BLOG DA INCLUNET</a></li>
            </ul>
        </section>
        <section class="section-social-medias-footer">
            <ul class="list-group flex-row">
                <li class="mx-3"><a href="https://www.facebook.com/CanalInclunet" target="_blank"><img src="/img/facebook.png" alt="Facebook" class="social-media-img list-group-item list-group-item-action"></a></li>
                <li class="mx-3"><a href="https://www.youtube.com/@Inclunet" target="_blank"><img src="/img/youtube.png" alt="YouTube" class="social-media-img list-group-item list-group-item-action"></a></li>
                <li class="mx-3"><a href="https://www.instagram.com/inclunet/" target="_blank"><img src="/img/instagram.png" alt="Instagram" class="social-media-img list-group-item list-group-item-action"></a></li>
                <li class="mx-3"><a href="https://www.linkedin.com/company/inclunet/" target="_blank"><img src="/img/linkedin.png" alt="Linkedin" class="social-media-img list-group-item list-group-item-action"></a></li>
                <li class="mx-3"><a href="https://www.twitter.com/inclunet" target="_blank"><img src="/img/twitter.png" alt="Twitter" class="social-media-img list-group-item list-group-item-action"></a></li>
            </ul>
        </section>
    </section>
    <hr>
    <section class="copyright-footer">
        <p>
            Copyright © Inclunet 2023
        </p>
        
    </section>
</footer>

<style>
    header{
        background-color: #2b7ef4;
        background-image: url("/img/fundo.png");
        background-repeat: repeat;
        background-size: 50%;
        padding: 50px 50px 20px 50px;
        align-items: center;
    }
    header h1{
        width: 50%;
        font-size: 6em;
        margin-left: 15%;
        background: #2b7ff4ab;
        color: #000;
        text-shadow: 
               -1px -1px 3px #FFF, 
               -1px 1px 3px #FFF,                    
                1px -1px 0px #FFF,                  
                1px 0px 0px #FFF;
        border-radius: 10px;
    }
    header .navbar{
        display: none;
    }
    .content{
        flex: 1 0 auto;
    }
    p{
        margin-top: 30px;
        padding: 0 10px 0 10px;
        font-size: 1.6em;
    }
    .copyright-footer p{
        margin: 0;
        padding: 0;
        font-size: 1.3em;
    }
    footer{
        display: flex;
        flex-direction: column;
        align-items: center;
        flex-shrink: 0;
        width: 100%;
        min-height: 30vh;
        background-color: #982a35;
        color: #fff;
        margin-top: 45px;
        padding-bottom: 30px;
    }
    footer hr{
        margin: 30px 0 30px 0;
        color: #fff;
        width: 95%;
    }
    .content-footer, footer p{
        display: flex;
        flex-direction: row;        
    }
    .content-footer{
        width: 100%;
        justify-content: space-around;
        align-items: start;
    }
    .social-media-img{
        margin: 0;
        padding: 0;
        border-radius: 5px;
        border: 2px solid #fff;
    }
    ul.list-group{
        display: flex;
        justify-content: space-between;
        margin: 0;
        padding: 0;
    }
    a{
        color: #fff;
    }
    .useful-links{
        border: none;
        background-color: #982a35;
    }
    
    @media(max-width: 450px){
        header{
            padding: 0 10px 5px 10px;
            justify-content: space-between;
            background-image: none;

        }
        header a {
            position: 0;
            display: flex;
            justify-content: center;
        }
        header img{
            margin: 0;
            padding: 0;
            width: 80%;
        }
        header h1{
            height: auto;
            width: 5%;
            margin: 0;
            width: 100%;
            font-size: 2.5em;
        }
        header .navbar{
            display: flex;
        }
        header .button-menu{
            background: #1d1d1d;
        }
        p{
            font-size: 1.3em;
        }
        .copyright-footer p{
            margin-top: 20px;
            font-size: 1em;
        }
        footer, .content-footer{
            flex-direction: column;
            justify-content: center;
            align-items: center;
        }
        footer{
            min-height: 26vh;
        }
        .logo-footer{
            width: 90px;
        }
        footer hr{
            margin: 0;
        }
        .section-useful-links-footer ul li a{
            margin: 10px 0 10px 0;
            padding: 15px;
            text-align: center;
        }
        footer hr,
        .section-logo-footer,
        .section-useful-links-footer, 
        .section-blog-footer,
        .section-social-medias-footer{
            margin-top: 40px
        }
    }
</style>