package auth

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

func (r *Repository) CreateSession(ctx context.Context, data *CreateSessionParams) (*AuthSessionDB, error) {
	q := `
	INSERT INTO auth_sessions(id, email, password_hash)
	VALUES(?, ?, ?);
	`

	q = r.db.Rebind(q)

	_, err := r.db.NamedExecContext(ctx, q, data)
	if err != nil {
		return nil, err
	}

	session, err := r.FindOne(ctx, data.ID)
	if err != nil {
		return nil, err
	}

	return session, nil
}

func (r *Repository) FindOne(ctx context.Context, id string) (*AuthSessionDB, error) {
	q := `
	SELECT *
	FROM auth_sessions
	WHERE id = ?;
	`

	q = r.db.Rebind(q)

	var session AuthSessionDB
	err := r.db.GetContext(ctx, &session, q)
	if err != nil {
		return nil, err
	}

	return &session, nil
}

func (r *Repository) FindOneByEmail(ctx context.Context, email string) (*AuthSessionDB, error) {
	q := `
	SELECT *
	FROM auth_sessions
	WHERE email = ?;
	`

	q = r.db.Rebind(q)

	var session AuthSessionDB
	err := r.db.GetContext(ctx, &session, q)
	if err != nil {
		return nil, err
	}

	return &session, nil
}

func (r *Repository) FindAllSessions(ctx context.Context) ([]*AuthSessionDB, error) {
	q := `
	SELECT *
	FROM auth_sessions
	ORDER BY created_at DESC;
	`

	q = r.db.Rebind(q)

	var sessions []*AuthSessionDB
	err := r.db.SelectContext(ctx, &sessions, q)
	if err != nil {
		return nil, err
	}

	return sessions, nil
}
