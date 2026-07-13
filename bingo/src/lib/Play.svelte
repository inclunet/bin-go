<script>
    import { goto } from "$app/navigation";
    import Button from "./Button.svelte";
    import { callApiResult } from "./api";
    import { card } from "./bingo";

    let playError = "";

    const startGame = async () => {
        const result = await callApiResult(
            $card,
            `/api/bingo/${$card.RoundID}/${$card.ID}`,
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
