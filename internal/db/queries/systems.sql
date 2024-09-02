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
	  sqlc.embed(systems)
	, CAST(IFNULL(GROUP_CONCAT(wp_connected.system, ','), '') AS TEXT) AS connected_systems
FROM systems
JOIN waypoints
	ON systems.symbol = waypoints.system
LEFT JOIN jump_gates
	ON waypoints.symbol = jump_gates.waypoint
LEFT JOIN waypoints wp_connected
	ON wp_connected.symbol = jump_gates.connects_to
WHERE TRUE
	AND systems.x >= sqlc.arg(x_min) AND systems.x <= sqlc.arg(x_max)
	AND systems.y >= sqlc.arg(y_min) AND systems.y <= sqlc.arg(y_max)
GROUP BY
	  systems.symbol
	, systems.x
	, systems.y
	, systems.type
	, factions
;

-- name: GetSystemByName :one
SELECT x, y FROM systems
WHERE symbol = sqlc.arg(system_name);
