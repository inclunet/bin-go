<script>
    import { HelpIcon } from "./icons";

    export let target;
    export let aria_label;
    export let title;
    export let icon;
    export let color;

    switch (color) {
        case "green":
            color = "--tertiary-button-color";
            break;

        default:
            color = "--quaternary-button-color";
            break;
    }
</script>

<nav class="navbar">
    <div class="container-fluid">
        <div class="fixed-div">
            <button
                class="button-rules"
                type="button"
                style="background-color: var({color});"
                data-bs-toggle="offcanvas"
                data-bs-target={`#${target}`}
                aria-controls="offcanvasNavbar"
                aria-label={aria_label}
            >
                {@html icon}
            </button>
        </div>
        <div
            class="offcanvas offcanvas-start"
            tabindex="-1"
            id={target}
            aria-labelledby="offcanvasNavbarLabel"
        >
            <div class="offcanvas-header">
                <h2>{title}</h2>
                <button
                    type="button"
                    class="btn-close button-rules_close"
                    data-bs-dismiss="offcanvas"
                    aria-label="Fechar"
                ></button>
            </div>
            <div class="offcanvas-body">
                <div class="container_card-letter">
                    <slot></slot>
                </div>
            </div>
            <button
                type="button"
                class="hidden-area"
                data-bs-dismiss="offcanvas"
                aria-label="Fechar"
            ></button>
        </div>
    </div>
</nav>

<style>
    .hidden-area {
        position: absolute;
        width: 1px;
        height: 1px;
        padding: 0;
        bottom: 0;
        margin: -1px;
        overflow: hidden;
        clip: rect(0, 0, 0, 0);
        border: 0;
    }
    .container-fluid {
        padding-left: 0;
    }

    .button-rules {
        display: flex;
        justify-content: center;
        align-items: center;
        width: 5rem;
        height: 5rem;
        left: 0.3rem;
        margin: 0;
        padding: 0 !important;
        border-radius: 0.8rem;
        cursor: pointer;
    }

    .button-rules:active {
        border: 1px solid var(--secondary-button-rules-color);
    }

    .offcanvas-header {
        display: flex;
        justify-content: space-between;
    }

    h2 {
        margin: 0;
        padding: 0;
        font-size: 2.8rem;
    }
    .offcanvas-body {
        display: flex;
        justify-content: center;
    }

    .container_card-letter {
        width: 40rem;
        height: 50rem;
        display: flex;
        justify-content: center;
        flex-wrap: wrap;
        margin-top: 3rem;
    }
    .button-rules_close {
        font-size: 3rem;
        padding: 2rem;
    }

    @media (hover: hover) {
        .button-rules:hover {
            background-color: var(--quinary-button-color) !important;
        }
    }
</style>
