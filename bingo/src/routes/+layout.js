export function load({ params }) {
    return ({
        Card: params.card === undefined ? 0 : params.card,
        Round: params.round === undefined ? 0 : params.round,
    });
}