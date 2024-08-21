CREATE TABLE IF NOT EXISTS ships (
	symbol TEXT PRIMARY KEY UNIQUE,
	current_system TEXT NOT NULL,
	current_waypoint TEXT NOT NULL
);

CREATE INDEX IF NOT EXISTS ship_symbol ON ships(symbol);
CREATE INDEX IF NOT EXISTS ship_current_systems ON ships(current_system);
