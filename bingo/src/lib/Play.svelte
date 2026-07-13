<script>
    import { goto } from "$app/navigation";
    import Button from "./Button.svelte";
    import { callApi } from "./api";
    import { card } from "./bingo";

    const startGame = async () => {
        $card = await callApi(
            $card,
            `/api/bingo/${$card.RoundID}/${$card.ID}`,
            "GET"
        );

        goto(`/bingo/${$card.RoundID}/${$card.ID}`);
    };
</script>

<Button on:click={startGame}>Jogar!</Button>
