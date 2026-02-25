-- Create message
-- name: CreateMessage :exec
INSERT INTO messages(id)
VALUES($1);
-- Get message Data
-- name: FindOneMessage :one
SELECT *
FROM messages
WHERE id = $1;
-- Get All messages With Data
-- name: FindAllmessages :many
SELECT *
FROM messages
ORDER BY created_at DESC;
-- Update message Data
-- name: UpdateMessage :exec
UPDATE messages
SET content = COALESCE($2, display_name),
    updated_at = CURRENT_TIMESTAMP
WHERE id = $1;
-- Toggle Pin Message
-- name: TogglePinMessage :exec
UPDATE messages_guilds
SET is_pinned = $3
WHERE message_id = $1 AND guild_id = $2;
-- Soft Delete message
-- name: SoftDeleteMessage :exec
UPDATE messages
SET status = 'pending_delete',
    deleted_at = CURRENT_TIMESTAMP
WHERE id = $1;
-- Delete message
-- name: DeleteMessage :exec
DELETE FROM messages
WHERE deleted_at IS NOT NULL
    AND deleted_after <= $1;
