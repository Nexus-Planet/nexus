package user

import "github.com/nexus-planet/nexus-planet-api/internal/media/image"

type User struct {
	ID          string        `json:"id"`
	DisplayName string        `json:"display_name"`
	Username    string        `json:"username"`
	Images      []image.Image `json:"images"`
	Roles       []string      `json:"roles"`
	Permissions []string      `json:"permissions"`
	CreatedAt   string        `json:"created_at"`
	UpdatedAt   string        `json:"updated_at"`
}

type UserResponse struct {
	ID          string        `json:"id"`
	Name        string        `json:"name"`
	Username    string        `json:"username"`
	Images      []image.Image `json:"images"`
	Roles       []string      `json:"roles"`
	Permissions []string      `json:"permissions"`
	CreatedAt   string        `json:"created_at"`
	UpdatedAt   string        `json:"updated_at"`
}

type UserCredentials struct {
	UserID string `json:"user_id"`
}
