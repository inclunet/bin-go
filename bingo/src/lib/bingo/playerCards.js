const storageKey = "inclubingo-player-cards";
const playerIDKey = "inclubingo-player-id";
const roundCreationIDKey = "inclubingo-pending-round-creation-id";
const pendingLobbyURLKey = "inclubingo-pending-organizer-lobby-url";
const pendingLobbySourceCardKey = "inclubingo-pending-lobby-source-card";
const memoryStorage = new Map();

/** @param {string} key */
const readStorage = (key) => {
    try {
        return localStorage.getItem(key) ?? memoryStorage.get(key) ?? null;
    } catch {
        return memoryStorage.get(key) ?? null;
    }
};

/** @param {string} key @param {string} value */
const writeStorage = (key, value) => {
    memoryStorage.set(key, value);
    try {
        localStorage.setItem(key, value);
    } catch {
        // Browser storage is a best-effort cache.
    }
};

/** @param {string} key */
const removeStorage = (key) => {
    memoryStorage.delete(key);
    try {
        localStorage.removeItem(key);
    } catch {
        // Browser storage is a best-effort cache.
    }
};

const createUUID = () => {
    if (typeof crypto !== "undefined" && crypto.randomUUID) {
        return crypto.randomUUID();
    }

    const bytes = new Uint8Array(16);
    if (typeof crypto !== "undefined" && crypto.getRandomValues) {
        crypto.getRandomValues(bytes);
    } else {
        for (let i = 0; i < bytes.length; i++) {
            bytes[i] = Math.floor(Math.random() * 256);
        }
    }
    bytes[6] = (bytes[6] & 0x0f) | 0x40;
    bytes[8] = (bytes[8] & 0x3f) | 0x80;
    const hex = Array.from(bytes, (byte) =>
        byte.toString(16).padStart(2, "0")
    ).join("");
    return `${hex.slice(0, 8)}-${hex.slice(8, 12)}-${hex.slice(
        12,
        16
    )}-${hex.slice(16, 20)}-${hex.slice(20)}`;
};

/** @returns {Record<string, string>} */
const loadCards = () => {
    try {
        const cards = JSON.parse(readStorage(storageKey) || "{}");
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
    writeStorage(storageKey, JSON.stringify(cards));
};

/** @param {string} roundID */
export const forgetPlayerCard = (roundID) => {
    const cards = loadCards();
    delete cards[roundID];
    writeStorage(storageKey, JSON.stringify(cards));
};

export const getPlayerID = () => {
    let playerID = readStorage(playerIDKey);
    if (!playerID) {
        playerID = createUUID();
        writeStorage(playerIDKey, playerID);
    }
    return playerID;
};

export const getRoundCreationID = () => {
    let creationID = readStorage(roundCreationIDKey);
    if (!creationID) {
        creationID = createUUID();
        writeStorage(roundCreationIDKey, creationID);
    }
    return creationID;
};

export const clearRoundCreationID = () => {
    removeStorage(roundCreationIDKey);
    clearPendingLobbyURL();
};

/**
 * @param {string} lobbyURL
 * @param {string} sourceCardID
 */
export const rememberPendingLobbyURL = (lobbyURL, sourceCardID = "") => {
    writeStorage(pendingLobbyURLKey, lobbyURL);
    if (sourceCardID) {
        writeStorage(pendingLobbySourceCardKey, sourceCardID);
    } else {
        removeStorage(pendingLobbySourceCardKey);
    }
};

export const getPendingLobbyURL = () =>
    readStorage(pendingLobbyURLKey) || "";

export const getPendingLobbySourceCardID = () =>
    readStorage(pendingLobbySourceCardKey) || "";

export const clearPendingLobbyURL = () => {
    removeStorage(pendingLobbyURLKey);
    removeStorage(pendingLobbySourceCardKey);
};
