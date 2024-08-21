-- name: InsertShip :exec
INSERT INTO ships (
		symbol,	current_system, current_waypoint
) VALUES (
	?, ?, ?
);

-- name: TruncateShips :exec
DELETE FROM ships;
