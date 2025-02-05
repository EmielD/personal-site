-- name: InsertContactRequest :one
INSERT INTO contact_requests (name, email, text, telephone, date)
VALUES (?, ?, ?, ?, ?)
RETURNING *;
