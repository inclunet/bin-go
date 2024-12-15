<nav class="navbar menu-navbar">
    <ul class="container-nav-inner navbar-nav navbar-ul-primary">
        <slot></slot>
    </ul>
</nav>

<style>
    :root {
        --primary-color: #2b7ef4;
        --quaternary-color: #fff;
        --quinary-color: #000;
        --senary-color: #1d1d1d;
        --setenary-color: #5799f5;
        --octonary-color: #006aff;

        font-size: 62.5%;
    }
    .container-nav-inner,
    .menu-navbar,
    ul {
        margin: 0;
    }
    .menu-navbar {
        width: 100%;
        background-color: var(--primary-color);
        border-top: 0.1rem solid var(--quaternary-color);
        display: flex;
        justify-content: center;
    }
    .container-nav-inner {
        width: 80%;
        display: flex;
        flex-direction: row;
        justify-content: space-between;
        flex-wrap: wrap;
    }

    @media (max-width: 838px) {
        .container-nav-inner {
            padding: 0 2.5rem 0 2.5rem;
        }
    }
    @media (max-width: 745px) {
        .container-nav-inner {
            padding: 0 1rem 0 1rem;
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
