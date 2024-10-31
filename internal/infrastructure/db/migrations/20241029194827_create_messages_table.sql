-- +goose Up
-- +goose StatementBegin
CREATE TABLE messages (
    id INTEGER PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    chat_id INTEGER REFERENCES chats(id) ON DELETE CASCADE,
    from_user_id INTEGER NOT NULL,
    text TEXT NOT NULL,
    timestamp TIMESTAMP NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE messages;
-- +goose StatementEnd
