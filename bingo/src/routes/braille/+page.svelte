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
        <button class="btn button-color" on:click={handleStartLearningNow}>
            <strong>Comece a aprender Agora</strong>
        </button>
    </section>
</div>

<style>
    :root {
        --primary-button-color: #2b7ef4;
        --secondary-button-color: #2868c2;
        --white: #fff;
        --black: #000;

        font-size: 62.5%;
    }

    h2 {
        font-size: 2.8rem;
    }

    section p,
    section button {
        font-size: 1.8rem;
    }

    p {
        line-height: 4rem;
    }

    .button-color {
        margin-top: 10px;
        background-color: var(--primary-button-color);
        color: var(--black);
    }

    .button-color:hover {
        background-color: var(--secondary-button-color);
        color: var(--white);
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

        p {
            line-height: 3.7rem;
        }
        section p {
            text-align: justify;
        }
    }
</style>
