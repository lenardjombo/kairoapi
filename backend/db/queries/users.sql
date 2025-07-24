-- name: CreateUser :one
INSERT INTO users (
  id,email ,password,created_at,updated_at
) VALUES (
  $1, $2, $3, $4, $5
)
RETURNING *;

-- name: GetUserById :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = $1 LIMIT 1;


-- name: ListUsers :many
SELECT * FROM users
ORDER BY email;


-- name: UpdateUser :execrows
UPDATE users
  set email = $2
WHERE id = $1;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;