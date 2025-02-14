<script>
    import MediaQuery from "$lib/MediaQuery.svelte";
</script>

<MediaQuery query="(max-width: 480px)" let:matches>
    {#if matches}
        <hr />
        <section class="container_social-medias-footer mt-4">
            <ul class="list-group flex-row">
                <slot></slot>
            </ul>
        </section>
    {/if}
</MediaQuery>

<button
    type="button"
    class="btn-close btn-close-white hidden-area"
    data-bs-dismiss="offcanvas"
    aria-label="Fechar"
></button>

<style>
    .container_social-medias-footer {
        position: relative;
        display: flex;
        justify-content: center;
    }
    .hidden-area {
        position: absolute;
        width: 1px;
        height: 1px;
        padding: 0;
        margin: -1px;
        overflow: hidden;
        clip: rect(0, 0, 0, 0);
        border: 0;
    }
</style>
