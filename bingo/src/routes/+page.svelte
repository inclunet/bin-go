<script>
    import PageTitle from "$lib/PageTitle.svelte";
    import CardPlay from "$lib/CardPlay.svelte";
    import Footer from "$lib/Footer.svelte";
</script>

<PageTitle title="Inclunet" game="Homepage" />
<header>
    <img src="/img/inclunet.png" alt="Inclunet" />
</header>
<main class="container">
    <slot>
        <div class="container-card">
            <CardPlay
                src="/img/chamada_card_inclubingo.png"
                alt="Inclubingo"
                href="/bingo"
                title="Jogo de Bingo Inclusivo"
            >
                Explore a emoção de um bingo acessível para todos no IncluBingo.
            </CardPlay>
            <CardPlay
                src="/img/chamada_card_inclubraille.png"
                alt="Inclubraille"
                href="/braille"
                title="Braille Personal Trainer"
            >
                O Braille Personal Trainer te deixará expert na escrita em
                Braille.
            </CardPlay>
        </div>
    </slot>
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

    header img {
        width: 100%;
    }
    .container {
        margin-top: 4rem;
    }
    .container-card {
        display: flex;
        flex-wrap: wrap;
        justify-content: space-around;
    }
</style>
