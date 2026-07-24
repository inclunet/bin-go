<script>
    import { onMount } from "svelte";

    export let compact = false;

    onMount(() => {
        (adsbygoogle = window.adsbygoogle || []).push({});
    });
</script>

<section class="container-adds" class:compact>
    <h2>{compact ? "Anúncio" : "Anúncios"}</h2>
    {#if !compact}
        <p>
            <em>
                Os Anúncios patrocinam o desenvolvimento e a manutenção deste
                jogo de Bingo acessível.
            </em>
        </p>
    {/if}

    <!-- Transição do Bingo -->
    <ins
        class="adsbygoogle"
        data-ad-client="ca-pub-6710397565046630"
        data-ad-slot="7204703685"
        data-ad-format="auto"
        data-full-width-responsive="true"
    ></ins>
</section>

<style>
    :root {
        font-size: 62.5%;
    }
    h2 {
        margin-top: 0;
        font-size: 2.8rem;
    }
    ins {
        display: flex;
        flex-wrap: wrap;
        width: 100%;
        overflow: hidden;
    }
    p {
        background-color: #f8f8b4;
        font-size: 1.8rem;
        text-align: justify;
        border: 0.1rem solid #f5f539;
        padding: 1.5rem;
        border-radius: 1rem;
    }

    .compact {
        width: 100%;
    }

    .compact h2 {
        margin: 0 0 0.4rem;
        color: #666;
        font-size: 1.4rem;
        font-weight: 400;
    }

    @media (max-width: 573px) {
        p {
            line-height: 3.7rem;
        }
    }
    @media (max-width: 450px) {
        h2 {
            font-size: 2.3rem;
        }
        .container-adds {
            margin: 0;
            padding: 0;
            width: 95%;
        }

        ins {
            width: 100%;
        }
    }
</style>
