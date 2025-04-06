<script>
    import CardNumber from "$lib/bingo/CardNumber.svelte";

    export let card = [];
    let tableHeader = ["B", "I", "N", "G", "O"];
</script>

<table summary="Cartela">
    {#each tableHeader as lineHeader, c}
        {#if 1 == 1}
            <tr>
                <th scope="row">{lineHeader}</th>
                {#each card.Numbers as line}
                    <td>
                        <CardNumber
                            number={line[c]}
                            on:checkNumber
                            on:playCheckSound
                        />
                    </td>
                {/each}
            </tr>
        {/if}
    {/each}
</table>

<style>
    td {
        padding: 0 !important;
    }
</style>
