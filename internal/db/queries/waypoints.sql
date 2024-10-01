-- name: InsertWaypoint :exec
INSERT INTO waypoints (
	symbol, system, x, y, orbits, type
) VALUES (
	?, ?, ?, ?, ?, ?
);

-- name: TruncateWaypoints :exec
DELETE FROM waypoints;

-- name: GetWaypointsByType :many
SELECT * FROM waypoints
WHERE type = sqlc.arg(type);

-- name: GetWaypointsForSystem :many
SELECT * FROM waypoints
WHERE system = sqlc.arg(system_name);

-- name: GetWaypointsByName :many
SELECT * FROM waypoints
WHERE symbol IN (sqlc.slice(waypoint_ids));
