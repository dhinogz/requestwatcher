-- name: CreateWatcher :one
INSERT INTO watchers (hash_uri)
VALUES ($1)
RETURNING watcher_id;

-- name: GetWatcher :one
SELECT watcher_id
FROM watchers w
WHERE w.hash_uri = $1;

