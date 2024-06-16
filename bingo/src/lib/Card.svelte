<script>
    import { card } from "./bingo";
    import Number from "./Number.svelte";
</script>

<div class="container-fluid d-flex text-center flex-column table-draw">
    <!-- <div class="table-draw"> -->

    <table summary="Cartela">
        <tr id="tr_first">
            <th scope="col">B</th>
            <th scope="col">I</th>
            <th scope="col">N</th>
            <th scope="col">G</th>
            <th scope="col">O</th>
        </tr>
        {#each $card.Numbers as row}
            <tr>
                {#each row as number}
                    <td>
                        <Number
                            {number}
                            mainCard={$card.Card == 1}
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
    @media (max-width: 991px) {
        .table-draw {
            align-items: end;
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
