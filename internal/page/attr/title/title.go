package title

import (
	"github.com/google/uuid"

	"daddy-link.me/internal/page/attr"
)

// write a Title attribute, similar to el.btn.Button

// Title represents a Title on the website.
type Title struct {
	Text string `json:"text" bson:"text"`
}

func (t Title) CSS() string {
	return "title"
}

func (t Title) DType() string {
	return "title"
}

func New(id uuid.UUID, text string) attr.Attribute {
	title := Title{text}
	return attr.NewAttribute(id, title)
}

// This line will prevent compilation if Title does not implement Content interface.
var _ attr.Data = (*Title)(nil)
