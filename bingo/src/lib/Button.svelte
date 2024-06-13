<script>
    export let color = "primary";
    export let data_bs_toggle = "";
    export let data_bs_target = "";
    export let data_bs_dismiss = "";
    export let aria_label = "";
</script>

<button
    class="btn btn-{color}"
    data-bs-toggle={data_bs_toggle}
    data-bs-target={data_bs_target}
    data-bs-dismiss={data_bs_dismiss}
    aria-label={aria_label}
    on:click
>
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
