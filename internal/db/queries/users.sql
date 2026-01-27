-- Create User  
-- name: CreateUser :exec
INSERT INTO users(id, email, password_hash,is_active) VALUES($1, $2, $3,1);

-- Get User Data
-- name: FindOneUser :one
SELECT * FROM users WHERE id = $1; 

-- Get User Data By Email
-- name: FindOneUserByEmail :one
SELECT * FROM users WHERE email = $1;

-- Get All Users With Data
-- name: FindAllUsers :many
SELECT * FROM users ORDER BY created_at DESC; 

-- Update User Date
UPDATE users SET display_name =  COALESCE($1,display_name), username =  COALESCE($2,username), email = COALESCE($3,email);

-- Update Profile Picture

-- Soft Delete User Account
-- name: SoftDeleteUser :exec
UPDATE users SET is_active = 0, deleted_at = CURRENT_TIMESTAMP WHERE id = $1;

-- Delete User Account
-- name: DeleteUser :exec
DELETE FROM users WHERE deleted_at IS NOT NULL AND deleted_after <= $1;

-- Cancel Delete Account
-- name: CancelDeleteUser :exec
UPDATE users SET is_active = 1, deleted_at = NULL WHERE id = $1;

-- Deactivate User Account
-- name: DeactivateUser :exec
UPDATE users SET is_active = COALESCE($1,is_active) WHERE id = $1;