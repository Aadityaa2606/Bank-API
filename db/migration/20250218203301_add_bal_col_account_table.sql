-- +goose Up
-- +goose StatementBegin
ALTER TABLE accounts ADD COLUMN balance bigint NOT NULL DEFAULT 0 CHECK (balance >= 0);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE accounts DROP COLUMN balance;
-- +goose StatementEnd
