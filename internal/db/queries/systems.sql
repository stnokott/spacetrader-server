-- name: InsertSystem :exec
INSERT INTO systems (
	symbol, x, y, type, factions	
) VALUES (
	?, ?, ?, ?, ?
);

-- name: GetSystemsInRect :many
SELECT
	sqlc.embed(systems), COUNT(ships.symbol) AS ship_count
FROM systems
LEFT JOIN ships
	ON ships.current_system = systems.symbol
WHERE TRUE
	AND x >= sqlc.arg(x_min) AND x <= sqlc.arg(x_max)
	AND y >= sqlc.arg(y_min) AND y <= sqlc.arg(y_max)
GROUP BY systems.symbol, x, y, type, factions
;

-- name: TruncateSystems :exec
DELETE FROM systems;

-- name: GetSystemByName :one
SELECT x, y FROM systems
WHERE symbol = sqlc.arg(system_name);
