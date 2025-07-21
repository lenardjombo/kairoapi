-- name: CreateUser :one
INSERT INTO users (
  email ,hashedpassword
) VALUES (
  $1, $2
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


-- name: UpdateUser :exec
UPDATE users
  set email = $2
WHERE id = $1;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;