<script>
    export let card = 0;
    export let completion = { Quantity: 0, Max: 0 };
</script>

<tr>
    <th scope="row">
        <slot></slot>
    </th>
    <td>{completion.Quantity}</td>
    <td>
        {#if card == 1}
            <input
                type="number"
                min={completion.Quantity}
                max="100"
                bind:value={completion.Max}
            />
        {:else}
            {completion.Max}
        {/if}
    </td>
</tr>
