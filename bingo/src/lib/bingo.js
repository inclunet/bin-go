import { writable } from "svelte/store";

export let card = writable({
    ID: "",
    RoundID: "",
    Autoplay: false,
    Bingo: false,
    Card: 0,
    Checked: 0,
    Completions: {},
    LastNumber: 0,
    NextRound: 0,
    NextRoundID: "",
    Round: 0,
    Type: 75,
    Numbers: [
        [
            { Checked: false, Number: 0, },
        ],
    ],
});
