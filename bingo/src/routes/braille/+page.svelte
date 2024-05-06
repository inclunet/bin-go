<script>
    import { goto } from "$app/navigation";
    import PageTitle from "$lib/PageTitle.svelte";

const handleStartLearningNow = () => {
goto(`/braille/new`);
};
</script>

<PageTitle title="Braille Personal Trainer" />

<h2>Braille Personal Trainer</h2>

<p>
    O Braille Personal Trainer é um aplicativo que ajuda a treinar a escrita em Braille. O aplicativo é composto por
    exercícios de escrita em Braille e de escrita em tinta. O aplicativo também possui um teclado Braille virtual para
    auxiliar na escrita.
</p>

<p>
    O aplicativo é composto por aulas de escrita em Braille. Cada aula é composta por exercícios de escrita em Braille e
    de escrita em tinta. O aplicativo também possui um teclado Braille virtual para auxiliar na escrita.
</p>

<button on:click={handleStartLearningNow}>Comece a aprender Agora</button>