<script>
    import { goto } from "$app/navigation";
    import PageTitle from "$lib/PageTitle.svelte";
    import { callApi } from "$lib/api";
    import { braille } from "$lib/braille";


    const handleStartNewBraillePlayer = async() => {
        $braille = await callApi($braille, `/api/braille/new`, "GET");

        if ($braille.Player >= 0) {
            goto(`/braille/${$braille.Player}`);
        }
    };
</script>

<PageTitle title="Novo Jogo - Braille Personal Trainer" />

<div>
    <h2>Novo Jogo - Braille Personal Trainer</h2>
    <button on:click={handleStartNewBraillePlayer}>Começar treino agora!</button>
</div>