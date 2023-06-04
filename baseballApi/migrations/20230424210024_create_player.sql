-- +goose Up
-- +goose StatementBegin
CREATE TABLE players (
   id VARCHAR(255) PRIMARY KEY,
   uniform_number INT NOT NULL,
   name VARCHAR(255) NOT NULL,
   at_bats INT NOT NULL,
   hits INT NOT NULL,
   walks INT NOT NULL,
   home_runs INT NOT NULL,
   runs_batted_in INT NOT NULL,
   created_at TIMESTAMP NOT NULL,
   updated_at TIMESTAMP NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS players;
-- +goose StatementEnd