-- create backup
CREATE TABLE jump_gates_backup AS SELECT * FROM jump_gates;

-- drop old table
DROP TABLE jump_gates;

-- create new table
CREATE TABLE jump_gates (
	system TEXT NOT NULL,
	waypoint TEXT NOT NULL,
	connects_to_sys TEXT NOT NULL,
	connects_to_wp TEXT NOT NULL,
	FOREIGN KEY(system) REFERENCES systems(symbol),
	FOREIGN KEY(waypoint) REFERENCES waypoints(symbol),
	FOREIGN KEY(connects_to_sys) REFERENCES systems(symbol) DEFERRABLE INITIALLY DEFERRED,
	FOREIGN KEY(connects_to_wp) REFERENCES waypoints(symbol) DEFERRABLE INITIALLY DEFERRED
);

-- migrate existing data
INSERT INTO jump_gates (
	system,
	waypoint,
	connects_to_sys,
	connects_to_wp
)
SELECT
	wp_from.system,
	wp_from.symbol,
	wp_to.system,
	wp_to.symbol
FROM jump_gates_backup src
INNER JOIN waypoints wp_from
	ON wp_from.symbol = src.waypoint
INNER JOIN waypoints wp_to
	ON wp_to.symbol = src.connects_to
;

-- drop backup
DROP TABLE jump_gates_backup;
