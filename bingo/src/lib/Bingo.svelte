<script>
    export let card;

function stopAudio () {
    var audio = document.getElementsByTagName("audio")[0];
    audio.pause();
    audio.currentTime = 0;
}

</script>

{#if card.Card > 1 && card.Bingo }
    <audio src="/sms.mp3" autoplay loop ></audio>
    <button class="btn btn-warning btn-lg" on:click={stopAudio}>Bingo!</button>
{/if}