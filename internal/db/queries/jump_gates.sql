-- name: InsertJumpGate :exec
INSERT INTO jump_gates (
	waypoint, connects_to
) VALUES (
	?, ?
);

-- name: TruncateJumpGates :exec
DELETE FROM jump_gates;
