export function load({ params }) {
    return ({
        Card: params.card === undefined ? 0 : params.card,
        Class: params.class === undefined ? 0 : params.class,
        Player: params.player === undefined ? 0 : params.player,
        Round: params.round === undefined ? 0 : params.round,
        });
}