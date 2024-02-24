-- +goose Up
DROP TABLE IF EXISTS users;



-- +goose Down
DROP TABLE users;

