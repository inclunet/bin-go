<script>
    import CardLetter from "./CardLetter.svelte";
    import { braille } from "$lib/braille";
    import { lettersList } from "./lettersList";
</script>

<nav class="navbar">
    <div class="container-fluid">
        <button
            class="button-rules"
            type="button"
            data-bs-toggle="offcanvas"
            data-bs-target="#offcanvasNavbar"
            aria-controls="offcanvasNavbar"
            aria-label="Toggle navigation"
        >
            Regras
        </button>
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
                    aria-label="Close"
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
                    <!-- <CardLetter
                        src="/img/letterList/a-j/a.png"
                        alt="vazio"
                        nameLetter="a"
                    />
                    <CardLetter
                        src="/img/letterList/a-j/b.png"
                        alt="vazio"
                        nameLetter="b"
                    />
                    <CardLetter
                        src="/img/letterList/a-j/c.png"
                        alt="vazio"
                        nameLetter="c"
                    />
                    <CardLetter
                        src="/img/letterList/a-j/d.png"
                        alt="vazio"
                        nameLetter="d"
                    />

                    <CardLetter
                        src="/img/letterList/a-j/e.png"
                        alt="vazio"
                        nameLetter="e"
                    />
                    <CardLetter
                        src="/img/letterList/a-j/f.png"
                        alt="vazio"
                        nameLetter="f"
                    />
                    <CardLetter
                        src="/img/letterList/a-j/g.png"
                        alt="vazio"
                        nameLetter="g"
                    />
                    <CardLetter
                        src="/img/letterList/a-j/h.png"
                        alt="vazio"
                        nameLetter="h"
                    />
                    <CardLetter
                        src="/img/letterList/a-j/i.png"
                        alt="vazio"
                        nameLetter="i"
                    />
                    <CardLetter
                        src="/img/letterList/a-j/j.png"
                        alt="vazio"
                        nameLetter="j"
                    /> -->
                </div>
            </div>
        </div>
    </div>
</nav>

<style>
    :root {
        --primary-button-rules-color: #f8cfab;
        --secondary-button-rules-color: #fabc86;

        font-size: 62.5%;
    }
    .container-fluid {
        padding-left: 0;
    }
    .button-rules {
        position: absolute;
        margin: 0;
        font-size: 2rem;
        padding: 0.45rem 1rem;
        background-color: var(--primary-button-rules-color);
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
    @media (hover: hover) {
        .button-rules:hover {
            background-color: var(--secondary-button-rules-color);
            /* border: none; */
        }
    }
</style>
