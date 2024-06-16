<script>
    import { onMount } from "svelte";

    onMount(() => {
        (adsbygoogle = window.adsbygoogle || []).push({});
    });
</script>

<section class="container">
    <h2>Anúncios</h2>
    <p>
        <em>
            Os Anúncios patrocinam o desenvolvimento e a manutenção deste jogo
            de Bingo acessível.
        </em>
    </p>

    <!-- Transição do Bingo -->
    <ins
        class="adsbygoogle"
        style="display:block;"
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
        font-size: 2.8rem;
    }
    p {
        background-color: #f8f8b4;
        font-size: 1.8rem;
        text-align: justify;
        border: 0.1rem solid #f5f539;
        padding: 1.5rem;
        border-radius: 1rem;
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
        .container {
            margin: 0;
            padding: 0;
            width: 95%;
        }
    }
</style>
