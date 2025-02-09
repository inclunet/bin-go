<script>
    import { goto } from "$app/navigation";
    import PageTitle from "$lib/PageTitle.svelte";
    import { callApi } from "$lib/api";
    import { braille } from "$lib/braille";

    const handleStartNewBraillePlayer = async () => {
        $braille = await callApi($braille, `/api/braille/new`, "GET");

        if ($braille.Player >= 0) {
            goto(`/braille/${$braille.Player}`);
        }
    };
</script>

<PageTitle title="Novo Jogo - Braille Personal Trainer" game="Inclubraille" />

<div class="container">
    <header class="my-5">
        <h2>Novo Jogo - Braille Personal Trainer</h2>
    </header>

    <section>
        <button class="btn button-color" on:click={handleStartNewBraillePlayer}>
            <strong>Começar treino agora!</strong>
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

    .container {
        min-height: 60vh;
    }
    h2 {
        font-size: 2.8rem;
    }
    section button {
        margin-top: 5rem;
        font-size: 2rem;
        padding: 1rem 2rem;
    }

    p {
        line-height: 32px;
    }

    .button-color {
        background-color: var(--primary-button-color);
        color: var(--black);
    }
    .button-color:active {
        background-color: var(--secondary-button-color);
        color: var(--white);
        border: 1px solid var(--black);
    }

    .term {
        padding: 3px 5px 3px 5px;
        margin: 0;
        font-size: 1em;
    }

    @media (hover: hover) {
        .button-color:hover {
            background-color: var(--secondary-button-color);
            color: var(--white);
        }
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
            line-height: 37px;
        }
        section p {
            text-align: justify;
        }
    }
</style>
