import { card } from "../../../bingo";

export function load({ params }) {
    card.Card = params.card;
    card.Round = params.round;
    return card;
}