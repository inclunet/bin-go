<script>
    import { createEventDispatcher } from "svelte";
    import Button from "./Button.svelte";

    const dispatch = createEventDispatcher();

    const handleClick = async () => {
        dispatch("callToAction");
    };
</script>

<Button on:click={handleClick}>Começar a Rodada</Button>
