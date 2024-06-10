<nav class="navbar menu-navbar">
    <ul
        class="container-nav-inner navbar-nav navbar-ul-primary d-flex flex-row"
    >
        <slot></slot>
    </ul>
</nav>

<style>
    :root {
        --primary-color: #2b7ef4;
        --secondary-color: #2b7ff4ab;
        --tertiary-color: #982a35;
        --quaternary-color: #fff;
        --quinary-color: #000;
        --senary-color: #1d1d1d;
        --setenary-color: #5799f5;
        --octonary-color: #006aff;
    }
    .container-nav-inner,
    .menu-navbar,
    ul {
        margin: 0;
    }

    .menu-navbar {
        background-color: var(--primary-color);
        border-top: 1px solid var(--quaternary-color);
        display: flex;
        justify-content: center;
    }
    .container-nav-inner {
        height: 48px;
        width: 80%;
        display: flex;
        justify-content: space-between;
        align-items: center;
        flex-wrap: wrap;
    }

    @media (max-width: 838px) {
        .container-nav-inner {
            padding: 0 25px 0 25px;
        }
    }
    @media (max-width: 745px) {
        .container-nav-inner {
            padding: 0 10px 0 10px;
        }
    }
    @media (max-width: 640px) {
        .container-nav-inner {
            padding: 0;
        }
        .navbar-ul-primary {
            width: 100%;
            justify-content: center;
        }
    }
    @media (max-width: 450px) {
        .menu-navbar {
            display: none;
        }
    }
</style>
