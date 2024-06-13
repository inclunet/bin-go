<script>
    export let color = "primary";
</script>

<button class="btn btn-{color}" on:click>
    <strong><slot></slot></strong>
</button>

<style>
    button {
        padding: 10px;
        font-size: 1.3em;
    }

    @media (max-width: 450px) {
        button {
            height: 48px;
            padding: 7px;
            font-size: 1em;
        }
    }
</style>
