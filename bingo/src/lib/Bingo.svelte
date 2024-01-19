<script>
    import { card } from "./bingo";
    import { afterUpdate, onMount } from "svelte";

    let audio = undefined;
    let drawed = 0;
    let muted = true;

    const isBingo = () => {
        if ($card.Card > 1 && $card.Bingo) {
            if ($card.Checked != drawed && muted) {
                drawed = $card.Checked;
                muted = false;
                audio.play();
            }
        }
    };

    function stopBingoAlert() {
        muted = true;
        audio.pause();
        audio.currentTime = 0;
    }

    afterUpdate(isBingo);
</script>

{#if $card.Card > 1 && $card.Bingo}
    <button class="btn btn-warning btn-lg" on:click={stopBingoAlert}
        >Bingo!</button
    >
{/if}

<audio bind:this={audio} src="/sms.mp3" loop></audio>
