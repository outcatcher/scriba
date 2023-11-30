-- goose Up

CREATE TABLE players
(
    id            UUID PRIMARY KEY,
    username      VARCHAR     NOT NULL,
    tg_user_id    INT8 UNIQUE NOT NULL,
    current_count INT4        NOT NULL DEFAULT 0,
);

CREATE TABLE count_history
(
    timestamp TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    delta     SMALLINT                 NOT NULL,
    player_id UUID                     NOT NULL REFERENCES players (id),
);

-- +goose StatementBegin
CREATE FUNCTION players_count_change() RETURNS TRIGGER AS $$
BEGIN

UPDATE players
SET current_count = current_count + OLD.delta;
RETURN NEW;

END;
$$
LANGUAGE plpgsql;
-- +goose StatementEnd

CREATE TRIGGER players_count_changes
    AFTER INSERT
    ON count_history
    FOR EACH ROW
    EXECUTE players_count_change();

-- goose Down

DROP TRIGGER players_count_changes;
DROP FUNCTION players_count_change;
DROP TABLE count_history;
DROP TABLE players;
