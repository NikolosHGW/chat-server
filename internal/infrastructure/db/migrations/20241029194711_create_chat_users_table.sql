-- +goose Up
-- +goose StatementBegin
CREATE TABLE chat_users (
    id INTEGER PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    chat_id INTEGER REFERENCES chats(id) ON DELETE CASCADE,
    user_id INTEGER NOT NULL,
    UNIQUE (chat_id, user_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE chat_users;
-- +goose StatementEnd
