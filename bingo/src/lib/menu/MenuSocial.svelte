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

<style>
    .container_social-medias-footer {
        position: relative;
        display: flex;
        justify-content: center;
    }
</style>
