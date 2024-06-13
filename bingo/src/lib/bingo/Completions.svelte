<script>
    import { createEventDispatcher } from "svelte";
    import Completion from "./Completion.svelte";

    export let card = 0;
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
    };
</script>

<div>
    <h1>Configurações</h1>
    {card}
    <table summary="Configurações da rodada">
        <tr>
            <th scope="col">Bingos Válidos</th>
            <th scope="col">Premiados</th>
            <th scope="col">Permitidos</th>
        </tr>
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
        <Completion {card} bind:completion={completions.Diagonal}>Diagonal</Completion>
    </table>
    {#if card == 1}
        <button on:click={handleClick}>Salvar</button>
    {/if}
</div>
