<script>
    import CardLetter from "./CardLetter.svelte";
    import { braille } from "$lib/braille";
    import { lettersList } from "./lettersList";
    import { HelpIcon } from "./icons";
    import MenuHelperBody from "./MenuHelperBody.svelte";
</script>

<MenuHelperBody
    target="alphabetList"
    aria_label="Regras da composição das letras em braille"
    title="Letras da Série {$braille.CurrentClass + 1}"
    icon={HelpIcon(50)}
>
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
</MenuHelperBody>

<!-- <nav class="navbar">
    <div class="container-fluid">
        <div class="fixed-div">
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
</nav> -->
