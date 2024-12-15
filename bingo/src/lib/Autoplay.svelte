<script>
    import Button from "$lib/Button.svelte";
    export let autoplay = false;
</script>

{#if autoplay}
    <Button color="success" on:click>Automático</Button>
{:else}
    <Button color="warning" on:click>Manual</Button>
{/if}
