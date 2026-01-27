package user

import "github.com/nexus-planet/nexus-planet-api/internal/media/image"

type User struct {
	ID           string        `json:"id"`
	Name         string        `json:"name"`
	Username     string        `json:"username"`
	Email        string        `json:"email"`
	PasswordHash string        `json:"password_hash"`
	Token        string        `json:"token"`
	Images       []image.Image `json:"images"`
	Roles        []string      `json:"roles"`
	Permissions  []string      `json:"permissions"`
	CreatedAt    string        `json:"created_at"`
	UpdatedAt    string        `json:"updated_at"`
}
