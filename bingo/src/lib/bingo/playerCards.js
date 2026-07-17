const storageKey = "inclubingo-player-cards";

/** @returns {Record<string, string>} */
const loadCards = () => {
    try {
        const cards = JSON.parse(localStorage.getItem(storageKey) || "{}");
        return cards && typeof cards === "object" && !Array.isArray(cards)
            ? cards
            : {};
    } catch {
        return {};
    }
};

/** @param {string} roundID */
export const getPlayerCardID = (roundID) => {
    const cardID = loadCards()[roundID];
    return typeof cardID === "string" ? cardID : "";
};

/**
 * @param {string} roundID
 * @param {string} cardID
 */
export const rememberPlayerCard = (roundID, cardID) => {
    if (!roundID || !cardID) {
        return;
    }
    const cards = loadCards();
    cards[roundID] = cardID;
    localStorage.setItem(storageKey, JSON.stringify(cards));
};

/** @param {string} roundID */
export const forgetPlayerCard = (roundID) => {
    const cards = loadCards();
    delete cards[roundID];
    localStorage.setItem(storageKey, JSON.stringify(cards));
};
