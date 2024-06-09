<script>
    import { links } from "$lib/menuHeaderTopLinks.js";
    console.log("meus links: ", links);
</script>

<nav class="navbar menu-navbar">
    <ul
        class="container-nav-inner navbar-nav navbar-ul-primary d-flex flex-row"
    >
        {#each links as menuLink}
            <li class="nav-item">
                <a href={menuLink.link} class="nav-link navbar-brand"
                    >{menuLink.title}</a
                >
            </li>
        {/each}
    </ul>
</nav>

<style>
    :root {
        --primary-color: #2b7ef4;
        --quaternary-color: #fff;
        --quinary-color: #000;
        --senary-color: #1d1d1d;
        --setenary-color: #5799f5;
        --octonary-color: #006aff;
    }
    .container-nav-inner,
    .menu-navbar,
    .menu-navbar a,
    ul,
    .container-nav-inner ul li a,
    ul {
        margin: 0;
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

    .section-social-medias-footer ul li {
        margin: 0 8px 0 8px;
    }

    a {
        color: var(--quaternary-color);
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
        .section-social-medias-nav ul li {
            margin: 0 3px 0 3px;
        }
    }
    @media (max-width: 745px) {
        .container-nav-inner {
            padding: 0 10px 0 10px;
        }

        .section-social-medias-nav ul li {
            margin-left: 3px;
        }
    }
    @media (max-width: 685px) {
    }
    @media (max-width: 650px) {
        .menu-navbar a {
            padding: 7px;
        }
    }
    @media (max-width: 640px) {
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
        .container-nav-inner .navbar-ul-primary a {
            padding: 0 12px 0 12px;
        }
    }

    @media (max-width: 450px) {
        .menu-navbar {
            display: none;
        }

        .section-useful-links-footer ul li a {
            margin: 10px 0 10px 0;
            padding: 15px;
            text-align: center;
        }
    }
</style>
