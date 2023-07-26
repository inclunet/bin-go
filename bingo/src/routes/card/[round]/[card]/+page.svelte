<script>
    import Card from "./Card.svelte";

    export let data;
    
</script>

<div>
    <h2>Cartela de Bingo #{data.Card} rodada #{data.Round}</h2>
    <Card bind:card={data} />
</div>
