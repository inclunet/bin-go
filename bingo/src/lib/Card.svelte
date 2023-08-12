<script>
    import CardHeader from "$lib/CardHeader.svelte";
    import Number from "./Number.svelte";
    export let card = {
        Autoplay: false,
        Card: 0,
        Bingo: false,
        Checked: 0,
        LastNumber: 0,
        Round: 0,
        Numbers: [[{ Checked: false, Number: 0 }]],
    };
</script>

<div>
    <CardHeader bind:card />
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
                        <Number bind:number bind:card />
                    </td>
                {/each}
            </tr>
        {/each}
    </table>
</div>
<audio src="/sms.mp3" />
