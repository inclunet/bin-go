<script>
    import Footer from "$lib/Footer.svelte";
    import ItemSocial from "$lib/menu/ItemSocial.svelte";
    import MenuHeader from "$lib/menu/MenuHeader.svelte";
    import MenuItem from "$lib/menu/MenuItem.svelte";
    import MenuSocial from "$lib/menu/MenuSocial.svelte";
</script>

<header>
    <div class="top-header_container">
        <a href="/"><img src="/img/logo.png" alt="Inclunet" /></a>
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
<Footer />

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

    header {
        display: flex;
        flex-direction: column;
        background-color: var(--primary-color);
        background-image: url("/img/fundo.png");
        background-repeat: repeat;
        background-size: 50%;
    }
    header h1 {
        width: 60rem;
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

    @media (max-width: 794px) {
        header h1 {
            font-size: 6.4rem;
        }
        header img {
            width: 100%;
        }
    }

    @media (max-width: 685px) {
        header {
            padding: 2.5rem 2.5rem 1rem 2.5rem;
        }
        header h1 {
            font-size: 5.9rem;
        }
    }
    @media (max-width: 640px) {
        header {
            padding: 2.5rem 1.5rem 1rem 1.5rem;
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
        header img {
            width: 80%;
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
    }
</style>
