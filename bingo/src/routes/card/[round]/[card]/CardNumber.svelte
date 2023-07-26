<script>
    import { createEventDispatcher } from "svelte";

    const dispatch = createEventDispatcher();
    export let data = { Round: 0 };
    export let number = { Checked: false, Number: 0 };

    async function checkNumber() {
        number.Checked = true;
        dispatch("checkNumber", number);
    }
</script>

<button aria-pressed={number.Checked} on:click={checkNumber}>{number.Number}</button
>
