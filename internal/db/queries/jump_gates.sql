-- name: InsertJumpGate :exec
INSERT INTO jump_gates (
	system, waypoint, connects_to_wp, connects_to_sys
) VALUES (
	?, ?, ?, ?
);

-- name: TruncateJumpGates :exec
DELETE FROM jump_gates;

-- name: GetConnectedSystemNames :many
	connects_to_sys
FROM jump_gates
WHERE system = sqlc.arg(system_name)
;
