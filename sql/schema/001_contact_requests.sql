-- +goose up
CREATE TABLE contact_requests (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    email TEXT NOT NULL,
    text TEXT NOT NULL,,
    telephone VARCHAR(15),
    created_at TEXT NOT NULL DEFAULT (datetime('now'))
);

-- +goose Down
DROP TABLE contact_requests;