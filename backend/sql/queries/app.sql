-- name: CreateWorker :one
INSERT INTO auth_db (id, worker_group, api_key_hash)
VALUES ($1, $2, $3)
RETURNING *;
