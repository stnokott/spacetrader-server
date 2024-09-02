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

-- name: GetAllSystems :many
SELECT
	  systems.symbol AS name
	, systems.x AS x
	, systems.y AS y
	, COUNT(jump_gates.waypoint) > 0 AS has_jumpgates
FROM systems
JOIN waypoints
	ON systems.symbol = waypoints.system
LEFT JOIN jump_gates
	ON waypoints.symbol = jump_gates.waypoint
GROUP BY
	  systems.symbol
	, systems.x
	, systems.y
;

-- name: GetSystemByName :one
SELECT x, y FROM systems
WHERE symbol = sqlc.arg(system_name);
