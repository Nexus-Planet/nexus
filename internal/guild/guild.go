package guild

type Guild struct {
	ID string `json:"id"`
}

type GuildRoles struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Color string `json:"color"`
}

type GuildPermissions struct {
}

type GuildDB struct {
	ID string `db:"id"`
}

type GuildRolesDB struct {
	ID    string `db:"id"`
	Title string `db:"title"`
	Color string `db:"color"`
}

func (gr *GuildRolesDB) ToGuildRoles() *GuildRoles {
	return &GuildRoles{
		ID:    gr.ID,
		Title: gr.Title,
		Color: gr.Color,
	}
}
