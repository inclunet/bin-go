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
        margin: 1em;
        width: 70px;
        height: 70px;
        color: black;
        font-weight: 900;
        font-size: 1.8em;
    }
    @media(max-width: 680px){
        button{
            margin: 0.6em;
            width: 55px;
            height: 55px;
            font-weight: 900;
            font-size: 1.6em;
        }
    }
    @media(max-width: 470px){
        button{
            margin: 0.5em;
            width: 50px;
            height: 50px;
            font-weight: 900;
            font-size: 1.5em;
        }
    }
    @media(max-width: 400px){
        button{
            margin: 0.3em;
            width: 50px;
            height: 50px;
            font-weight: 900;
            font-size: 1.3em;
        }
    }
    @media(max-width: 350px){
        button{
            margin: 0.2em;
            width: 45px;
            height: 45px;
            font-weight: 900;
            font-size: 1.1em;
        }
    }
</style>