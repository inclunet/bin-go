<script>
    import Button from "./Button.svelte";
    import { callApiResult } from "./api";
    import { card } from "./bingo";
    import {
        forgetPlayerCard,
        getPlayerCardID,
        rememberPlayerCard,
    } from "./bingo/playerCards";

    export let roundID = "";
    let playError = "";
    let opening = false;

    /** @param {string} cardID */
    const openCard = (cardID) => {
        window.location.assign(`/bingo/${roundID}/${cardID}`);
    };

    const startGame = async () => {
        if (opening) {
            return;
        }
        if (!roundID) {
            playError = "Não foi possível identificar a rodada.";
            return;
        }

        opening = true;
        playError = "";

        const existingCardID = getPlayerCardID(roundID);
        if (existingCardID) {
            const existing = await callApiResult(
                $card,
                `/api/bingo/${roundID}/${existingCardID}`,
                "GET"
            );
            if (
                existing.ok &&
                existing.data.RoundID === roundID &&
                existing.data.ID === existingCardID &&
                existing.data.Card > 1
            ) {
                $card = existing.data;
                openCard(existingCardID);
                return;
            }
            forgetPlayerCard(roundID);
        }

        const created = await callApiResult(
            $card,
            `/api/bingo/${roundID}/0`,
            "GET"
        );
        if (
            !created.ok ||
            created.data.RoundID !== roundID ||
            !created.data.ID ||
            created.data.Card <= 1
        ) {
            playError =
                "Não foi possível abrir a cartela. Verifique sua conexão e tente novamente.";
            opening = false;
            return;
        }
        $card = created.data;
        rememberPlayerCard(roundID, created.data.ID);
        openCard(created.data.ID);
    };
</script>

<Button on:click={startGame} disabled={opening} ariaBusy={opening}>
    {opening ? "Abrindo cartela…" : "Jogar!"}
</Button>
{#if playError}
    <p class="alert alert-danger" role="alert">{playError}</p>
{/if}
