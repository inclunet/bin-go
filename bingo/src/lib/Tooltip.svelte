<script>
    export let enableSpaceTip = false;
</script>

{#if enableSpaceTip}
    <div id="container-texto_botao-espaco">
        <div class="seta"></div>
        <div id="texto_botao-espaco" role="alert">
            <slot></slot>
            </div>
    </div>
{/if}

<style>
    #container-texto_botao-espaco {
        position: absolute;
        width: 15em;
        margin-top: 0.6em;
        margin-left: -3em;
    }
    #texto_botao-espaco {
        background-color: var(--tooltip);
        margin-left: -5em;
        padding: 1em;
        border-radius: 0.4em;
        font-weight: 500;
    }
    #texto_botao-espaco p {
        margin: 0;
        padding: 0;
        font-size: 1.2em;
    }
    .seta {
        width: 1.4em;
        height: 1.4em;
        margin-left: 75px;
        background: var(--tooltip);
        clip-path: polygon(50% 0, 100% 100%, 0 100%);
    }
</style>
