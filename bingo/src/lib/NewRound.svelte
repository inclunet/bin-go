<script>
    import { goto } from "$app/navigation";
    import { callApi } from "./api";
    import { card } from "./bingo";

    const newRound = async () => {
        $card = await callApi(
            $card,
            `/api/bingo/${$card.Round}/new/${$card.Type}`,
            "GET",
        );

        goto(`/bingo/${$card.Round}`);
    };
</script>

{#if ($card.Card == 1 && $card.Checked == $card.Type) || $card.Round == 0}
    <button class="btn" on:click={newRound}><strong>Nova Rodada</strong></button
    >
{/if}

<style>
    :root {
        --primary-button-color: #2b7ef4;
        --secondary-button-color: #2868c2;
        --tertiary-button-color: #2157a1;
    }
    button {
        padding: 10px;
        font-size: 1.3em;
        background-color: var(--primary-button-color);
    }
    button:hover {
        color: #fff;
        background-color: var(--secondary-button-color);
    }
    .btn:active {
        color: #fff;
        background-color: var(--tertiary-button-color);
    }
    @media (max-width: 450px) {
        button {
            height: 48px;
            padding: 7px;
            font-size: 1em;
        }
    }
</style>
