<header class="d-flex flex-row">
    <a href="/braille"><img src="/img/logo.png" alt="Inclunet" /></a>
    <h1 class="text-center">Inclubraille</h1>

    <nav class="navbar">
        <button
            class="navbar-toggler button-menu navbar-dark"
            type="button"
            data-bs-toggle="offcanvas"
            data-bs-target="#offcanvasDarkNavbar"
            aria-controls="offcanvasDarkNavbar"
            aria-label="Toggle navigation"
        >
            <span class="navbar-toggler-icon"></span>
        </button>
        <div class="container-fluid">
            <div
                class="offcanvas offcanvas-end text-bg-dark"
                tabindex="-1"
                id="offcanvasDarkNavbar"
                aria-labelledby="offcanvasDarkNavbarLabel"
            >
                <div class="offcanvas-header">
                    <h5 class="offcanvas-title" id="offcanvasDarkNavbarLabel">
                        Inclubraille
                    </h5>
                    <button
                        type="button"
                        class="btn-close btn-close-white"
                        data-bs-dismiss="offcanvas"
                        aria-label="Close"
                    ></button>
                </div>
                <div class="offcanvas-body">
                    <ul class="navbar-nav justify-content-end flex-grow-1 pe-3">
                        <li
                            class="nav-item"
                            data-bs-dismiss="offcanvas"
                            aria-label="Close"
                        >
                            <a
                                class="nav-link active"
                                aria-current="page"
                                href="/braille">Início</a
                            >
                        </li>
                        <li
                            class="nav-item"
                            data-bs-dismiss="offcanvas"
                            aria-label="Close"
                        >
                            <a class="nav-link" href="/braille/new"
                                >Ir para o jogo</a
                            >
                        </li>
                        <li
                            class="nav-item"
                            data-bs-dismiss="offcanvas"
                            aria-label="Close"
                        >
                            <a class="nav-link" href="/braille/instructions"
                                >Como Jogar</a
                            >
                        </li>
                        <li
                            class="nav-item"
                            data-bs-dismiss="offcanvas"
                            aria-label="Close"
                        >
                            <a class="nav-link" href="/braille/terms"
                                >Termo de uso</a
                            >
                        </li>
                    </ul>
                    <hr />
                    <section class="section-social-medias-footer mt-4">
                        <ul class="list-group flex-row">
                            <li class="mx-3">
                                <a
                                    href="https://www.facebook.com/CanalInclunet"
                                    target="_blank"
                                    ><img
                                        src="/img/facebook.png"
                                        alt="Facebook"
                                        class="social-media-img list-group-item list-group-item-action"
                                    /></a
                                >
                            </li>
                            <li class="mx-3">
                                <a
                                    href="https://www.youtube.com/@Inclunet"
                                    target="_blank"
                                    ><img
                                        src="/img/youtube.png"
                                        alt="YouTube"
                                        class="social-media-img list-group-item list-group-item-action"
                                    /></a
                                >
                            </li>
                            <li class="mx-3">
                                <a
                                    href="https://www.instagram.com/inclunet/"
                                    target="_blank"
                                    ><img
                                        src="/img/instagram.png"
                                        alt="Instagram"
                                        class="social-media-img list-group-item list-group-item-action"
                                    /></a
                                >
                            </li>
                            <li class="mx-3">
                                <a
                                    href="https://www.linkedin.com/company/inclunet/"
                                    target="_blank"
                                    ><img
                                        src="/img/linkedin.png"
                                        alt="Linkedin"
                                        class="social-media-img list-group-item list-group-item-action"
                                    /></a
                                >
                            </li>
                            <li class="mx-3">
                                <a
                                    href="https://www.twitter.com/inclunet"
                                    target="_blank"
                                    ><img
                                        src="/img/twitter.png"
                                        alt="Twitter"
                                        class="social-media-img list-group-item list-group-item-action"
                                    /></a
                                >
                            </li>
                        </ul>
                    </section>
                </div>
            </div>
        </div>
    </nav>
</header>
<nav class="navbar menu-navbar">
    <ul
        class="container-nav-inner navbar-nav navbar-ul-primary d-flex flex-row"
    >
        <li class="nav-item">
            <a href="/braille" class="nav-link navbar-brand">Início</a>
        </li>
        <li class="nav-item">
            <a href="/braille/new" class="nav-link navbar-brand">Jogar</a>
        </li>
        <li class="nav-item">
            <a href="/braille/instructions" class="nav-link navbar-brand"
                >Como Jogar</a
            >
        </li>
        <li class="nav-item">
            <a href="/braille/terms" class="nav-link navbar-brand"
                >Termo de Uso</a
            >
        </li>
    </ul>
</nav>
<main class="content">
    <slot></slot>
