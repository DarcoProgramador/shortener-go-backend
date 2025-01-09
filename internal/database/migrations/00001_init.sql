-- +goose Up
-- +goose StatementBegin
CREATE TABLE urls (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    url TEXT NOT NULL,
    shortCode TEXT NOT NULL UNIQUE,
    createdAt TEXT NOT NULL,
    updatedAt TEXT NOT NULL,
    accessCount INTEGER DEFAULT 0
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS urls;
-- +goose StatementEnd
