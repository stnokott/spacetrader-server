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

-- name: GetSystemsInRect :many
SELECT
	symbol, x, y, type, factions
FROM systems
WHERE TRUE
	AND x >= sqlc.arg(x_min) AND x <= sqlc.arg(x_max)
	AND y >= sqlc.arg(y_min) AND y <= sqlc.arg(y_max)
;

-- name: GetSystemByName :one
SELECT x, y FROM systems
WHERE symbol = sqlc.arg(system_name);
