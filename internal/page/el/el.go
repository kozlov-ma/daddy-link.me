package el

import (
	"github.com/google/uuid"

	"daddy-link.me/internal/styled"
)

// An Element is a container for a page element, i.e. Button, Image, etc.
type Element struct {
	// A unique identifier for an element
	ID uuid.UUID `json:"id" bson:"_id"`

	// A rendering order of an element. Lower numbers are rendered first.
	RendOrder int `json:"rendOrder" bson:"rend_order"`

	// CSS classes string for the element. Must be taken from C.CSS().
	CSS string `json:"css" bson:"css"`

	// Type of the element's content, i.e. "button", "image", etc. Will be taken from C.CType() by a constructor.
	CType string `json:"ctype" bson:"ctype"`

	// Content of the element.
	C Content `json:"content" bson:"content"`
}

// A Content represents the content of an element, like a link and a title for a Button.
// An implementation of Content MUST contain all the fields necessary
// to be exchanged between frontend and backend, and to store it in the
// database. It must also contain struct tags for json and bson accordingly.
type Content interface {
	styled.Styled

	// CType returns a string type the element's content, i.e. "button", "image", etc.
	// to be used in json and bson. Must return a constant string.
	CType() string
}

// NewElement creates a new Element, filling CType and CSS fields.
// USE ONLY VALIDATED DATA IN THIS CONSTRUCTOR.
func NewElement(id uuid.UUID, rendOrder int, c Content) Element {
	return Element{
		ID:        id,
		RendOrder: rendOrder,
		CType:     c.CType(),
		C:         c,
		CSS:       c.CSS(),
	}
}
