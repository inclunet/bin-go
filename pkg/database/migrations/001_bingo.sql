CREATE TABLE bingo_rounds (
    id UUID PRIMARY KEY,
    display_number BIGINT NOT NULL UNIQUE,
    type INTEGER NOT NULL CHECK (type > 0),
    next_round_id UUID REFERENCES bingo_rounds(id),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE bingo_cards (
    id UUID PRIMARY KEY,
    round_id UUID NOT NULL REFERENCES bingo_rounds(id) ON DELETE CASCADE,
    display_number INTEGER NOT NULL CHECK (display_number > 0),
    state JSONB NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    UNIQUE (round_id, display_number)
);

CREATE INDEX bingo_cards_round_id_idx ON bingo_cards(round_id);

CREATE TABLE bingo_draws (
    round_id UUID NOT NULL REFERENCES bingo_rounds(id) ON DELETE CASCADE,
    number INTEGER NOT NULL CHECK (number > 0),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    PRIMARY KEY (round_id, number)
);
