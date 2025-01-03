-- name: InsertValues :one
INSERT INTO
    test (name)
VALUES
    ('one') RETURNING *;