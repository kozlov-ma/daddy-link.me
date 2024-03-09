package attr

import (
	"github.com/google/uuid"

	"daddy-link.me/internal/styled"
)

// An Attribute is a container for a page attribute, i.e. Title, Avatar, etc.
// An Attribute of one given type may be present at most once in a page, i.e.
// a Page cannot have two Avatars.
type Attribute struct {
	// A unique identifier for an attribute
	ID uuid.UUID `json:"id" bson:"_id"`

	// CSS is a CSS classes string for the attribute. Must be taken from D.CSS().
	CSS string `json:"css" bson:"css"`

	// DType represents the data type of Attribute, i.e. "avatar". Will be taken from D.DType() by a constructor.
	DType string `json:"dtype" bson:"dtype"`

	// Data of the attribute.
	D Data `json:"data" bson:"data"`
}

// A Data represents the data of an attribute, like a text for a Title or source for the Avatar.
// An implementation of Data MUST contain all the fields necessary
// to be exchanged between frontend and backend, and to store it in the
// database. It must also contain struct tags for json and bson accordingly.
type Data interface {
	styled.Styled

	// DType returns a data type to be serialized in json, i.e. "avatar".
	// MUST return a constant string.
	DType() string
}

// NewAttribute creates a new Attribute, filling DType and CSS fields.
// USE ONLY VALIDATED DATA IN THIS CONSTRUCTOR.
func NewAttribute(id uuid.UUID, d Data) Attribute {
	return Attribute{
		ID:    id,
		DType: d.DType(),
		D:     d,
		CSS:   d.CSS(),
	}
}
