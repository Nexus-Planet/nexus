package attachment

type Attachment struct {
	ID        string `json:"id"`
	MessageID string `json:"message_id"`
	URL       string `json:"uri"`
	Type      string `json:"type"`
	Size      int64  `json:"Size"`
}

type AttachmentDB struct {
	ID        string `db:"id"`
	MessageID string `db:"message_id"`
	URL       string `db:"uri"`
	Type      string `db:"type"`
	Size      int64  `db:"Size"`
}
