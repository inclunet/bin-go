<script>
    import { createEventDispatcher, tick } from "svelte";

    export let open = false;
    export let participantUrl = "";
    export let organizerUrl = "";

    const dispatch = createEventDispatcher();
    /** @type {HTMLButtonElement | null} */
    let firstButton = null;
    /** @type {HTMLElement | null} */
    let dialog = null;
    /** @type {HTMLElement | null} */
    let previouslyFocused = null;
    let wasOpen = false;
    let statusMessage = "";

    $: if (open && !wasOpen) {
        wasOpen = true;
        statusMessage = "";
        previouslyFocused =
            typeof document === "undefined"
                ? null
                : /** @type {HTMLElement | null} */ (document.activeElement);
        tick().then(() => firstButton?.focus());
    }

    $: if (!open && wasOpen) {
        wasOpen = false;
        const focusTarget = previouslyFocused;
        previouslyFocused = null;
        tick().then(() => focusTarget?.focus?.());
    }

    const close = () => {
        dispatch("close");
    };

    /** @param {string} url @param {string} successMessage */
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
            if (
                !(error instanceof DOMException) ||
                error.name !== "AbortError"
            ) {
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

    /** @param {KeyboardEvent} event */
    const handleKeydown = (event) => {
        if (!open) {
            return;
        }
        if (event.key === "Escape") {
            close();
            return;
        }
        if (event.key !== "Tab" || !dialog) {
            return;
        }

        const focusable = /** @type {HTMLElement[]} */ (
            Array.from(
                dialog.querySelectorAll(
                    'a[href], button:not([disabled]), input:not([disabled]), select:not([disabled]), textarea:not([disabled]), [tabindex]:not([tabindex="-1"])'
                )
            )
        );
        if (focusable.length === 0) {
            event.preventDefault();
            dialog.focus();
            return;
        }

        const first = focusable[0];
        const last = focusable[focusable.length - 1];
        if (
            event.shiftKey &&
            (document.activeElement === first ||
                !dialog.contains(document.activeElement))
        ) {
            event.preventDefault();
            last.focus();
        } else if (
            !event.shiftKey &&
            (document.activeElement === last ||
                !dialog.contains(document.activeElement))
        ) {
            event.preventDefault();
            first.focus();
        }
    };
</script>

<svelte:window on:keydown={handleKeydown} />

{#if open}
    <div class="modal-backdrop" role="presentation" on:click|self={close}>
        <section
            bind:this={dialog}
            class="share-modal"
            role="dialog"
            tabindex="-1"
            aria-modal="true"
            aria-labelledby="share-round-title"
        >
            <header>
                <div>
                    <p class="eyebrow">Opções da rodada</p>
                    <h2 id="share-round-title">Compartilhar rodada</h2>
                </div>
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
                    {#if typeof navigator !== "undefined" && Reflect.has(navigator, "share")}
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
        background: rgba(29, 29, 29, 0.72);
    }

    .share-modal {
        width: min(64rem, 100%);
        max-height: 90vh;
        overflow-y: auto;
        padding: 0;
        border: 0;
        border-radius: 1.2rem;
        background: var(--white, #fff);
        color: var(--senary-color, #1d1d1d);
        box-shadow: 0 0.8rem 2.4rem rgba(29, 29, 29, 0.16);
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
        padding: 2.8rem 3.2rem;
        background: linear-gradient(
            135deg,
            var(--primary-color, #2b7ef4),
            var(--octonary-color, #006aff)
        );
        color: var(--white, #fff);
    }

    header h2 {
        margin: 0.3rem 0 0;
        color: inherit;
        font-size: 3.2rem;
    }

    .eyebrow {
        margin: 0;
        color: inherit;
        font-size: 1.6rem;
        font-weight: 700;
        letter-spacing: 0.08em;
        text-transform: uppercase;
    }

    h3 {
        margin: 0 0 1.2rem;
        color: var(--senary-color, #1d1d1d);
    }

    .close-button {
        display: grid;
        width: 4.4rem;
        height: 4.4rem;
        place-items: center;
        padding: 0;
        border: 0;
        border-radius: 50%;
        background: transparent;
        color: var(--white, #fff);
        font-size: 3rem;
        line-height: 1;
    }

    .close-button:hover {
        background: rgba(255, 255, 255, 0.18);
    }

    .share-section {
        margin: 0;
        padding: 3.2rem;
        border: 0;
        background: var(--white, #fff);
    }

    .share-section p {
        margin-bottom: 1.4rem;
        line-height: 1.7;
    }

    .organizer {
        margin: 0 3.2rem 2.4rem;
        padding: 1.4rem 1.6rem;
        border-left: 0.6rem solid var(--tertiary-color, #982a35);
        border-radius: 0.4rem;
        background: #fff4f5;
    }

    .organizer h3 {
        color: var(--tertiary-color, #982a35);
    }

    button {
        padding: 1rem 1.4rem;
        border: 0;
        border-radius: 0.4rem;
        background: var(--primary-button-color, #2b7ef4);
        color: var(--black, #000);
        font-size: 2rem;
        font-weight: 700;
        cursor: pointer;
    }

    button:hover {
        background: var(--secondary-button-color, #2868c2);
        color: var(--white, #fff);
    }

    button:focus-visible {
        outline: 0.3rem solid var(--tertiary-button-color, #00f279);
        outline-offset: 0.3rem;
    }

    .actions {
        flex-wrap: wrap;
        gap: 1.2rem;
    }

    .status {
        min-height: 2.4rem;
        margin: 0 3.2rem 1.2rem;
        line-height: 1.5;
    }

    footer {
        justify-content: flex-end;
        padding: 2.4rem 3.2rem;
        border-top: 0.1rem solid #d8e5f8;
        background: #f5f9ff;
    }

    @media (max-width: 520px) {
        .modal-backdrop {
            align-items: flex-start;
            padding: 1rem;
        }

        header,
        footer {
            padding: 2rem;
        }

        .share-section {
            padding: 2rem;
        }

        header h2 {
            font-size: 2.6rem;
        }

        .organizer {
            margin: 0 2rem 2rem;
            padding: 1.4rem 1.6rem;
        }

        .actions,
        .actions button,
        .organizer button,
        footer button {
            width: 100%;
        }

        .status {
            margin: 0 2rem 1rem;
        }
    }
</style>
