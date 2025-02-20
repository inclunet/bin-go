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
    .container {
        min-height: 60vh;
    }
</style>
