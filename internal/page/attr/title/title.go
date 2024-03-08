package title

import (
	"github.com/google/uuid"

	"changemedaddy/internal/page/attr"
)

// write a Title attribute, similar to el.btn.Button

// Title represents a Title on the website.
type Title struct {
	Text string `json:"text" bson:"text"`
}

func (t Title) CSS() string {
	return "Title"
}

func (t Title) DType() string {
	return "Title"
}

func New(id uuid.UUID, text string) attr.Attribute {
	title := Title{text}
	return attr.NewAttribute(id, title)
}

// This line will prevent compilation if Title does not implement Content interface.
var _ attr.Data = (*Title)(nil)
