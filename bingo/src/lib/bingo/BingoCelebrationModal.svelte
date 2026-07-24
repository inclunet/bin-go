<script>
    import {
        afterUpdate,
        createEventDispatcher,
        onDestroy,
        onMount,
        tick,
    } from "svelte";

    export let cardNumber = 0;
    export let completion = "";
    export let soundStatus = "idle";
    export let dismissError = "";
    export let dismissing = false;

    const dispatch = createEventDispatcher();
    /** @type {HTMLDialogElement | null} */
    let dialog = null;
    /** @type {HTMLButtonElement | null} */
    let dismissButton = null;
    /** @type {HTMLElement | null} */
    let previouslyFocused = null;
    let previousBodyOverflow = "";
    let completionLabel = "";

    const completionLabels = {
        Full: "Cartela cheia",
        Horizontal: "Linha horizontal",
        Vertical: "Linha vertical",
        Diagonal: "Diagonal",
    };

    $: completionLabel =
        completionLabels[completion] || "Você completou um bingo";

    const dismiss = () => dispatch("dismiss");
    const playSound = () => dispatch("playSound");

    /** @param {KeyboardEvent} event */
    const handleKeydown = (event) => {
        if (event.key === "Escape") {
            event.preventDefault();
            dismiss();
            return;
        }
        if (event.key !== "Tab" || !dialog) {
            return;
        }

        const focusable = /** @type {HTMLElement[]} */ (
            Array.from(dialog.querySelectorAll("button:not([disabled])"))
        );
        const first = focusable[0];
        const last = focusable[focusable.length - 1];
        if (!first || !last) {
            event.preventDefault();
            dialog.focus();
        } else if (
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

    onMount(() => {
        previouslyFocused = /** @type {HTMLElement | null} */ (
            document.activeElement
        );
        previousBodyOverflow = document.body.style.overflow;
        document.body.style.overflow = "hidden";
        if (dialog && !dialog.open) {
            dialog.showModal();
        }
        tick().then(() => dismissButton?.focus());
    });

    afterUpdate(() => {
        if (dialog && !dialog.contains(document.activeElement)) {
            if (dismissButton && !dismissButton.disabled) {
                dismissButton.focus();
            } else {
                dialog.focus();
            }
        }
    });

    onDestroy(() => {
        if (dialog?.open) {
            dialog.close();
        }
        document.body.style.overflow = previousBodyOverflow;
        previouslyFocused?.focus?.();
    });
</script>

<svelte:window on:keydown={handleKeydown} />

<dialog
    bind:this={dialog}
    class="celebration-modal"
    role="alertdialog"
    tabindex="-1"
    aria-modal="true"
    aria-labelledby="bingo-title"
    aria-describedby="bingo-message"
    on:cancel|preventDefault={dismiss}
>
        <header>
            <p class="eyebrow">Cartela #{cardNumber}</p>
            <h2 id="bingo-title">BINGO!</h2>
            <p id="bingo-message">
                Parabéns! <strong>{completionLabel}</strong>.
            </p>
        </header>

        {#if soundStatus === "blocked" || soundStatus === "stopped"}
            <div class="sound-status" aria-live="polite">
                <span aria-hidden="true">🔇</span>
                <p>
                    {soundStatus === "blocked"
                        ? "O navegador bloqueou o som."
                        : "O aviso sonoro foi interrompido."}
                </p>
            </div>
        {/if}

        {#if dismissError}
            <p class="dismiss-error" role="alert">{dismissError}</p>
        {/if}

        <div class="actions">
            {#if soundStatus === "blocked" || soundStatus === "stopped"}
                <button class="secondary-action" type="button" on:click={playSound}>
                    Ativar som
                </button>
            {/if}
            <button
                bind:this={dismissButton}
                class="primary-action"
                type="button"
                disabled={dismissing}
                aria-busy={dismissing}
                on:click={dismiss}
            >
                {dismissing ? "Fechando…" : "Parar aviso"}
            </button>
        </div>
</dialog>

<style>
    .celebration-modal {
        position: fixed;
        inset: 0;
        display: flex;
        width: 100vw;
        max-width: none;
        height: 100vh;
        height: 100dvh;
        max-height: none;
        align-items: center;
        justify-content: center;
        flex-direction: column;
        overflow-y: auto;
        margin: 0;
        padding: 3.2rem 2rem;
        border: 0;
        border-radius: 0;
        background:
            radial-gradient(
                circle at center,
                rgba(43, 126, 244, 0.1),
                transparent 42%
            ),
            var(--white, #fff);
        color: var(--senary-color, #1d1d1d);
        text-align: center;
    }

    .celebration-modal:not([open]) {
        display: none;
    }

    .celebration-modal::backdrop {
        background: var(--white, #fff);
    }

    header {
        width: min(48rem, 100%);
        padding: 0 0 2rem;
    }

    .eyebrow {
        margin: 0;
        color: var(--tertiary-color, #982a35);
        font-size: 1.6rem;
        font-weight: 700;
        letter-spacing: 0.08em;
        text-transform: uppercase;
    }

    h2 {
        margin: 0.4rem 0;
        color: var(--primary-color, #2b7ef4);
        font-size: clamp(4.2rem, 10vw, 6rem);
        line-height: 1;
    }

    header p:last-child {
        margin: 1.6rem 0 0;
        font-size: 2rem;
        line-height: 1.5;
    }

    .sound-status {
        display: flex;
        align-items: center;
        justify-content: center;
        gap: 1rem;
        width: min(42rem, 100%);
        margin: 0;
        padding: 1rem 1.4rem;
        border-left: 0.6rem solid var(--primary-color, #2b7ef4);
        border-radius: 0.4rem;
        background: #f5f9ff;
        text-align: left;
    }

    .sound-status span {
        font-size: 2.4rem;
    }

    .sound-status p {
        margin: 0;
        font-size: 1.6rem;
    }

    .actions {
        display: flex;
        justify-content: center;
        gap: 1.2rem;
        margin-top: 2rem;
        padding: 0;
    }

    .dismiss-error {
        width: min(42rem, 100%);
        margin: 1.6rem 0 0;
        padding: 1.2rem 1.6rem;
        border-left: 0.6rem solid var(--tertiary-color, #982a35);
        border-radius: 0.4rem;
        background: #fff4f5;
        font-size: 1.6rem;
        text-align: left;
    }

    button {
        min-width: 15rem;
        padding: 1rem 1.4rem;
        border: 0;
        border-radius: 0.4rem;
        color: var(--black, #000);
        font-size: 2rem;
        font-weight: 700;
        cursor: pointer;
    }

    .primary-action {
        background: var(--primary-button-color, #2b7ef4);
    }

    .secondary-action {
        background: var(--secondary-button-color, #e9b949);
    }

    button:hover,
    button:focus-visible {
        color: var(--white, #fff);
        filter: brightness(0.88);
    }

    button:disabled {
        cursor: wait;
        opacity: 0.72;
    }

    @media (max-width: 520px) {
        header {
            padding-bottom: 1.6rem;
        }

        .sound-status {
            margin: 0;
        }

        .dismiss-error {
            margin: 1.6rem 0 0;
        }

        .actions {
            flex-direction: column;
            width: min(32rem, 100%);
        }

        button {
            width: 100%;
            min-height: 4.8rem;
            font-size: 1.6rem;
        }
    }

    @media (prefers-reduced-motion: no-preference) {
        .celebration-modal {
            animation: celebrate-in 220ms ease-out;
        }
    }

    @keyframes celebrate-in {
        from {
            opacity: 0;
        }
    }
</style>
