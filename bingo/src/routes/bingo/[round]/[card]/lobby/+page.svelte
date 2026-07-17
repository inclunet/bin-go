<script>
    import { onMount } from "svelte";
    import Button from "$lib/Button.svelte";
    import PageTitle from "$lib/PageTitle.svelte";
    import StartRound from "$lib/StartRound.svelte";
    import ShareRoundModal from "$lib/bingo/ShareRoundModal.svelte";
    import { clearRoundCreationID } from "$lib/bingo/playerCards";
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
        clearRoundCreationID();
        participantUrl = `${window.location.origin}/bingo/${$card.RoundID}/new`;
        organizerUrl = `${window.location.origin}/bingo/${$card.RoundID}/${$card.ID}/lobby`;
        loaded = true;
    };

    const startRound = () => {
        window.location.assign(`/bingo/${$card.RoundID}/${$card.ID}`);
    };

    onMount(loadRound);
</script>

<PageTitle title="Tela de espera da rodada" game="Inclubingo" />

<main class="lobby-page container">
    {#if loadError}
        <p class="alert alert-danger text-center" role="alert">{loadError}</p>
    {:else if !loaded}
        <p class="loading text-center" role="status">Carregando a rodada…</p>
    {:else if loaded}
        <section class="lobby-card" aria-labelledby="lobby-title">
            <header class="lobby-heading">
                <p class="eyebrow">Painel do organizador</p>
                <h2 id="lobby-title">Rodada #{$card.Round}</h2>
                <p>
                    Convide os participantes e inicie a rodada quando todos
                    estiverem com suas cartelas.
                </p>
            </header>

            <div class="lobby-content">
                <div class="instructions">
                    <h3>Como convidar os participantes</h3>
                    <ol>
                        <li>
                            <span aria-hidden="true">1</span>
                            Mostre o QR Code ou use o botão de compartilhamento.
                        </li>
                        <li>
                            <span aria-hidden="true">2</span>
                            Cada pessoa abrirá o convite e gerará sua própria
                            cartela.
                        </li>
                        <li>
                            <span aria-hidden="true">3</span>
                            Quando todos estiverem prontos, comece a rodada.
                        </li>
                    </ol>

                    <p class="organizer-note">
                        <strong>Seu acesso é privado.</strong>
                        O link do organizador fica disponível somente nas opções
                        de compartilhamento.
                    </p>
                </div>

                <figure class="qr-panel">
                    <div class="qr-code">
                        <img
                            src="/qr/bingo/{$card.RoundID}"
                            alt="QR Code para participantes entrarem na rodada {$card.Round}"
                        />
                    </div>
                    <figcaption>
                        Aponte a câmera do celular para entrar na rodada
                    </figcaption>
                </figure>
            </div>

            <div class="actions" aria-label="Ações da rodada">
                <Button on:click={() => (shareOpen = true)}>
                    Compartilhar rodada
                </Button>
                <StartRound on:callToAction={startRound} />
            </div>
        </section>
    {/if}
</main>

<ShareRoundModal
    open={shareOpen}
    {participantUrl}
    {organizerUrl}
    on:close={() => (shareOpen = false)}
/>

<style>
    .lobby-page {
        width: min(104rem, calc(100% - 3rem));
        min-height: 55vh;
        padding: 4rem 0;
    }

    .loading {
        margin: 5rem 0;
    }

    .lobby-card {
        overflow: hidden;
        border: 0;
        border-radius: 1.2rem;
        background: var(--white, #fff);
        box-shadow: 0 0.8rem 2.4rem rgba(29, 29, 29, 0.16);
    }

    .lobby-heading {
        padding: 2.8rem 3.2rem;
        background: linear-gradient(
            135deg,
            var(--primary-color, #2b7ef4),
            var(--octonary-color, #006aff)
        );
        color: var(--white, #fff);
    }

    .lobby-heading h2 {
        margin: 0.3rem 0 0.8rem;
        color: inherit;
        font-size: 3.2rem;
    }

    .lobby-heading p {
        max-width: 76rem;
        margin: 0;
        color: inherit;
        line-height: 1.6;
    }

    .eyebrow {
        font-size: 1.6rem;
        font-weight: 700;
        letter-spacing: 0.08em;
        text-transform: uppercase;
    }

    .lobby-content {
        display: grid;
        grid-template-columns: minmax(0, 1fr) auto;
        align-items: center;
        gap: 4rem;
        padding: 3.2rem;
    }

    .instructions h3 {
        margin-bottom: 2rem;
        color: var(--senary-color, #1d1d1d);
    }

    .instructions ol {
        display: grid;
        gap: 1.6rem;
        margin: 0;
        padding: 0;
    }

    .instructions li {
        display: grid;
        grid-template-columns: 3.6rem 1fr;
        align-items: center;
        gap: 1.2rem;
        font-size: 1.8rem;
        line-height: 1.5;
    }

    .instructions li span {
        display: grid;
        width: 3.6rem;
        height: 3.6rem;
        place-items: center;
        border-radius: 50%;
        background: var(--primary-color, #2b7ef4);
        color: var(--white, #fff);
        font-weight: 700;
    }

    .organizer-note {
        margin: 2.4rem 0 0;
        padding: 1.4rem 1.6rem;
        border-left: 0.5rem solid var(--tertiary-color, #982a35);
        border-radius: 0.4rem;
        background: #fff4f5;
        line-height: 1.6;
    }

    .qr-panel {
        margin: 0;
        text-align: center;
    }

    .qr-code {
        padding: 1.4rem;
        border: 0.3rem solid var(--primary-color, #2b7ef4);
        border-radius: 1.2rem;
        background: var(--white, #fff);
        box-shadow: 0 0.4rem 1.2rem rgba(43, 126, 244, 0.2);
    }

    .qr-code img {
        display: block;
        width: 26rem;
        max-width: 65vw;
        height: auto;
    }

    figcaption {
        max-width: 28rem;
        margin-top: 1rem;
        font-size: 1.6rem;
        font-weight: 600;
        line-height: 1.5;
    }

    .actions {
        display: flex;
        flex-wrap: wrap;
        justify-content: center;
        gap: 1.2rem;
        padding: 2.4rem 3.2rem 3.2rem;
        border-top: 0.1rem solid #d8e5f8;
        background: #f5f9ff;
    }

    .actions :global(button) {
        min-width: 22rem;
    }

    @media (max-width: 760px) {
        .lobby-content {
            grid-template-columns: 1fr;
        }

        .qr-panel {
            justify-self: center;
        }
    }

    @media (max-width: 450px) {
        .lobby-page {
            width: calc(100% - 2rem);
            padding: 1.5rem 0 2.5rem;
        }

        .lobby-heading,
        .lobby-content,
        .actions {
            padding: 2rem;
        }

        .lobby-heading h2 {
            font-size: 2.6rem;
        }

        .lobby-heading p,
        .organizer-note {
            text-align: left;
        }

        .instructions li {
            font-size: 1.6rem;
        }

        .actions {
            flex-direction: column;
        }

        .actions :global(button) {
            width: 100%;
        }
    }
</style>
