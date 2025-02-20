<script>
    import CardLetter from "./CardLetter.svelte";
    import { braille } from "$lib/braille";
    import { lettersList } from "./lettersList";
    import { onMount, onDestroy } from "svelte";
    import { HelpIcon } from "./icons";

    let showDiv = false;
    let triggerPosition = 0;
    let triggerPositionOff = 400;
    let divStyle = "display: none;";

    function updateTriggerPosition() {
        if (window.matchMedia("(max-width: 494px)").matches) {
            triggerPosition = 200;
            triggerPositionOff = 500;
        } else {
            triggerPosition = 350;
            triggerPositionOff = 550;
        }
    }

    const handleScroll = () => {
        updateTriggerPosition();

        showDiv =
            window.scrollY >= triggerPosition &&
            window.scrollY <= triggerPositionOff;

        divStyle = showDiv ? "display: block;" : "display: none;";
    };

    if (typeof window !== "undefined") {
        onMount(() => {
            window.addEventListener("scroll", handleScroll);
        });

        onDestroy(() => {
            window.removeEventListener("scroll", handleScroll);
        });
    }
</script>

<nav class="navbar">
    <div class="container-fluid">
        <div class="fixed-div" style={divStyle}>
            <button
                class="button-rules"
                type="button"
                data-bs-toggle="offcanvas"
                data-bs-target="#offcanvasNavbar"
                aria-controls="offcanvasNavbar"
                aria-label="Regras da composição das letras em braille"
            >
                {@html HelpIcon(50)}
            </button>
        </div>
        <div
            class="offcanvas offcanvas-start"
            tabindex="-1"
            id="offcanvasNavbar"
            aria-labelledby="offcanvasNavbarLabel"
        >
            <div class="offcanvas-header">
                <h2>Letras da Série {$braille.CurrentClass + 1}</h2>
                <button
                    type="button"
                    class="btn-close button-rules_close"
                    data-bs-dismiss="offcanvas"
                    aria-label="Fechar"
                ></button>
            </div>
            <div class="offcanvas-body">
                <div class="container_card-letter">
                    {#each lettersList as letters, index}
                        {#if index + 1 == $braille.CurrentClass + 1}
                            {#each letters as letter}
                                <CardLetter
                                    src={letter.src}
                                    alt={letter.alt}
                                    nameLetter={letter.nameLetter}
                                />
                            {/each}
                        {/if}
                    {/each}
                </div>
            </div>
            <button
                type="button"
                class="hidden-area"
                data-bs-dismiss="offcanvas"
                aria-label="Fechar"
            ></button>
        </div>
    </div>
</nav>

<style>
    :root {
        --primary-button-rules-color: #f8cfab;
        --secondary-button-rules-color: #fabc86;
    }

    .hidden-area {
        position: absolute;
        width: 1px;
        height: 1px;
        padding: 0;
        margin: -1px;
        overflow: hidden;
        clip: rect(0, 0, 0, 0);
        border: 0;
    }
    .container-fluid {
        padding-left: 0;
    }

    .button-rules {
        display: flex;
        justify-content: center;
        align-items: center;
        width: 5rem;
        height: 5rem;
        position: fixed;
        top: 40rem;
        left: 0.3rem;
        margin: 0;
        padding: 0 !important;
        background-color: var(--setenary-color);
        border-radius: 0.8rem;
        cursor: pointer;
    }

    .button-rules:active {
        border: 1px solid var(--secondary-button-rules-color);
    }

    .offcanvas-header {
        display: flex;
        justify-content: space-between;
    }

    h2 {
        margin: 0;
        padding: 0;
        font-size: 2.8rem;
    }
    .offcanvas-body {
        display: flex;
        justify-content: center;
    }

    .container_card-letter {
        width: 40rem;
        height: 50rem;
        display: flex;
        justify-content: center;
        flex-wrap: wrap;
        margin-top: 3rem;
    }
    .button-rules_close {
        font-size: 3rem;
        padding: 2rem;
    }

    .fixed-div {
        position: fixed;
    }
    @media (hover: hover) {
        .button-rules:hover {
            background-color: var(--secondary-button-rules-color);
            /* border: none; */
        }
    }
</style>
