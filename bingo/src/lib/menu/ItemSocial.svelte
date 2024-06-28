<script>
    import MediaQuery from "$lib/MediaQuery.svelte";
    export let href;
    export let src;
    export let alt;
</script>

<MediaQuery query="(max-width: 480px)" let:matches>
    {#if matches}
        <li class="mx-3">
            <a {href} target="_blank"
                ><img
                    {src}
                    {alt}
                    class="social-media-img list-group-item list-group-item-action"
                /></a
            >
        </li>
    {/if}
</MediaQuery>

<style>
    a {
        border-radius: 0.3rem;
    }
    img {
        padding: 0.16rem;
    }
</style>
