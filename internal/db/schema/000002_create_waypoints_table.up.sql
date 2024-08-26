CREATE TABLE IF NOT EXISTS waypoints (
	symbol TEXT PRIMARY KEY UNIQUE,
	system TEXT NOT NULL,
	orbits TEXT NULL,
	x INTEGER NOT NULL,
	y INTEGER NOT NULL,
	type TEXT NOT NULL,
	charted BOOL NOT NULL,
	FOREIGN KEY(system) REFERENCES systems(symbol)
);

CREATE INDEX IF NOT EXISTS waypoint_system_index ON waypoints(system);
CREATE INDEX IF NOT EXISTS waypoint_orbits_index ON waypoints(orbits);
