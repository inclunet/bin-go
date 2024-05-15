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
                        <Number {number} on:checkNumber on:playCheckSound />
                    </td>
                {/each}
            </tr>
        {/each}
    </table>
</div>

<style>
    .table-draw {
        align-items: center;
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
    }
</style>
