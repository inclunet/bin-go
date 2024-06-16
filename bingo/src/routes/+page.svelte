<script>
    import CardPlay from "$lib/CardPlay.svelte";
    import Footer from "$lib/Footer.svelte";
</script>

<header>
    <img src="/img/inclunet.png" alt="Inclunet" />
</header>
<main class="container">
    <slot>
        <div class="container-card">
            <CardPlay
                href="/bingo"
                title="Jogo de Bingo Inclusivo"
                titleButton="Inclubingo"
            >
                Explore a emoção de um bingo acessível para todos no IncluBingo.
            </CardPlay>
            <CardPlay
                href="/braille"
                title="Braille Personal Trainer"
                titleButton="Inclubraille"
            >
                O Braille Personal Trainer te deixará expert na escrita em
                Braille.
            </CardPlay>
        </div>
    </slot>
</main>
<Footer>Aprenda Braille brincando.</Footer>

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
        justify-content: center;
    }
</style>
