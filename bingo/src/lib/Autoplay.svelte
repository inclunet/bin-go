<script>
    export let autoplay = false;
</script>

{#if autoplay}
    <button class="btn btn-success" on:click>
        <strong>Automático</strong>
    </button>
{:else}
    <button class="btn btn-warning" on:click> Manual </button>
{/if}

<style>
    button {
        padding: 10px;
        font-size: 1.3em;
        color: #000;
    }
    @media (max-width: 450px) {
        button {
            height: 48px;
            padding: 7px;
            font-size: 1em;
        }
    }
</style>
