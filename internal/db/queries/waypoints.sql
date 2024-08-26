-- name: InsertWaypoint :exec
INSERT INTO waypoints (
	symbol, system, x, y, orbits, type, charted
) VALUES (
	?, ?, ?, ?, ?, ?, ?
);

-- name: TruncateWaypoints :exec
DELETE FROM waypoints;
