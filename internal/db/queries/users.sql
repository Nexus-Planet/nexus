-- Create User  
-- name: CreateUser :exec
INSERT INTO users(id,email,password_hash) VALUES($1,$2,$3);

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


-- Update Profile Picture

-- Delete User Account
-- name: DeleteUser :exec
DELETE FROM users WHERE id = $1;

-- Deactivate User Account
-- name: DeactivateUser :exec