</main>
<footer>
    <p class="text-center">
        <strong> Aprenda Braille brincando. </strong>
    </p>
    <hr />
    <section class="content-footer">
        <section class="section-logo-footer">
            <a href="/braille"
                ><img
                    src="/img/logo.png"
                    alt="Inclunet"
                    class="logo-footer"
                /></a
            >
        </section>
        <section class="section-useful-links-footer">
            <ul class="list-group">
                <li>
                    <a
                        href="https://inclunet.com.br/quem-somos/"
                        class="useful-links list-group-item list-group-item-action"
                        >SOBRE A INCLUNET</a
                    >
                </li>
                <li>
                    <a
                        href="."
                        class="useful-links list-group-item list-group-item-action"
                        >DÚVIDAS FREQUENTES</a
                    >
                </li>
                <li>
                    <a
                        href="."
                        class="useful-links list-group-item list-group-item-action"
                        >SUGIRA UMA FUNCIONALIDADE</a
                    >
                </li>
            </ul>
        </section>
        <section class="section-blog-footer">
            <ul class="list-group">
                <li>
                    <a
                        href="https://inclunet.com.br/"
                        class="useful-links list-group-item list-group-item-action"
                        target="_blank">BLOG DA INCLUNET</a
                    >
                </li>
            </ul>
        </section>
        <section class="section-social-medias-footer">
            <ul class="d-flex flex-row">
                <li class="">
                    <a
                        href="https://www.facebook.com/CanalInclunet"
                        target="_blank"
                        ><img
                            src="/img/facebook.png"
                            alt="Facebook"
                            class="social-media-img"
                        /></a
                    >
                </li>
                <li class="">
                    <a href="https://www.youtube.com/@Inclunet" target="_blank"
                        ><img
                            src="/img/youtube.png"
                            alt="YouTube"
                            class="social-media-img"
                        /></a
                    >
                </li>
                <li class="">
                    <a
                        href="https://www.instagram.com/inclunet/"
                        target="_blank"
                        ><img
                            src="/img/instagram.png"
                            alt="Instagram"
                            class="social-media-img"
                        /></a
                    >
                </li>
                <li class="">
                    <a
                        href="https://www.linkedin.com/company/inclunet/"
                        target="_blank"
                        ><img
                            src="/img/linkedin.png"
                            alt="Linkedin"
                            class="social-media-img"
                        /></a
                    >
                </li>
                <li class="">
                    <a href="https://www.twitter.com/inclunet" target="_blank"
                        ><img
                            src="/img/twitter.png"
                            alt="Twitter"
                            class="social-media-img"
                        /></a
                    >
                </li>
            </ul>
        </section>
    </section>
    <hr />
    <section class="copyright-footer">
        <p>Copyright © Inclunet 2023</p>
    </section>
</footer>

