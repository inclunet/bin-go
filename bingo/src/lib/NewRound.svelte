<script>
    import Button from "./Button.svelte";

    export let disabled = false;
    export let ariaBusy = false;
</script>

<Button on:click {disabled} {ariaBusy}>
    {ariaBusy ? "Criando rodada…" : "Nova Rodada"}
</Button>
