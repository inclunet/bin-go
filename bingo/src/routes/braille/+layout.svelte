<script>
    import ItemSocial from "$lib/menu/ItemSocial.svelte";
    import MenuHeader from "$lib/menu/MenuHeader.svelte";
    import MenuItem from "$lib/menu/MenuItem.svelte";
    import MenuSocial from "$lib/menu/MenuSocial.svelte";
</script>

<header>
    <div class="top-header_container">
        <a href="/braille"><img src="/img/logo.png" alt="Inclunet" /></a>
        <h1 class="text-center">Inclubraille</h1>
    </div>

    <div>
        <MenuHeader title="IncluBraille">
            <MenuItem href="/braille">Início</MenuItem>
            <MenuItem href="/braille/new">Jogar</MenuItem>
            <MenuItem href="/braille/instructions">Como Jogar</MenuItem>
            <MenuItem href="/braille/terms">Termo de Uso</MenuItem>
            <MenuItem href="/braille/terms" mobile={true}>Qualquer</MenuItem>
            <MenuSocial>
                <ItemSocial
                    href="https://www.facebook.com/CanalInclunet"
                    src="/img/facebook.png"
                    alt="Facebook"
                />
                <ItemSocial
                    href="https://www.youtube.com/@Inclunet"
                    src="/img/youtube.png"
                    alt="YouTube"
                />
                <ItemSocial
                    href="https://www.instagram.com/inclunet/"
                    src="/img/instagram.png"
                    alt="Instagram"
                />
                <ItemSocial
                    href="https://www.linkedin.com/company/inclunet/"
                    src="/img/linkedin.png"
                    alt="Linkedin"
                />
                <ItemSocial
                    href="https://www.twitter.com/inclunet"
                    src="/img/twitter.png"
                    alt="Twitter"
                />
            </MenuSocial>
        </MenuHeader>
    </div>
</header>
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

    header {
        display: flex;
        flex-direction: column;
        background-color: var(--primary-color);
        background-image: url("/img/fundo.png");
        background-repeat: repeat;
        background-size: 50%;
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
    .top-header_container {
        display: flex;
        margin: 2.5em 0 3em 0;
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
    @media (max-width: 640px) {
        header {
            padding: 25px 15px 10px 15px;
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
    }
    @media (max-width: 565px) {
        .section-blog-footer {
            margin-top: 50px;
        }
    }
    @media (max-width: 450px) {
        header {
            flex-direction: row;
            padding: 0.2em 10px;
            align-items: center;
            background-image: none;
        }
        .top-header_container {
            margin: 0;
        }
        header a {
            display: none;
        }
        header h1 {
            width: 100%;
            font-size: 2.5em;
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
