<script>
    import { onMount } from "svelte";

    onMount(() => {
        (adsbygoogle = window.adsbygoogle || []).push({});
    })

</script>
<section class="container">
<h2>Anúncios</h2>
    <p>
        <em>
            Os Anúncios patrocinam o desenvolvimento e a manutenção 
            deste jogo de Bingo acessível.
        </em>
    </p>
    
    <!-- Transição do Bingo -->
    <ins class="adsbygoogle"
    style="display:block;"
    data-ad-client="ca-pub-6710397565046630"
    data-ad-slot="7204703685"
    data-ad-format="auto"
    data-full-width-responsive="true"></ins>
</section>

<style>
    section{
        width: auto;

    }
    h2{
        font-size: 1.6em;
    }
    p{
        background-color: #f8f8b4;
        font-size: 1.2em;
        text-align: justify;
        border: 1px solid #f5f539;
        padding: 15px;
        border-radius: 10px;
    }

    @media(max-width: 573px){
        p{
            line-height: 37px;
        }
    }
</style>