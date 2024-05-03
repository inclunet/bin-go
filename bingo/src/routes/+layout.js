export function load({ params }) {
    return ({
        Card: params.card === undefined ? 0 : params.card,
        Player: params.round === undefined ? 0 : params.player,
        Round: params.round === undefined ? 0 : params.round,
        });
}