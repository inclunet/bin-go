<script>
    import OneClickButton from "$lib/atoms/OneClickButton.svelte";
    import { onMount } from "svelte";

    /**
     * @type {string[][]}
     */
    export let card = [];

    /**
     * @type {OneClickButton[][]}
     */
    let cardButtons = [[]];

    let x = 0;
    let y = 0;
    const moveCellHandler = (
        /** @type {{ getAttribute: (arg0: string) => string; key: any; }} */ e,
    ) => {
        // x = parseInt(e.getAttribute("data-x"));
        // y = parseInt(e.getAttribute("data-y"));

        switch (e.key) {
            case "ArrowUp":
                if (x > 0) {
                    x--;
                }
                break;

            case "ArrowDown":
                if (x < cardButtons.length - 1) {
                    x++;
                }
                break;

            case "ArrowLeft":
                if (y > 0) {
                    y--;
                }
                break;

            case "ArrowRight":
                if (y < cardButtons[x].length - 1) {
                    y++;
                }
                break;

            case "Home":
                y = 0;
                break;

            case "End":
                y = cardButtons.length - 1;
                break;

            case "PageUp":
                x = 0;
                break;

            case "PageDown":
                x = cardButtons[x].length - 1;
                break;
        }

        cardButtons[x][y].focus();
    };

    onMount(() => {
        cardButtons = card.map((row, i) =>
            row.map((cell, j) => {
                let button = OneClickButton;
                button.setAttribute("data-x", "{i}");
                button.setAttribute("data-y", "{j}");
                return button;
            }),
        );
        cardTable.setAttribute("tabindex", "0");
        cardTable.focus();
    });
</script>

<p>teste</p>

<table
    bind:this={cardTable}
    role="grid"
    tabindex="0"
    on:keyup={moveCellHandler}
>
    <tbody>
        {#each card as row, i}
            <tr>
                {#each row as cell, j}
                    <td role="gridcell">
                        <OneClickButton
                            bind:this={cardButtons[x][y]}
                            x={i}
                            y={j}>{cell}</OneClickButton
                        >
                    </td>
                {/each}
            </tr>
        {/each}
    </tbody>
</table>
