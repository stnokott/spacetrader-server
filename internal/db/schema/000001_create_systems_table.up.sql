CREATE TABLE IF NOT EXISTS systems (
	symbol TEXT PRIMARY KEY UNIQUE,
	x INTEGER NOT NULL,
	y INTEGER NOT NULL,
	type TEXT NOT NULL,
	factions TEXT NOT NULL
);

CREATE INDEX IF NOT EXISTS system_x_pos_index ON systems(x);
CREATE INDEX IF NOT EXISTS system_y_pos_index ON systems(y);
CREATE INDEX IF NOT EXISTS system_type_index ON systems(type);
