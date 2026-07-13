<script>
    import { goto } from "$app/navigation";
    import { onMount } from "svelte";
    import Button from "$lib/Button.svelte";
    import PageTitle from "$lib/PageTitle.svelte";
    import StartRound from "$lib/StartRound.svelte";
    import ShareRoundModal from "$lib/bingo/ShareRoundModal.svelte";
    import { callApiResult } from "$lib/api";
    import { card } from "$lib/bingo";

    export let data;

    let loaded = false;
    let loadError = "";
    let shareOpen = false;
    let participantUrl = "";
    let organizerUrl = "";

    const loadRound = async () => {
        const roundID = String(data.Round);
        const cardID = String(data.Card);
        const result = await callApiResult(
            $card,
            `/api/bingo/${roundID}/${cardID}`,
            "GET"
        );

        if (
            !result.ok ||
            !result.data.ID ||
            !result.data.RoundID ||
            result.data.Card !== 1
        ) {
            loadError =
                "Não foi possível abrir o acesso do organizador desta rodada.";
            return;
        }

        $card = result.data;
        participantUrl = `${window.location.origin}/bingo/${$card.RoundID}/new`;
        organizerUrl = `${window.location.origin}/bingo/${$card.RoundID}/${$card.ID}/lobby`;
        loaded = true;
    };

    const startRound = () => {
        goto(`/bingo/${$card.RoundID}/${$card.ID}`);
    };

    onMount(loadRound);
</script>

<PageTitle title="Tela de espera da rodada" game="Inclubingo" />

<main class="container-fluid d-flex align-items-center flex-column">
    {#if loadError}
        <p class="alert alert-danger text-center" role="alert">{loadError}</p>
    {:else if loaded}
        <h2 class="text-center">Rodada #{$card.Round}</h2>
        <p class="instructions text-center">
            Peça aos participantes para apontarem a câmera para o QR Code. Cada
            pessoa poderá gerar sua própria cartela antes de começar.
        </p>

        <div class="qr-code">
            <img
                src="/qr/bingo/{$card.RoundID}"
                alt="QR Code para participantes entrarem na rodada {$card.Round}"
            />
        </div>

        <div class="actions">
            <Button on:click={() => (shareOpen = true)}>
                Compartilhar rodada
            </Button>
            <StartRound on:callToAction={startRound} />
        </div>
    {/if}
</main>

<ShareRoundModal
    open={shareOpen}
    {participantUrl}
    {organizerUrl}
    on:close={() => (shareOpen = false)}
/>

<style>
    main {
        padding-bottom: 3rem;
    }

    h2 {
        margin-top: 4rem;
        font-size: 2.8rem;
    }

    .instructions {
        max-width: 65rem;
        padding: 0 1rem;
        font-size: 1.8rem;
    }

    .qr-code {
        margin: 1rem 0 1.5rem;
        padding: 1rem;
        border-radius: 0.8rem;
        background: #fff;
    }

    .qr-code img {
        display: block;
        width: 24.8rem;
        max-width: 70vw;
        height: auto;
    }

    .actions {
        display: flex;
        flex-wrap: wrap;
        justify-content: center;
        gap: 1rem;
    }

    @media (max-width: 450px) {
        h2 {
            margin-top: 1rem;
            font-size: 2.3rem;
        }
    }
</style>
