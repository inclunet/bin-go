<script>
    import { createEventDispatcher } from "svelte";

    const dispatch = createEventDispatcher();

    const handleClick = () => {
        dispatch("draw");
    };
</script>

<button class="btn btn-primary" on:click={handleClick}
    ><strong>Sortear</strong></button
>

<style>
    button {
        padding: 10px;
        font-size: 1.3em;
        color: #000;
    }
    @media (max-width: 450px) {
        button {
            height: 48px;
            padding: 7px;
            font-size: 1em;
        }
    }
</style>
