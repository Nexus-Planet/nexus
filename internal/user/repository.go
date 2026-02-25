package user

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

// Query to insert a new user
func (r *Repository) CreateUser(ctx context.Context, id string) (*UserDB, error) {
	q := `
	INSERT INTO users (id)
	VALUES (?);
	`

	q = r.db.Rebind(q)

	_, err := r.db.ExecContext(ctx, q, id)
	if err != nil {
		return nil, err
	}

	user, err := r.FindOne(ctx, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// Query to fetch one user by id
func (r *Repository) FindOne(ctx context.Context, id string) (*UserDB, error) {
	q := `
	SELECT *
	FROM users
	WHERE id = ?;
	`

	q = r.db.Rebind(q)

	var user UserDB
	err := r.db.GetContext(ctx, &user, q, id)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// Query to fetch all users
func (r *Repository) FindAll(ctx context.Context) ([]*UserDB, error) {
	q := `
	SELECT *
	FROM users
	ORDER BY created_at DESC;
	`

	q = r.db.Rebind(q)

	var users []*UserDB
	err := r.db.SelectContext(ctx, &users, q)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (r *Repository) SetUsername(ctx context.Context, params *SetUsernameParams) (*UserDB, error) {
	q := `
	UPDATE users
	SET username = ?,
    updated_at = CURRENT_TIMESTAMP
	WHERE id = ?
    AND (
        username IS NULL
        OR username_changed_at IS NULL
        OR username_changed_at <= ?
    );
	`

	q = r.db.Rebind(q)

	_, err := r.db.NamedExecContext(ctx, q, params)
	if err != nil {
		return nil, err
	}

	user, err := r.FindOne(ctx, params.ID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *Repository) UpdateData(ctx context.Context, params *UpdateUserParams) (*UserDB, error) {
	q := `
 	UPDATE users
	SET display_name = COALESCE(?, display_name),
    updated_at = CURRENT_TIMESTAMP
	WHERE id = ?;
	`

	q = r.db.Rebind(q)

	_, err := r.db.NamedExecContext(ctx, q, params)
	if err != nil {
		return nil, err
	}

	user, err := r.FindOne(ctx, params.ID)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *Repository) SoftDelete(ctx context.Context, id string) error {
	q := `
	UPDATE users
	SET status = 'pending_delete',
	deleted_at = CURRENT_TIMESTAMP
	WHERE id = ?;
	`

	q = r.db.Rebind(q)

	_, err := r.db.ExecContext(ctx, q, id)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) Deactivate(ctx context.Context, id string) error {
	q := `
	UPDATE users
	SET account_status = 'deactivated'
	WHERE id = ?;
	`

	q = r.db.Rebind(q)

	_, err := r.db.ExecContext(ctx, q, id)
	if err != nil {
		return err
	}

	return nil
}
func (r *Repository) Reactivate(ctx context.Context, id string) error {
	q := `
	UPDATE Users
	SET account_status = 'active'
	WHERE id = ?;
	`

	q = r.db.Rebind(q)

	_, err := r.db.ExecContext(ctx, q, id)
	if err != nil {
		return err
	}

	return nil
}
