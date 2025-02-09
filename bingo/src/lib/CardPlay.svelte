<script>
    export let title;
    export let href;
    export let src;
    export let alt;
</script>

<section>
    <a {href}>
        <div class="card">
            <img {src} class="card-img-top" {alt} />
            <div class="card-body">
                <h5 class="card-title">{title}</h5>
                <p class="card-text">
                    <slot></slot>
                </p>
            </div>
        </div>
    </a>
</section>

<style>
    .card {
        width: 35rem;
        margin-bottom: 5rem;
    }
    .card-title {
        font-size: 2.5rem;
    }
    .card-body p {
        font-size: 1.8rem;
    }
    a {
        text-decoration: none;
    }
</style>
