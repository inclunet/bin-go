<script>
    import { goto } from "$app/navigation";
    import Button from "./Button.svelte";
    import { callApiResult } from "./api";
    import { card } from "./bingo";

    const startGame = async () => {
        const result = await callApiResult(
            $card,
            `/api/bingo/${$card.RoundID}/${$card.ID}`,
            "GET"
        );
        if (!result.ok || !result.data.RoundID || !result.data.ID) {
            return;
        }
        $card = result.data;
        goto(`/bingo/${$card.RoundID}/${$card.ID}`);
    };
</script>

<Button on:click={startGame}>Jogar!</Button>
