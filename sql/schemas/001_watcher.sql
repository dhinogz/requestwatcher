-- +goose Up
CREATE TABLE watchers (
    watcher_id BIGSERIAL PRIMARY KEY,

    -- TODO: make index
    hash_uri TEXT NOT NULL UNIQUE,

    created_at TIMESTAMP(0) with time zone NOT NULL DEFAULT NOW(),
    active BOOLEAN NOT NULL DEFAULT TRUE
);

-- +goose Down
DROP TABLE watchers;

