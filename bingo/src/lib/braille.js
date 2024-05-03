import { writable } from "svelte/store";

export let braille = writable({
    Description: "",
    CurrentClass: 0,
        CurrentPunctuation: 0,
    CurrentRound: 0,
    TotalPunctuation: 0,
    Challenge: {
        Braille: [],
        Word: ""
    },
});
