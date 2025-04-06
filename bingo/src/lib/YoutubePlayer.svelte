<script>
    export let videoId = "";
    export let autoplay = false;
    export let controls = true;
    export let loop = false;
</script>

<div class="container">
    <div class="container-video">
        <iframe
            title="Youtube Video Player"
            src={`https://www.youtube.com/embed/${videoId}?autoplay=${autoplay ? 1 : 0}&controls=${controls ? 1 : 0}&loop=${loop ? 1 : 0}`}
            frameborder="0"
            allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture"
            allowfullscreen
        >
        </iframe>
    </div>
</div>

<style>
    .container {
        min-height: auto;
        width: 70%;
        margin-bottom: 40px;
        padding: 0;
    }
    .container-video {
        position: relative;
        display: flex;
        width: 100%;
        justify-content: center;
        padding-top: 56.25%;
    }
    iframe {
        position: absolute;
        top: 0;
        bottom: 0;
        left: 0;
        right: 0;
        width: 100%;
        height: 100%;
        border-radius: 8px;
        box-shadow: 0px 6px 15px 4px rgba(0, 0, 0, 0.7);
    }

    @media (max-width: 768px) {
        .container {
            width: 90%;
        }
    }
</style>
