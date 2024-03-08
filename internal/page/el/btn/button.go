package btn

import (
	"github.com/google/uuid"

	"changemedaddy/internal/page/el"
)

// A Button represents a Button on the website. It displays a Title.
// On click, user should be redirected to a Link
type Button struct {
	Title string `json:"title" bson:"title"`
	Link  string `json:"link" bson:"link"`
}

func (b Button) CType() string {
	return "Button"
}

func (b Button) CSS() string { // here we use a value receiver to state that Button is immutable
	return "Button"
}

// New creates a new Button as an Element, filling CType and CSS fields.
// USE ONLY VALIDATED DATA IN THIS CONSTRUCTOR.
func New(id uuid.UUID, rendOrder int, title, link string) el.Element {
	return el.NewElement(
		id, rendOrder, Button{
			Title: title,
			Link:  link,
		},
	)
}

// This line will prevent compilation if Button does not implement Content interface.
var _ el.Content = (*Button)(nil)
