<script>
    import MediaQuery from "$lib/MediaQuery.svelte";
</script>

<MediaQuery query="(max-width: 480px)" let:matches>
    {#if matches}
        <hr />
        <section class="section-social-medias-footer mt-4">
            <ul class="list-group flex-row">
                <slot></slot>
            </ul>
        </section>
    {/if}
</MediaQuery>
