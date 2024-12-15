<script>
    import { goto } from "$app/navigation";
    import Button from "./Button.svelte";
    import { callApi } from "./api";
    import { card } from "./bingo";

    const startGame = async () => {
        $card = await callApi(
            $card,
            `/api/bingo/${$card.Round}/${$card.Card}`,
            "GET"
        );

        goto(`/bingo/${$card.Round}/${$card.Card}`);
    };
</script>

<Button on:click={startGame}>Jogar!</Button>
