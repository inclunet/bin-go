import { card } from "$lib/bingo";

export function load({ params }) {
    card.Card = params.card === undefined ? 0 : params.card;
    card.Round = params.round === undefined ? 0 : params.round;
    return card;
}