<script>
    import CardNumber from "$lib/bingo/CardNumber.svelte";

    export let card = { Numbers: [] };

    let table_client = false;

    $: table_client = card.Card > 1 ? true : false;
</script>

<div class="text-center flex-column table-draw" class:table_client>
    <table summary="Cartela">
        <tr id="tr_first">
            <th scope="col">B</th>
            <th scope="col">I</th>
            <th scope="col">N</th>
            <th scope="col">G</th>
            <th scope="col">O</th>
        </tr>
        {#each card.Numbers as row}
            <tr>
                {#each row as number}
                    <td>
                        <CardNumber
                            {number}
                            mainCard={card.Card == 1}
                            on:checkNumber
                            on:playCheckSound
                        />
                    </td>
                {/each}
            </tr>
        {/each}
    </table>
</div>

<style>
    :root {
        font-size: 62.5%;
    }
    .table-draw {
        align-items: center;
    }
    th {
        font-size: 2rem;
    }

    .table_client {
        /* width: 50vw !important; */
        /* width: 30vw; */
        display: flex;
        align-items: center !important;
    }
    @media (max-width: 991px) {
        .table-draw {
            align-items: start;
        }
    }

    @media (max-width: 910px) {
        .table-draw {
            align-items: start;
        }
    }
    @media (max-width: 767px) {
        .table-draw {
            align-items: center;
        }
        th {
            font-size: 1.3rem;
        }
    }
</style>
