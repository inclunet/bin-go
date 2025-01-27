<script>
    export let videoId = '';
    export let width = '560';
    export let height = '315';
    export let autoplay = false;
    export let controls = true;
    export let loop = false;
</script>

<iframe
    width={width}
    height={height}
    src={`https://www.youtube.com/embed/${videoId}?autoplay=${autoplay ? 1 : 0}&controls=${controls ? 1 : 0}&loop=${loop ? 1 : 0}`}
    frameborder="0"
    allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture"
    allowfullscreen>
</iframe>

<style>
    iframe {
        border: none;
    }
</style>