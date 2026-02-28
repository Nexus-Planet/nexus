package message

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) CreateMessage(ctx context.Context, data CreateMessageParams) (*MessageDB, error) {
	q := `
	INSERT INTO messages(id)
	VALUES(?);
	`
	q = r.db.Rebind(q)

	_, err := r.db.NamedExecContext(ctx, q, data)
	if err != nil {
		return nil, err
	}

	msg, err := r.FindOne(ctx, data.ID)
	if err != nil {
		return nil, err
	}

	return msg, err
}

func (r *Repository) FindOne(ctx context.Context, id string) (*MessageDB, error) {
	q := `
	SELECT *
	FROM messages
	WHERE id = ?;
	`
	q = r.db.Rebind(q)

	var msg MessageDB
	err := r.db.GetContext(ctx, &msg, q)
	if err != nil {
		return nil, err
	}

	return &msg, nil
}

func (r *Repository) FindAll(ctx context.Context) ([]MessageDB, error) {
	q := `
	SELECT *
	FROM messages
	ORDER BY created_at DESC;
	`

	q = r.db.Rebind(q)

	var msgs []MessageDB
	err := r.db.GetContext(ctx, &msgs, q)
	if err != nil {
		return nil, err
	}

	return msgs, nil
}

func (r *Repository) UpdateData(ctx context.Context, data *UpdateMessageParams) (*MessageDB, error) {
	q := `
	UPDATE messages
	SET content = COALESCE($2, display_name),
    updated_at = CURRENT_TIMESTAMP
    WHERE id = ?;
	`

	q = r.db.Rebind(q)

	_, err := r.db.NamedExecContext(ctx, q, data)
	if err != nil {
		return nil, err
	}

	msg, err := r.FindOne(ctx, data.ID)
	if err != nil {
		return nil, err
	}

	return msg, nil
}

func (r *Repository) SoftDelete(ctx context.Context, id string) error {
	q := `
	UPDATE messages
	SET status = 'pending_delete',
    deleted_at = CURRENT_TIMESTAMP
    WHERE id = ?;
	`

	q = r.db.Rebind(q)

	_, err := r.db.NamedExecContext(ctx, q, id)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) TogglePin(ctx context.Context, data *TogglePinParams) error {
	q := `
	UPDATE messages_guilds
	SET is_pinned = $3
	WHERE message_id = $1 AND guild_id = $2;
	`

	q = r.db.Rebind(q)

	_, err := r.db.NamedExecContext(ctx, q, data)
	if err != nil {
		return err
	}

	return nil
}
