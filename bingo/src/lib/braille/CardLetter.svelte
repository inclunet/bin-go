<script>
    export let src;
    export let alt;
    export let nameLetter;
</script>

<div class="container_card">
    <div class="card">
        <img {src} class="card-img" {alt} />
        <div class="card-body">
            <h5 aria-hidden="true" class="card-title">{nameLetter}</h5>
        </div>
    </div>
</div>

<style>
    :root {
        font-size: 62.5%;
    }
    .container_card {
        width: 9rem;
        display: flex;
        justify-content: center;
    }
    .card {
        width: 6rem;
        height: 12rem;
        display: flex;
        flex-direction: column;
        align-items: center;
        border: none;
    }
    .card-img {
        background-color: #58e0de;
        border-radius: 0.8rem;
    }
    .card-body {
        display: flex;
        align-items: center;
        padding: 0;
    }
    .card-title {
        font-size: 2.5rem;
    }

    @media (max-width: 490px) {
        .card {
            width: 5rem;
        }
    }
</style>
