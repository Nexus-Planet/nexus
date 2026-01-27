package auth

import (
	"github.com/nexus-planet/nexus-planet-api/internal/db"
)

type Repository struct {
	q *db.Queries
}

func NewRepository(q *db.Queries) *Repository {
	return &Repository{q: q}
}
