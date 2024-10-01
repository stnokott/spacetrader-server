-- name: InsertJumpGate :exec
INSERT INTO jump_gates (
	waypoint, connects_to
) VALUES (
	?, ?
);

-- name: TruncateJumpGates :exec
DELETE FROM jump_gates;

-- name: HasJumpgateRows :one
SELECT EXISTS (SELECT 1 FROM jump_gates) AS "exists";

-- name: GetJumpgatesInSystem :many
SELECT
	*
FROM jump_gates
WHERE waypoint IN (
	SELECT symbol
	FROM waypoints
	WHERE system = sqlc.arg(system_id)
);
