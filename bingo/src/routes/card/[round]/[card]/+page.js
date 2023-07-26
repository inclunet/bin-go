export function load({ params }) {
    return {
        Card: params.card,
        Checked: 0,
        Name: "",
        Round: params.round,
        Numbers: [],
    };
}