<script>
    import PageTitle from "$lib/PageTitle.svelte";
    import { onMount } from "svelte";
    import { goto } from "$app/navigation";

    export let data;

    onMount(() => goto(`/bingo/${data.Round}/new`));
</script>

<PageTitle title="Entrar na rodada" game="Inclubingo" />

<div class="container-fluid d-flex align-items-center flex-column">
    <p class="text-center">Redirecionando para criar sua cartela…</p>
</div>

<style>
    :root {
        font-size: 62.5%;
    }
    p {
        padding: 0 1rem 0 1rem;
        font-size: 1.8rem;
    }

    @media (max-width: 450px) {
        p {
            margin: 0;
            margin-bottom: 2rem;
        }
    }
</style>
