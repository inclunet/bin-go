<script>
    import MediaQuery from "../MediaQuery.svelte";
    import MenuDesktop from "./MenuDesktop.svelte";
    import MenuMobile from "./MenuMobile.svelte";
    export let title;
</script>

<MediaQuery query="(min-width: 481px)" let:matches>
    {#if matches}
        <MenuDesktop><slot></slot></MenuDesktop>
    {/if}
</MediaQuery>

<MediaQuery query="(max-width: 480px)" let:matches>
    {#if matches}
        <MenuMobile {title}><slot></slot></MenuMobile>
    {/if}
</MediaQuery>
