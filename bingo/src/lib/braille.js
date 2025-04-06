import { writable } from "svelte/store";

export let braille = writable({
    Description: "",
    CurrentClass: 0,
    CurrentPunctuation: 0,
    CurrentRound: 0,
    Player: -1,
    TotalPunctuation: 0,
    Challenge: {
        Braille: [],
        Word: "",
    },
    TotalRounds: 0,
});