<style>
    :root {
        --primary-color: #2b7ef4;
        --secondary-color: #2b7ff4ab;
        --tertiary-color: #982a35;
        --quaternary-color: #fff;
        --quinary-color: #000;
        --senary-color: #1d1d1d;
        --setenary-color: #5799f5;
        --octonary-color: #006aff;
    }
    .container-nav-inner,
    .menu-navbar,
    .menu-navbar a,
    .copyright-footer p,
    .social-media-img,
    ul,
    ul.list-group,
    .container-nav-inner ul li a,
    .social-medias-nav-img .copyright-footer p,
    .social-media-img,
    ul,
    ul.list-group {
        margin: 0;
    }

    header {
        background-color: var(--primary-color);
        background-image: url("/img/fundo.png");
        background-repeat: repeat;
        background-size: 50%;
        padding: 25px 50px 10px 50px;
        align-items: center;
    }
    header h1 {
        width: 600px;
        font-size: 5em;
        margin-left: 15%;
        background: var(--secondary-color);
        color: var(--quinary-color);
        text-shadow:
            -1px -1px 3px var(--quaternary-color),
            -1px 1px 3px var(--quaternary-color),
            1px -1px 0px var(--quaternary-color),
            1px 0px 0px var(--quaternary-color);
        border-radius: 10px;
    }
    header .navbar {
        display: none;
    }
    .menu-navbar {
        background-color: var(--primary-color);
        border-top: 1px solid var(--quaternary-color);
        display: flex;
        justify-content: center;
    }
    .container-nav-inner {
        height: 48px;
        width: 80%;
        display: flex;
        justify-content: space-between;
        align-items: center;
        flex-wrap: wrap;
    }
    .menu-navbar a {
        height: 48px;
        line-height: 48px;
        padding: 0 30px;
        color: var(--quinary-color);
        font-weight: 600;
    }
    .menu-navbar a:hover {
        background-color: var(--setenary-color);
        color: var(--quaternary-color);
        text-shadow:
            -1px -1px 1px var(--senary-color),
            -1px 1px 1px var(--senary-color),
            1px -1px 0px var(--senary-color),
            1px 0px 0px var(--senary-color);
    }

    .menu-navbar a:active {
        background-color: var(--octonary-color);
        color: var(--senary-color);
    }

    .section-social-medias-nav ul li {
        margin: 0 8px 0 8px;
    }

    .content {
        flex: 1 0 auto;
    }
    p {
        margin-top: 30px;
        padding: 0 10px 0 10px;
        font-size: 1.6em;
    }
    .copyright-footer p {
        font-size: 1.3em;
    }
    footer {
        display: flex;
        flex-direction: column;
        align-items: center;
        flex-shrink: 0;
        width: 100%;
        min-height: 30vh;
        background-color: var(--tertiary-color);
        color: var(--quaternary-color);
        margin-top: 45px;
        padding-bottom: 30px;
    }
    footer hr {
        margin: 30px 0 30px 0;
        color: var(--quaternary-color);
        width: 95%;
    }
    .content-footer,
    footer p {
        display: flex;
        flex-direction: row;
    }
    .content-footer {
        width: 100%;
        justify-content: space-around;
        align-items: start;
    }
    .section-social-medias-footer ul li {
        margin: 0 8px 0 8px;
    }
    .social-media-img {
        border-radius: 5px;
        border: 2px solid var(--quaternary-color);
    }
    ul.list-group {
        display: flex;
        justify-content: space-between;
    }
    a {
        color: var(--quaternary-color);
    }
    .useful-links {
        border: none;
        background-color: var(--tertiary-color);
    }
    @media (max-width: 838px) {
        .menu-navbar a {
            padding: 0 10px;
        }
        .container-nav-inner {
            padding: 0 25px 0 25px;
        }
        .section-social-medias-footer ul li {
            margin: 0 3px 0 3px;
        }
    }
    @media (max-width: 794px) {
        header h1 {
            font-size: 4em;
        }
        header img {
            width: 100%;
        }
        .content-footer {
            width: 95%;
            margin: 0;
            padding: 0;
        }
        .section-logo-footer {
            display: flex;
            justify-content: center;
            margin: 0;
        }
        .section-logo-footer img {
            margin: 0;
            width: 100px;
        }
        .section-social-medias-nav ul li {
            margin: 0 3px 0 3px;
        }
    }
    @media (max-width: 745px) {
        .container-nav-inner {
            padding: 0 10px 0 10px;
        }
        .content-footer {
            flex-wrap: wrap;
        }
        .section-social-medias-footer {
            margin-top: 50px;
        }
        .section-social-medias-nav ul li {
            margin-left: 3px;
        }
    }
    @media (max-width: 685px) {
        header {
            padding: 25px 25px 10px 25px;
        }
        header h1 {
            font-size: 3.7em;
        }
    }
    @media (max-width: 650px) {
        .menu-navbar a {
            padding: 7px;
        }
    }
    @media (max-width: 640px) {
        header {
            padding: 25px 15px 10px 15px;
        }
        .container-nav-inner {
            padding: 0;
        }
        .navbar-ul-primary {
            width: 100%;
            justify-content: center;
        }
        .container-nav-inner .navbar-ul-primary a {
            padding: 0 20px 0 20px;
        }
    }
    @media (max-width: 571px) {
        header {
            justify-content: space-between;
        }

        header h1 {
            width: 350px;
            margin: 0;
            padding: 0;
        }
        header img {
            width: 80%;
        }
        .container-nav-inner .navbar-ul-primary a {
            padding: 0 12px 0 12px;
        }
    }
    @media (max-width: 565px) {
        .section-blog-footer {
            margin-top: 50px;
        }
    }
    @media (max-width: 450px) {
        header {
            padding: 0 10px 5px 10px;
            justify-content: space-between;
            background-image: none;
        }
        header a {
            position: 0;
            display: flex;
            justify-content: center;
        }
        header img {
            margin: 0;
            padding: 0;
            width: 80%;
        }
        header h1 {
            height: auto;
            width: 5%;
            margin: 0;
            width: 100%;
            font-size: 2.5em;
        }
        header .navbar {
            display: flex;
        }
        header .button-menu {
            background: var(--senary-color);
        }
        .menu-navbar {
            display: none;
        }
        p {
            font-size: 1.3em;
        }
        .copyright-footer p {
            margin-top: 20px;
            font-size: 1em;
        }
        footer,
        .content-footer {
            flex-direction: column;
            justify-content: center;
            align-items: center;
        }
        footer {
            min-height: 26vh;
        }
        .logo-footer {
            width: 90px;
        }
        footer hr {
            margin: 0;
        }
        .section-useful-links-footer ul li a {
            margin: 10px 0 10px 0;
            padding: 15px;
            text-align: center;
        }
        footer hr,
        .section-logo-footer,
        .section-useful-links-footer,
        .section-blog-footer,
        .section-social-medias-footer {
            margin-top: 40px;
        }
    }
</style>
