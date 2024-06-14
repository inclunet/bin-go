<script>
    import MediaQuery from "$lib/MediaQuery.svelte";
    export let href;
    export let mobile = false;
</script>

<MediaQuery query="(min-width: 481px)" let:matches>
    {#if matches && !mobile}
        <li class="nav-item">
            <a {href} class="nav-link navbar-brand">
                <slot></slot>
            </a>
        </li>
    {/if}
</MediaQuery>

<MediaQuery query="(max-width: 480px)" let:matches>
    {#if matches}
        <li class="nav-item" data-bs-dismiss="offcanvas" aria-label="Close">
            <a class="nav-link" aria-current="page" {href}>
                <slot></slot>
            </a>
        </li>
    {/if}
</MediaQuery>

<style>
    :root {
        --primary-color: #2b7ef4;
        --secondary-color: #2b7ff4ab;
        --quaternary-color: #fff;
        --quinary-color: #000;
        --senary-color: #1d1d1d;
        --setenary-color: #5799f5;
        --octonary-color: #006aff;
    }
    .nav-link,
    .nav-item {
        margin: 0;
    }
    .nav-link {
        height: 48px;
        line-height: 48px;
        padding: 0 30px;
        color: var(--quinary-color);
        font-weight: 600;
    }
    .nav-link:hover {
        background-color: var(--setenary-color);
        color: var(--quaternary-color);
        text-shadow:
            -1px -1px 1px var(--senary-color),
            -1px 1px 1px var(--senary-color),
            1px -1px 0px var(--senary-color),
            1px 0px 0px var(--senary-color);
    }
    .nav-link:active {
        background-color: var(--octonary-color);
        color: var(--senary-color);
    }
    @media (max-width: 480px) {
        .nav-link {
            color: var(--quaternary-color);
        }
        .nav-item {
            text-align: center;
        }
    }
</style>
