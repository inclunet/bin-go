<script>
    export let title;
</script>

<nav class="navbar">
    <button
        class="navbar-toggler button-menu navbar-dark"
        type="button"
        data-bs-toggle="offcanvas"
        data-bs-target="#offcanvasDarkNavbar"
        aria-controls="offcanvasDarkNavbar"
        aria-label="Menu de navegação"
    >
        <span class="navbar-toggler-icon"></span>
    </button>
    <div class="container-fluid">
        <div
            class="offcanvas offcanvas-end text-bg-dark"
            tabindex="-1"
            id="offcanvasDarkNavbar"
            aria-labelledby="offcanvasDarkNavbarLabel"
        >
            <div class="offcanvas-header">
                <h5 class="offcanvas-title" id="offcanvasDarkNavbarLabel">
                    {title}
                </h5>
                <button
                    type="button"
                    class="btn-close btn-close-white"
                    data-bs-dismiss="offcanvas"
                    aria-label="Fechar"
                ></button>
            </div>
            <div class="offcanvas-body">
                <ul class="navbar-nav justify-content-end flex-grow-1 pe-3">
                    <slot></slot>
                </ul>
            </div>
            <button
                type="button"
                class="btn-close btn-close-white hidden-area"
                data-bs-dismiss="offcanvas"
                aria-label="Fechar"
            ></button>
        </div>
    </div>
</nav>

<style>
    .btn-close {
        font-size: 3rem;
        padding: 2rem;
    }

    .offcanvas-header h5 {
        font-size: 2.5rem;
    }
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

    @media (max-width: 450px) {
        .navbar {
            display: flex;
            padding: 0;
        }
        .button-menu {
            display: flex;
            justify-content: center;
            align-items: center;
            width: 5rem;
            height: 5rem;
            margin: 0;
            padding: 0;
            background-color: var(--senary-color);
        }
        .navbar-toggler-icon {
            padding: 1.8rem;
        }
        .container-fluid {
            margin: 0;
            padding: 0;
        }
    }
</style>
