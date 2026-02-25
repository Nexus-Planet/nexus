-- Create User  
-- name: CreateUser :exec
INSERT INTO users(id)
VALUES($1);
-- Get User Data
-- name: FindOneUser :one
SELECT *
FROM users
WHERE id = $1;
-- Get All Users With Data
-- name: FindAllUsers :many
SELECT *
FROM users
ORDER BY created_at DESC;
-- Update User Data
-- name: UpdateUser :exec
UPDATE users
SET display_name = COALESCE($2, display_name),
    updated_at = CURRENT_TIMESTAMP
WHERE id = $1;
-- Set Username
-- name: SetUserName :exec
UPDATE users
SET username = $2,
    updated_at = CURRENT_TIMESTAMP
WHERE id = $1
    AND (
        username IS NULL
        OR username_changed_at IS NULL
        OR username_changed_at <= $3
    );
-- Soft Delete User Account
-- name: SoftDeleteUser :exec
UPDATE users
SET status = 'pending_delete',
    deleted_at = CURRENT_TIMESTAMP
WHERE id = $1;
-- Delete User Account
-- name: DeleteUser :exec
DELETE FROM users
WHERE deleted_at IS NOT NULL
    AND deleted_after <= $1;
-- Cancel Delete Account
-- name: CancelDeleteUser :exec
UPDATE users
SET account_status = 'active',
    deleted_at = NULL
WHERE id = $1;
-- Deactivate User Account
-- name: DeactivateUser :exec
UPDATE users
SET account_status = 'deactivated'
WHERE id = $1;
-- Reactivate User Account
-- name: ReactivateUser :exec
UPDATE Users
SET account_status = 'active'
WHERE id = $1;