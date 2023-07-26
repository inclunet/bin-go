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

<button aria-pressed={number.Checked} on:click={checkNumber} class="btn active">{number.Number}</button>
<style>
    button[aria-pressed="true"]{
        background-color: #008000;
        color: white;
        border-color: black;
    }
    button{
        margin: 0.5em;
        width: 60px;
        height: 60px;
        color: black;
        font-weight: 900;
        font-size: 1.6em;
    }
</style>