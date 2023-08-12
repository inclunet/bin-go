<script>
    import Play from "$lib/Play.svelte";

export let data;
export let card = data;
</script>
<h2>Rodada #{card.Round}</h2>
<p>Para jogar clique no botão "Jogar" a baixo</p>

<Play bind:card />