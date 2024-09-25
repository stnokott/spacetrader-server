-- name: InsertJumpGate :exec
INSERT INTO jump_gates (
	waypoint, connects_to
) VALUES (
	?, ?
);

-- name: TruncateJumpGates :exec
DELETE FROM jump_gates;

-- name: GetConnectionsForWaypoint :many
SELECT
	connects_to AS connected_wp
FROM jump_gates
WHERE waypoint = sqlc.arg(waypoint);
