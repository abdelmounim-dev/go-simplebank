-- name: CreateEntry :one
INSERT INTO entries ( account_id, amount, created_at)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetEntry :one
SELECT * from entries
where id = $1 LIMIT 1;

-- name: ListEntries :many
SELECT * FROM entries
order by id
LIMIT $1
offset $2;
