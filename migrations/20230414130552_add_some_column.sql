-- +goose Up
-- +goose StatementBegin
CREATE TABLE user (
   id INT(11) NOT NULL AUTO_INCREMENT,
   name VARCHAR(50) NOT NULL,
   email VARCHAR(100) NOT NULL,
   created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
   updated_at TIMESTAMP NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
   PRIMARY KEY (id)
);

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS user;

-- +goose StatementEnd