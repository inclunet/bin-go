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
        width: 15em;
        margin-top: 0.6em;
    }
    .texto_tip-container {
        background-color: var(--tooltip);
        margin-left: -5em;
        padding: 1em;
        border-radius: 0.4em;
        font-weight: 500;
    }

    .seta {
        width: 1.4em;
        height: 1.4em;
        margin-left: 75px;
        background: var(--tooltip);
        clip-path: polygon(50% 0, 100% 100%, 0 100%);
    }
</style>
