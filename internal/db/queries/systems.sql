-- name: InsertSystem :exec
INSERT INTO systems (
	symbol, x, y, type, factions	
) VALUES (
	?, ?, ?, ?, ?
);

-- name: SelectSystemsInRect :many
SELECT symbol, x, y, type, factions FROM systems
	WHERE TRUE
		AND x >= sqlc.arg(x_min) AND x <= sqlc.arg(x_max)
		AND y >= sqlc.arg(y_min) AND y <= sqlc.arg(y_max);

-- name: TruncateSystems :exec
DELETE FROM systems;
