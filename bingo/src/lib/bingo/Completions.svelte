<script>
    import { createEventDispatcher, onMount } from "svelte";
    import Completion from "./Completion.svelte";

    export let card = 0;
    export let config = true;
    export let completions = {
        Total: {
            Max: 0,
        },
        Full: {
            Max: 0,
        },
        Horizontal: {
            Max: 0,
        },
        Vertical: {
            Max: 0,
        },
        Diagonal: {
            Max: 0,
        },
    };

    const dispatch = createEventDispatcher();

    const handleClick = () => {
        dispatch("saveCompletions");
        config = false;
    };
</script>

{#if config}
    <div role="dialog">
        <!-- <h1 tabindex="0">Configurações</h1> -->
        <table class="table table-striped" summary="Configurações da rodada">
            <thead>
                <tr>
                    <th scope="col">Bingos Válidos</th>
                    <th scope="col">Premiados</th>
                    <th scope="col">Permitidos</th>
                </tr>
            </thead>
            <tbody>
                <Completion {card} bind:completion={completions.Total}
                    >Total de ganhadores</Completion
                >
                <Completion {card} bind:completion={completions.Intermediary}
                    >Total de quinas</Completion
                >
                <Completion {card} bind:completion={completions.Full}
                    >Cartela cheia</Completion
                >
                <Completion {card} bind:completion={completions.Horizontal}
                    >Horizontal</Completion
                >
                <Completion {card} bind:completion={completions.Vertical}
                    >Vertical</Completion
                >
                <Completion {card} bind:completion={completions.Diagonal}
                    >Diagonal</Completion
                >
            </tbody>
        </table>
        {#if card == 1}
            <button on:click={handleClick}>Salvar</button>
        {/if}
    </div>
{/if}
