const storageKey = "inclubingo-player-cards";
const playerIDKey = "inclubingo-player-id";
const roundCreationIDKey = "inclubingo-pending-round-creation-id";
const pendingLobbyURLKey = "inclubingo-pending-organizer-lobby-url";

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

export const getPlayerID = () => {
    let playerID = localStorage.getItem(playerIDKey);
    if (!playerID) {
        playerID = crypto.randomUUID();
        localStorage.setItem(playerIDKey, playerID);
    }
    return playerID;
};

export const getRoundCreationID = () => {
    let creationID = localStorage.getItem(roundCreationIDKey);
    if (!creationID) {
        creationID = crypto.randomUUID();
        localStorage.setItem(roundCreationIDKey, creationID);
    }
    return creationID;
};

export const clearRoundCreationID = () => {
    localStorage.removeItem(roundCreationIDKey);
    localStorage.removeItem(pendingLobbyURLKey);
};

/** @param {string} lobbyURL */
export const rememberPendingLobbyURL = (lobbyURL) => {
    localStorage.setItem(pendingLobbyURLKey, lobbyURL);
};

export const getPendingLobbyURL = () =>
    localStorage.getItem(pendingLobbyURLKey) || "";
