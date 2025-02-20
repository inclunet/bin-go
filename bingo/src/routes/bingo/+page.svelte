<script>
    import LandingPage from "$lib/LandingPage.svelte";
    import PageTitle from "$lib/PageTitle.svelte";
</script>

<PageTitle title="Início" game="Inclubingo" />
<div class="m-3 me-5 term-use">
    <a href="/terms" class="btn btn-warning btn-lg term-link">Termos de uso</a>
</div>

<LandingPage />

<style>
    .term-link {
        font-size: 2.2rem;
        font-weight: 500;
    }
    @media (max-width: 450px) {
        .term-use {
            display: none;
        }
    }
</style>
