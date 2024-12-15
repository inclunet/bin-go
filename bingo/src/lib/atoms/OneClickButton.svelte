<script>
    import { createEventDispatcher } from "svelte";

    export let x = 0;
    export let y = 0;

    /**
     * @type {HTMLButtonElement}
     */
    let btn;
    let dispatch = createEventDispatcher();

    export const focus = () => {
        btn.focus();
    };
    
    const handleClick = () => {
        if (btn.getAttribute("enabled") !== "false") {
            btn.setAttribute("disabled", "true");
            dispatch("click", { x, y });
        }
    };
</script>

<button bind:this={btn} on:click={handleClick}><slot /></button>
