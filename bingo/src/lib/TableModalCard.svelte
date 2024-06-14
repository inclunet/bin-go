<script>
    import Button from "./Button.svelte";
    export let title;
</script>

<div
    class="modal fade"
    id="exampleModal"
    tabindex="-1"
    aria-labelledby="exampleModalLabel"
    aria-hidden="true"
>
    <div class="modal-dialog">
        <div class="modal-content">
            <div class="modal-header">
                <h1 class="modal-title fs-5" id="exampleModalLabel">
                    {title}
                </h1>
                <Button color="close" data_bs_dismiss="modal" aria_label="Close"
                ></Button>
            </div>
            <div class="modal-body"><slot></slot></div>
            <div class="modal-footer">
                <Button color="secondary" data_bs_dismiss="modal">Fechar</Button
                >
                <Button color="success">Salvar</Button>
            </div>
        </div>
    </div>
</div>
