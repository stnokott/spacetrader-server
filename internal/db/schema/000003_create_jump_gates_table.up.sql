CREATE TABLE IF NOT EXISTS jump_gates (
	waypoint TEXT NOT NULL,
	connects_to TEXT NOT NULL,
	FOREIGN KEY(waypoint) REFERENCES waypoints(symbol),
	FOREIGN KEY(connects_to) REFERENCES waypoints(symbol) DEFERRABLE INITIALLY DEFERRED -- defer since waypoint might not exist yet (TODO: refactor system cache creation so that waypoints can be created on-demand and we dont need this anymore)
);

CREATE INDEX IF NOT EXISTS jump_gate_waypoint_index ON jump_gates(waypoint);
CREATE INDEX IF NOT EXISTS jump_gate_connects_to_index ON jump_gates(connects_to);
