-- name: GetTodoById :one
SELECT * from todos 
WHERE id = $1 
AND deleted_at is NULL
LIMIT 1;

-- name: GetTodo :many
SELECT * from todos 
WHERE deleted_at is NULL
LIMIT $1 OFFSET $2;

-- name: CreateTodo :one
INSERT INTO todos (
    description, is_completed, created_at, updated_at
) VALUES (
    $1, $2, $3, $4
) 
RETURNING *;

-- name: UpdateTodo :one
UPDATE todos
SET description = COALESCE(sqlc.arg(description), description), is_completed = COALESCE(sqlc.arg(is_completed), is_completed)
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: DeleteTodo :exec
UPDATE todos
set deleted_at = now()
where id = $1;

