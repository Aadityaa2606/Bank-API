-- +goose Up
-- +goose StatementBegin
CREATE TABLE "sessions"  (
    "id" varchar PRIMARY KEY,
    "username" varchar NOT NULL,
    "refresh_token" varchar(512) NOT NULL,
    "is_revoked" boolean NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT (now()),
    "expires_at" timestamptz NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "sessions";
-- +goose StatementEnd
