ALTER TABLE bingo_rounds
    ADD COLUMN IF NOT EXISTS creation_id UUID;

CREATE UNIQUE INDEX IF NOT EXISTS bingo_rounds_creation_id_idx
    ON bingo_rounds(creation_id)
    WHERE creation_id IS NOT NULL;
