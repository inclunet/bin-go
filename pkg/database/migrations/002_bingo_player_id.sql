ALTER TABLE bingo_cards
    ADD COLUMN IF NOT EXISTS player_id UUID;

CREATE UNIQUE INDEX IF NOT EXISTS bingo_cards_round_player_idx
    ON bingo_cards(round_id, player_id)
    WHERE player_id IS NOT NULL;
