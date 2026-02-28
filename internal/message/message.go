package message

import "github.com/nexus-planet/nexus/internal/attachment"

type Message struct {
	ID          string                  `json:"id"`
	Content     string                  `json:"content"`
	Type        string                  `json:"type"`
	Attachments []attachment.Attachment `json:"attachments"`
}

type UpdateMessage struct {
	ID      string `json:"id"`
	Content string `json:"content"`
}

type MessageDB struct {
	ID          string                  `db:"id"`
	Content     string                  `db:"content"`
	Type        string                  `db:"type"`
	Attachments []attachment.Attachment `db:"attachments"`
}

type CreateMessageParams struct {
	ID          string                  `db:"id"`
	Content     string                  `db:"content"`
	Type        string                  `db:"type"`
	Attachments []attachment.Attachment `db:"attachments"`
}

type UpdateMessageParams struct {
	ID      string `db:"id"`
	Content string `db:"content"`
}

type TogglePinParams struct {
	MessageID string `db:"message_id"`
	GuildID   string `db:"guild_id"`
	IsPinned  int    `db:"is_pinned"`
}

func (m *MessageDB) ToMessage() *Message {
	return &Message{
		ID:          m.ID,
		Content:     m.Content,
		Type:        m.Type,
		Attachments: m.Attachments,
	}
}
