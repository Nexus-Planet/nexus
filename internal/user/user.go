package user

import (
	"database/sql"
	"image"
	"time"

	"github.com/nexus-planet/nexus/internal/guild"
)

type User struct {
	ID          string             `json:"id"`
	DisplayName string             `json:"display_name"`
	Username    string             `json:"username"`
	Images      []image.Image      `json:"images"`
	Roles       []guild.GuildRoles `json:"roles"`
	CreatedAt   string             `json:"created_at"`
	UpdatedAt   string             `json:"updated_at"`
}

type UserCredentials struct {
	ID string `json:"id"`
}

type SetUsernameCredentials struct {
	ID                string `json:"id"`
	Username          string `json:"username"`
	UsernameChangedAt string `json:"username_changed_at"`
}

type UpdateUserCredentials struct {
	ID          *string `json:"id"`
	DisplayName *string `json:"display_name,omitempty"`
}

type UserDB struct {
	ID          string               `db:"id"`
	DisplayName sql.NullString       `db:"display_name"`
	Username    sql.NullString       `db:"username"`
	Images      []image.Image        `db:"images"`
	Roles       []guild.GuildRolesDB `db:"roles"`
	CreatedAt   time.Time            `db:"created_at"`
	UpdatedAt   time.Time            `db:"updated_at"`
}

type SetUsernameParams struct {
	ID                string    `db:"id"`
	Username          string    `db:"username"`
	UsernameChangedAt time.Time `db:"username_changed_at"`
}

type UpdateUserParams struct {
	ID          string         `db:"id"`
	DisplayName sql.NullString `db:"display_name"`
}

func (u *UserDB) ToUser() *User {
	roles := make([]guild.GuildRoles, len(u.Roles))
	for i, role := range u.Roles {
		r := role.ToGuildRoles()
		if r != nil {
			roles[i] = *r
		}
	}

	return &User{
		ID:          u.ID,
		DisplayName: u.DisplayName.String,
		Username:    u.Username.String,
		Images:      u.Images,
		Roles:       roles,
		CreatedAt:   u.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   u.UpdatedAt.Format(time.RFC3339),
	}
}
