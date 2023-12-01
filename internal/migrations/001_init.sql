-- goose Up

CREATE TABLE players
(
    id         UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    username   VARCHAR NOT NULL,
    tg_user_id INT8    NOT NULL UNIQUE,
);

CREATE TABLE count_history
(
    timestamp TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    delta     SMALLINT                 NOT NULL,
    player_id UUID                     NOT NULL REFERENCES players (id),
);

-- goose Down

DROP INDEX count_history_player_id_idx;
DROP TABLE count_history;
DROP TABLE players;
