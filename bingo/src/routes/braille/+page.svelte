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
