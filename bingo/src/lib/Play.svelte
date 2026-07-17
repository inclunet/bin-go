<script>
    import Button from "./Button.svelte";
    import { callApiResult } from "./api";
    import { card } from "./bingo";
    import {
        forgetPlayerCard,
        getPlayerCardID,
        getPlayerID,
        rememberPlayerCard,
    } from "./bingo/playerCards";

    export let roundID = "";
    let playError = "";
    let opening = false;

    /** @param {string} cardID */
    const openCard = (cardID) => {
        window.location.assign(`/bingo/${roundID}/${cardID}`);
        window.setTimeout(() => {
            opening = false;
        }, 2000);
    };

    const openOrCreateCard = async () => {
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
                return true;
            }
            forgetPlayerCard(roundID);
        }

        const created = await callApiResult(
            $card,
            `/api/bingo/${roundID}/0`,
            "GET",
            null,
            false,
            { "X-Bingo-Player-ID": getPlayerID() }
        );
        if (
            !created.ok ||
            created.data.RoundID !== roundID ||
            !created.data.ID ||
            created.data.Card <= 1
        ) {
            playError =
                "Não foi possível abrir a cartela. Verifique sua conexão e tente novamente.";
            return false;
        }
        $card = created.data;
        rememberPlayerCard(roundID, created.data.ID);
        openCard(created.data.ID);
        return true;
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

        try {
            const navigated = navigator.locks
                ? await navigator.locks.request(
                      `inclubingo-card-${roundID}`,
                      openOrCreateCard
                  )
                : await openOrCreateCard();
            if (!navigated) {
                opening = false;
            }
        } catch {
            playError =
                "Não foi possível abrir a cartela. Verifique sua conexão e tente novamente.";
            opening = false;
        }
    };
</script>

<Button on:click={startGame} disabled={opening} ariaBusy={opening}>
    {opening ? "Abrindo cartela…" : "Jogar!"}
</Button>
{#if playError}
    <p class="alert alert-danger" role="alert">{playError}</p>
{/if}
