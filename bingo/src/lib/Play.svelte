<script>
    import { goto } from "$app/navigation";
    import Button from "./Button.svelte";
    import { callApiResult } from "./api";
    import { card } from "./bingo";

    export let roundID = "";
    let playError = "";

    const startGame = async () => {
        if (!roundID) {
            playError = "Não foi possível identificar a rodada.";
            return;
        }
        const result = await callApiResult(
            $card,
            `/api/bingo/${roundID}/0`,
            "GET"
        );
        if (!result.ok || !result.data.RoundID || !result.data.ID) {
            playError =
                "Não foi possível abrir a cartela. Verifique sua conexão e tente novamente.";
            return;
        }
        playError = "";
        $card = result.data;
        goto(`/bingo/${$card.RoundID}/${$card.ID}`);
    };
</script>

<Button on:click={startGame}>Jogar!</Button>
{#if playError}
    <p class="alert alert-danger" role="alert">{playError}</p>
{/if}
