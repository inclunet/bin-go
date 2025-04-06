<script>
    import Footer from "$lib/Footer.svelte";
    import ItemSocial from "$lib/menu/ItemSocial.svelte";
    import MenuHeader from "$lib/menu/MenuHeader.svelte";
    import MenuItem from "$lib/menu/MenuItem.svelte";
    import MenuSocial from "$lib/menu/MenuSocial.svelte";
</script>

<header>
    <div class="top-header_container">
        <a href="/"
            ><img class="img_logo" src="/img/logo.png" alt="Inclunet" /></a
        >
        <h1 class="text-center">
            <span> Inclubraille </span>
            <span
                ><img
                    class="img_inclubraille"
                    src="/img/inclubraille-braille.png"
                    alt="Nome Inclubraille na escrita Braille"
                /></span
            >
        </h1>
    </div>

    <div>
        <MenuHeader title="IncluBraille">
            <MenuItem href="/">Portal Jogos</MenuItem>
            <MenuItem href="/braille">Início</MenuItem>
            <MenuItem href="/braille/new">Jogar</MenuItem>
            <MenuItem href="/braille/instructions">Como Jogar</MenuItem>
            <MenuItem href="/braille/terms">Termo de Uso</MenuItem>
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
<Footer />

<style>
    header {
        display: flex;
        flex-direction: column;
        background-color: var(--primary-color);
        background-image: url("/img/banner_braille.png");
        background-repeat: no-repeat;
        background-size: 75%;
        background-position: top 10% center;
    }
    header h1 {
        width: 70rem;
        font-size: 8rem;
        margin-left: 15%;
        background: var(--secondary-color);
        color: var(--quinary-color);
        text-shadow:
            -0.1rem -0.1rem 0.3rem var(--quaternary-color),
            -0.1rem 0.1rem 0.3rem var(--quaternary-color),
            0.1rem -0.1rem 0 var(--quaternary-color),
            0.1rem 0 0 var(--quaternary-color);
        border-radius: 1rem;
    }
    .top-header_container {
        display: flex;
        margin: 3rem 0 3.5rem 0;
    }

    .content {
        flex: 1 0 auto;
    }

    a {
        color: var(--quaternary-color);
    }

    @media (max-width: 1182px) {
        header {
            background-size: 80vw;
        }
    }

    @media (max-width: 1040px) {
        header {
            background-size: 85vw;
            background-position: 80% 10%;
        }
    }
    @media (max-width: 867px) {
        header {
            background-size: 90rem;
            background-position: -250% 20%;
        }
    }

    @media (max-width: 794px) {
        header h1 {
            font-size: 6.4rem;
        }
        header .img_logo {
            width: 100%;
        }
    }
    @media (max-width: 768px) {
        header {
            background-position: -70% 20%;
        }
    }
    @media (max-width: 743px) {
        header {
            background-image: none;
        }
        header h1 {
            margin-left: 0;
        }
        header .img_logo {
            width: 15vw;
        }
        header .img_inclubraille {
            width: 60vw;
        }
    }

    @media (max-width: 685px) {
        header {
            padding: 0;
        }
        header h1 {
            font-size: 5.9rem;
        }
    }
    @media (max-width: 650px) {
        header h1 {
            width: 40rem;
            margin: 0;
            margin-left: 5rem;
            padding: 0;
        }
        header .img_logo {
            width: 110%;
        }
        .img_inclubraille {
            width: 40rem;
        }
    }

    @media (max-width: 571px) {
        header {
            justify-content: space-between;
        }

        header h1 {
            width: 35rem;
            margin: 0;
            padding: 0;
        }
        header .img_logo {
            width: 80%;
        }
    }
    @media (max-width: 520px) {
        .img_inclubraille {
            width: 100%;
        }
    }

    @media (max-width: 450px) {
        header {
            flex-direction: row;
            padding: 0.3rem 1rem;
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
            font-size: 4rem;
        }
        .img_inclubraille {
            display: none;
        }
    }
</style>
