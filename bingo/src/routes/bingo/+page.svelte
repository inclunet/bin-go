<script>
    import LandingPage from "$lib/LandingPage.svelte";
    import PageTitle from "$lib/PageTitle.svelte";
</script>

<PageTitle title="Início" game="Inclubingo" />
<div class="m-3 me-5 term-use">
    <a href="/terms" class="btn btn-warning btn-lg">Termos de uso</a>
</div>

<LandingPage />

<style>
    @media (max-width: 450px) {
        .term-use {
            display: none;
        }
    }
</style>
