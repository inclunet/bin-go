<script>
    import { createEventDispatcher, tick } from "svelte";

    export let open = false;
    export let participantUrl = "";
    export let organizerUrl = "";

    const dispatch = createEventDispatcher();
    let firstButton;
    let statusMessage = "";

    $: if (open) {
        statusMessage = "";
        tick().then(() => firstButton?.focus());
    }

    const close = () => {
        dispatch("close");
    };

    const copyLink = async (url, successMessage) => {
        try {
            await navigator.clipboard.writeText(url);
            statusMessage = successMessage;
        } catch (error) {
            statusMessage = "Não foi possível copiar o link.";
        }
    };

    const shareParticipantLink = async () => {
        try {
            await navigator.share({
                title: "Rodada de Inclubingo",
                text: "Use este link para gerar sua cartela de bingo.",
                url: participantUrl,
            });
            statusMessage = "Convite compartilhado.";
        } catch (error) {
            if (error?.name !== "AbortError") {
                statusMessage = "Não foi possível compartilhar o convite.";
            }
        }
    };

    const shareOnWhatsApp = () => {
        const message = encodeURIComponent(
            `Participe desta rodada de Inclubingo: ${participantUrl}`
        );
        window.open(
            `https://wa.me/?text=${message}`,
            "_blank",
            "noopener,noreferrer"
        );
    };

    const handleKeydown = (event) => {
        if (open && event.key === "Escape") {
            close();
        }
    };
</script>

<svelte:window on:keydown={handleKeydown} />

{#if open}
    <div class="modal-backdrop" role="presentation" on:click|self={close}>
        <section
            class="share-modal"
            role="dialog"
            aria-modal="true"
            aria-labelledby="share-round-title"
        >
            <header>
                <h2 id="share-round-title">Compartilhar rodada</h2>
                <button
                    class="close-button"
                    type="button"
                    aria-label="Fechar"
                    on:click={close}>×</button
                >
            </header>

            <div class="share-section">
                <h3>Convite para participantes</h3>
                <p>
                    Este acesso permite que cada participante gere sua própria
                    cartela.
                </p>
                <div class="actions">
                    <button
                        bind:this={firstButton}
                        type="button"
                        on:click={() =>
                            copyLink(participantUrl, "Convite copiado.")}>
                        Copiar convite
                    </button>
                    <button type="button" on:click={shareOnWhatsApp}>
                        Compartilhar pelo WhatsApp
                    </button>
                    {#if typeof navigator !== "undefined" && navigator.share}
                        <button type="button" on:click={shareParticipantLink}>
                            Mais opções
                        </button>
                    {/if}
                </div>
            </div>

            <div class="share-section organizer">
                <h3>Acesso do organizador</h3>
                <p>
                    Este acesso controla o sorteio. Guarde-o em um local seguro
                    e não o compartilhe com os participantes.
                </p>
                <button
                    type="button"
                    on:click={() =>
                        copyLink(
                            organizerUrl,
                            "Link do organizador copiado."
                        )}>
                    Copiar meu link de organizador
                </button>
            </div>

            <p class="status" aria-live="polite">{statusMessage}</p>
            <footer>
                <button type="button" on:click={close}>Fechar</button>
            </footer>
        </section>
    </div>
{/if}

<style>
    .modal-backdrop {
        position: fixed;
        inset: 0;
        z-index: 1050;
        display: flex;
        align-items: center;
        justify-content: center;
        padding: 1.5rem;
        background: rgba(0, 0, 0, 0.72);
    }

    .share-modal {
        width: min(58rem, 100%);
        max-height: 90vh;
        overflow-y: auto;
        padding: 2rem;
        border-radius: 0.8rem;
        background: var(--background-color, #fff);
        color: var(--text-color, #111);
    }

    header,
    footer,
    .actions {
        display: flex;
        align-items: center;
        gap: 1rem;
    }

    header {
        justify-content: space-between;
    }

    h2,
    h3 {
        margin-top: 0;
    }

    .close-button {
        border: 0;
        background: transparent;
        color: inherit;
        font-size: 3rem;
        line-height: 1;
    }

    .share-section {
        margin-top: 1.5rem;
        padding-top: 1.5rem;
        border-top: 1px solid currentColor;
    }

    .organizer {
        padding: 1.5rem;
        border: 2px solid #b45309;
        border-radius: 0.6rem;
    }

    button {
        padding: 0.8rem 1.2rem;
        border: 0;
        border-radius: 0.4rem;
        background: var(--primary-button-color, #f7c948);
        color: #111;
        font-weight: 700;
        cursor: pointer;
    }

    .actions {
        flex-wrap: wrap;
    }

    .status {
        min-height: 2.4rem;
        margin: 1rem 0;
    }

    footer {
        justify-content: flex-end;
    }
</style>
