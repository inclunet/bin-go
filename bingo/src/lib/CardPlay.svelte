<script>
    export let title;
    export let titleButton;
    export let href;
    export let src;
    export let alt;
</script>

<div class="card">
    <img {src} class="card-img-top" {alt} />
    <div class="card-body">
        <h5 class="card-title">{title}</h5>
        <p class="card-text">
            <slot></slot>
        </p>
        <a {href} class="btn btn-primary">{titleButton}</a>
    </div>
</div>

<style>
    .card {
        width: 30rem;
        margin: 2rem;
    }
    .card-title {
        font-size: 2.5rem;
    }
    .card-body p {
        font-size: 1.7rem;
    }
    .card-body a {
        font-size: 1.6rem;
    }
</style>
