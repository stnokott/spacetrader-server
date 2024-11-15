-- name: InsertWaypoint :exec
INSERT INTO waypoints (
	symbol, system, x, y, orbits, type
) VALUES (
	?, ?, ?, ?, ?, ?
);

-- name: TruncateWaypoints :exec
DELETE FROM waypoints;
