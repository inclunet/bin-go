<footer>
    <p class="text-center">
        <strong> <slot></slot> </strong>
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

        font-size: 62.5%;
    }
    * {
        margin: 0;
        padding: 0;
        box-sizing: border-box;
    }
    .section-social-medias-nav ul li {
        margin: 0 0.8rem 0 0.8rem;
    }
    p {
        margin-top: 2rem;
        padding: 0 1rem 0 1rem;
        font-size: 2.5rem;
    }
    .copyright-footer p {
        margin: 0;
        font-size: 2rem;
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
        margin-top: 4.5rem;
        padding-bottom: 3rem;
    }
    footer hr {
        margin: 1.5rem 0;
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
    .section-useful-links-footer li {
        line-height: 5rem;
    }
    .section-social-medias-footer ul li {
        margin: 0 0.8rem 0 0.8rem;
    }
    .social-media-img {
        border-radius: 5px;
        border: 0.2rem solid var(--quaternary-color);
    }
    ul.list-group {
        display: flex;
        justify-content: space-between;
    }
    .useful-links {
        border: none;
        background-color: var(--tertiary-color);
    }
    a {
        color: var(--quaternary-color);
        font-size: 1.8rem;
    }
    @media (max-width: 838px) {
        .section-social-medias-footer ul li {
            margin: 0 0.3rem 0 0.3rem;
        }
    }
    @media (max-width: 794px) {
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
            width: 10rem;
        }
        .section-social-medias-nav ul li {
            margin: 0 0.3rem 0 0.3rem;
        }
    }
    @media (max-width: 745px) {
        .content-footer {
            flex-wrap: wrap;
        }
        .section-social-medias-footer {
            margin-top: 5rem;
        }
        .section-social-medias-nav ul li {
            margin-left: 3px;
        }
    }
    @media (max-width: 565px) {
        .section-blog-footer {
            margin-top: 5rem;
        }
    }
    @media (max-width: 450px) {
        p {
            font-size: 2rem;
        }
        .copyright-footer p {
            margin-top: 2rem;
            font-size: 1.6rem;
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
            width: 9rem;
        }
        footer hr {
            margin: 0;
        }
        .section-useful-links-footer ul li a {
            margin: 1rem 0 1rem 0;
            padding: 1.5rem;
            text-align: center;
        }
        footer hr,
        .section-logo-footer,
        .section-useful-links-footer,
        .section-blog-footer,
        .section-social-medias-footer {
            margin-top: 4rem;
        }
        .section-useful-links-footer li {
            line-height: 1rem;
        }
    }
</style>
