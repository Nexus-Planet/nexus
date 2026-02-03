-- Create User  
-- name: CreateSession :exec
INSERT INTO auth_sessions(id, email, password_hash)
VALUES($1, $2, $3);
-- Get User Data
-- name: FindOneSession :one
SELECT *
FROM auth_sessions
WHERE id = $1;
-- Get User Data By Email
-- name: FindOneSessionByEmail :one
SELECT *
FROM auth_sessions
WHERE email = $1;
-- Get All Users With Data
-- name: FindAllSessions :many
SELECT *
FROM auth_sessions
ORDER BY created_at DESC;
-- Update User Date
-- name: UpdatePassword :exec
UPDATE auth_sessions
SET password_hash = $2,
    updated_at = CURRENT_TIMESTAMP
WHERE id = $1;
-- Delete User Account
-- name: DeleteSession :exec
DELETE FROM auth_sessions
WHERE id = $1;
-- Toggle Deactivate User Account
-- name: ToggleDeactivateSession :exec
UPDATE auth_sessions
SET is_active = COALESCE($1, is_active)
WHERE id = $1;