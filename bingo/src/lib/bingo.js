import { goto } from "$app/navigation";
import { page } from "$app/stores";
import { writable } from "svelte/store";

export let card = writable({
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
});

export async function getCard(endpoint = "") {
    const response = await fetch(getEndpointUrl(endpoint));
    const data = await response.json();
    return data;
}

function getEndpointProtocol(protocol = "http") {
    if (window.location.protocol == "https:" && protocol == "ws") {
        return ("wss:");
    }

    if (window.location.protocol == "http:" && protocol == "ws") {
        return ("ws:");
    }

    return (window.location.protocol);
}

export function getEndpointUrl(call = "", protocol = "http") {
    if (window.location.port == "5173") {
        return (
            getEndpointProtocol(protocol) +
            "//" +
            window.location.hostname +
            ":8080/api" +
            call
        );
    } else {
        return getEndpointProtocol(protocol) + "//" + window.location.host + "/api" + call;
    }
}
