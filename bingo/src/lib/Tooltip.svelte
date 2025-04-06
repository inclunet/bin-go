<script>
    export let enableSpaceTip = false;
    export let marginTop = "0rem";
</script>

{#if enableSpaceTip}
    <div id="container-texto_botao" style="margin-top: {marginTop};">
        <div class="seta"></div>
        <div class="texto_tip-container" role="alert">
            <slot></slot>
        </div>
    </div>
{/if}

<style>
    :root {
        font-size: 62.5%;
    }
    #container-texto_botao {
        position: absolute;
        width: 30rem;
        margin-top: 0.6rem;
    }
    .texto_tip-container {
        background-color: var(--tooltip);
        padding: 1rem;
        border-radius: 0.4rem;
        font-weight: 500;
    }

    .seta {
        width: 1.8rem;
        height: 1.8rem;
        margin-left: 14rem;
        background: var(--tooltip);
        clip-path: polygon(50% 0, 100% 100%, 0 100%);
    }
</style>
