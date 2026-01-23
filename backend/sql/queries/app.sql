-- name: CreateWorker :one
INSERT INTO auth_db (id, worker_group, api_key_hash,is_disabled,expiry)
VALUES ($1, $2, $3, $4 , $5)
RETURNING *;
