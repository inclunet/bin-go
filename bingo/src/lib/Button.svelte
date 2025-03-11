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
    :root {
        font-size: 62.5%;
    }
    button {
        margin-top: 0;
        padding: 1rem;
        font-size: 2rem;
    }

    @media (max-width: 450px) {
        button {
            height: 4.8rem;
            padding: 0.5rem;
            font-size: 1.5rem;
        }
    }

    @media (max-width: 410px) {
        button {
            height: 4.8rem;
            padding: 0.5rem 0.3rem;
            font-size: 1.4rem;
        }
    }
</style>
