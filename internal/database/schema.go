package database

var DatabaseSchema = `
CREATE TABLE IF NOT EXISTS users (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT NOT NULL,
	role TEXT NOT NULL DEFAULT 'USER',
	email TEXT NOT NULL UNIQUE,
	age INTEGER,
	password TEXT NOT NULL
);
`
