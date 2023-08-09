import { goto } from "$app/navigation";
import { page } from "$app/stores";

export let audio;

export let card = {
    Autoplay: false,
    Bingo: false,
    Card: 0,
    Checked: 0,
    LastNumber: 0,
    NextRound: 0,
    Round: 0,
    Type: 75,
    Numbers: [
        [
            { Checked: false, Number: 0, },
        ],
    ],
};

export async function sendCheckNumber(number = 0) {
    return await getCard("/card/" + card.Round + "/" + card.Card + "/" + number);
}

export async function drawNumber() {
    card = await getCard("/card/" + card.Round + "/1/0");
}

export async function getCard(endpoint = "") {
    const response = await fetch(getEndpointUrl(endpoint));
    return await response.json();
}

export function getEndpointUrl(call = "") {
    if (window.location.port == "5173") {
        return (
            window.location.protocol +
            "//" +
            window.location.hostname +
            ":8080/api" +
            call
        );
    } else {
        return window.location.protocol + "//" + window.location.host + "/api" + call;
    }
}

export function redirectToNextRound() {
    if (card.NextRound > 0) {
        goto("/card/" + card.NextRound + "/0");
    }
}

export async function newGame() {
    card = await getCard("/card/" + card.Round + "/new/" + card.Type);
    goto("/card/" + card.Round);
}

export function playAudio() {
    if (card.Bingo) {
        audio.play();
    }
}

export function startNewGame() {
    goto("/card/" + card.Round + "/1");
}

export async function toggleAutoplay() {
    if (card.Autoplay) {
        card.Autoplay = false;
    } else {
        card.Autoplay = true;
    }

    card = await getCard("/card/" + card.Round + "/" + card.Card + "/autoplay");
}

