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
    /** @type {HTMLElement | null} */
    let dialog = null;
    /** @type {HTMLButtonElement | null} */
    let dismissButton = null;
    /** @type {HTMLElement | null} */
    let previouslyFocused = null;
    let previousBodyOverflow = "";

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
        document.body.style.overflow = previousBodyOverflow;
        previouslyFocused?.focus?.();
    });
</script>

<svelte:window on:keydown={handleKeydown} />

<div class="celebration-backdrop" role="presentation">
    <section
        bind:this={dialog}
        class="celebration-modal"
        role="alertdialog"
        tabindex="-1"
        aria-modal="true"
        aria-labelledby="bingo-title"
        aria-describedby="bingo-message"
    >
        <div class="confetti" aria-hidden="true">
            <span>B</span><span>I</span><span>N</span><span>G</span><span>O</span>
        </div>

        <header>
            <p class="eyebrow">Cartela #{cardNumber}</p>
            <h2 id="bingo-title">BINGO!</h2>
            <p id="bingo-message">
                Parabéns! <strong>{completionLabel}</strong>.
            </p>
        </header>

        <div class="sound-status" aria-live="polite">
            {#if soundStatus === "blocked"}
                <span aria-hidden="true">🔇</span>
                <p>O navegador bloqueou o som. Ative-o pelo botão abaixo.</p>
            {:else if soundStatus === "playing"}
                <span aria-hidden="true">🔊</span>
                <p>O aviso sonoro está tocando.</p>
            {:else if soundStatus === "stopped"}
                <span aria-hidden="true">🔇</span>
                <p>O aviso sonoro foi interrompido.</p>
            {:else}
                <span aria-hidden="true">🔊</span>
                <p>Preparando o aviso sonoro…</p>
            {/if}
        </div>

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
    </section>
</div>

<style>
    .celebration-backdrop {
        position: fixed;
        inset: 0;
        z-index: 1100;
        display: flex;
        align-items: center;
        justify-content: center;
        padding: 1.5rem;
        background: rgba(29, 29, 29, 0.78);
    }

    .celebration-modal {
        width: min(56rem, 100%);
        overflow: hidden;
        border: 0;
        border-radius: 1.2rem;
        background: var(--white, #fff);
        color: var(--senary-color, #1d1d1d);
        text-align: center;
        box-shadow: 0 1.2rem 3.6rem rgba(0, 0, 0, 0.3);
    }

    .confetti {
        display: flex;
        justify-content: center;
        gap: 1rem;
        padding: 2.4rem 2rem 0;
    }

    .confetti span {
        display: grid;
        width: 5.2rem;
        height: 5.2rem;
        place-items: center;
        border-radius: 50%;
        background: var(--primary-color, #2b7ef4);
        color: var(--white, #fff);
        font-size: 2.6rem;
        font-weight: 800;
        box-shadow: 0 0.4rem 0 #174f9d;
    }

    .confetti span:nth-child(even) {
        background: var(--tertiary-color, #982a35);
        box-shadow: 0 0.4rem 0 #681923;
    }

    header {
        padding: 2.4rem 3.2rem 2rem;
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
        font-size: clamp(4.8rem, 12vw, 7.2rem);
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
        margin: 0 3.2rem;
        padding: 1.4rem 1.6rem;
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
        margin-top: 2.4rem;
        padding: 2.4rem 3.2rem;
        border-top: 0.1rem solid #d8e5f8;
        background: #f5f9ff;
    }

    .dismiss-error {
        margin: 1.6rem 3.2rem 0;
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
        .confetti {
            gap: 0.6rem;
        }

        .confetti span {
            width: 4rem;
            height: 4rem;
            font-size: 2rem;
        }

        header {
            padding: 2rem;
        }

        .sound-status {
            margin: 0 2rem;
        }

        .dismiss-error {
            margin: 1.6rem 2rem 0;
        }

        .actions {
            flex-direction: column;
            padding: 2rem;
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

        .confetti span {
            animation: ball-pop 420ms ease-out both;
        }

        .confetti span:nth-child(2) {
            animation-delay: 50ms;
        }

        .confetti span:nth-child(3) {
            animation-delay: 100ms;
        }

        .confetti span:nth-child(4) {
            animation-delay: 150ms;
        }

        .confetti span:nth-child(5) {
            animation-delay: 200ms;
        }
    }

    @keyframes celebrate-in {
        from {
            opacity: 0;
            transform: scale(0.92);
        }
    }

    @keyframes ball-pop {
        from {
            opacity: 0;
            transform: translateY(-1.5rem) scale(0.7);
        }
    }
</style>
