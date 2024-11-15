-- name: HasSystemsRows :one
SELECT EXISTS (SELECT 1 FROM systems) AS "exists";

-- name: InsertSystem :exec
INSERT INTO systems (
	symbol, x, y, type, factions	
) VALUES (
	?, ?, ?, ?, ?
);

-- name: TruncateSystems :exec
DELETE FROM systems;

-- name: GetSystemCount :one
SELECT
	COUNT(*) AS n
FROM systems
;

-- name: GetSystemsOffset :many
SELECT
	*
FROM systems
ORDER BY symbol
LIMIT sqlc.arg(limit) OFFSET sqlc.arg(offset)
;

-- name: GetSystemsByName :many
SELECT * FROM systems
WHERE symbol IN (sqlc.slice(system_ids))
;
