<script>
    export let videoId = "";
    export let width = "720";
    export let height = "405";
    export let autoplay = false;
    export let controls = true;
    export let loop = false;
</script>

<div class="container-video">
    <iframe
        title="Youtube Video Player"
        {width}
        {height}
        src={`https://www.youtube.com/embed/${videoId}?autoplay=${autoplay ? 1 : 0}&controls=${controls ? 1 : 0}&loop=${loop ? 1 : 0}`}
        frameborder="0"
        allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture"
        allowfullscreen
    >
    </iframe>
</div>

<style>
    iframe {
        border-radius: 8px;
        box-shadow: 0px 6px 15px 4px rgba(0, 0, 0, 0.7);
    }

    .container-video {
        display: flex;
        width: 100%;
        justify-content: center;
    }
</style>
