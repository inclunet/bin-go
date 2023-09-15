<script>
    
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
	$: outerWidth = 0
	$: innerWidth = 0
</script>

<svelte:window bind:innerWidth bind:outerWidth />

<div class="container-fluid d-flex align-items-center text-center flex-column">
    
    <table summary="Cartela">
        

        {#if card.Card == 1 && outerWidth >= 940}
            <tbody class="d-flex flex-row">
                {#each card.Numbers as row}
                    <tr>
                        {#each row as number}
                            <tr>
                                <Number bind:number bind:card />
                            </tr>
                        {/each}
                    </tr>
                {/each}
            </tbody>
        {:else}
            <tbody class="">
                {#each card.Numbers as row}
                    <tr>
                        {#each row as number}
                            <td>
                                <Number bind:number bind:card />
                            </td>
                        {/each}
                    </tr>
                {/each}
            </tbody>
        {/if}

        {#if card.Card > 1}
        <thead>
            <tr id="tr_first">
                <th scope="col">B</th>
                <th scope="col">I</th>
                <th scope="col">N</th>
                <th scope="col">G</th>
                <th scope="col">O</th>
            </tr>
        </thead>
            <tbody>
                {#each card.Numbers as row}
                    <tr>
                        {#each row as number}
                            <td>
                                <Number bind:number bind:card />
                            </td>
                        {/each}
                    </tr>
                {/each}
            </tbody>
        {/if}

    </table>
</div>
<audio src="/sms.mp3" />
