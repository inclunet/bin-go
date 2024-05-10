<script>
    import { goto } from "$app/navigation";
    import PageTitle from "$lib/PageTitle.svelte";

    const handleStartLearningNow = () => {
        goto(`/braille/new`);
    };
</script>

<PageTitle title="Braille Personal Trainer" game="Inclubraille" />

<div class="container">
    <header class="my-5">
        <h2>Braille Personal Trainer</h2>
    </header>

    <section>
        <p>
            O Braille Personal Trainer é um aplicativo que ajuda a treinar a
            escrita em Braille. O aplicativo é composto por exercícios de
            escrita em Braille e de escrita em tinta. O aplicativo também possui
            um teclado Braille virtual para auxiliar na escrita.
        </p>

        <p>
            O aplicativo é composto por aulas de escrita em Braille. Cada aula é
            composta por exercícios de escrita em Braille e de escrita em tinta.
            O aplicativo também possui um teclado Braille virtual para auxiliar
            na escrita.
        </p>
    </section>

    <section>
        <button
            class="btn button-color text-black"
            on:click={handleStartLearningNow}
        >
            <strong>Comece a aprender Agora</strong>
        </button>
    </section>
</div>

<style>
    @import "style.css";

    h2:nth-child(2n) {
        font-size: 1.6em;
    }

    section p,
    section button {
        font-size: 1.2em;
    }

    p {
        line-height: 32px;
    }

    h3 {
        font-size: 1.4em;
    }

    .term {
        padding: 3px 5px 3px 5px;
        margin: 0;
        font-size: 1em;
    }

    @media (max-width: 573px) {
        .container {
            padding: 0 20px 0 20px;
        }
        header h2 {
            margin-top: -20px;
            line-height: 40px;
            text-align: center;
        }
        h3 {
            line-height: 40px;
        }
        header p:nth-child(n + 3):nth-child(-n + 4) {
            text-align: justify;
        }
        .row {
            padding: 0px;
        }
        .card {
            border-radius: 10px;
            box-shadow: 5px 5px 5px #e2e2e2;
        }

        .main-features-img {
            margin-left: 10px;
        }
        p {
            line-height: 37px;
        }
        section p {
            text-align: justify;
        }
    }
</style>
