import { goto } from "$app/navigation";
import { page } from "$app/stores";

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

export function getWsEndpointUrl(call = "") {
    if (window.location.port == "5173") {
        return (
            "ws://" +
            window.location.hostname +
            ":8080/api" +
            call
        );
    } else {
        return "wss://" + window.location.host + "/api" + call;
    }
}

