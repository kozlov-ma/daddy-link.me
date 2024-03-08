package page

import (
	"github.com/google/uuid"

	"daddy-link.me/internal/page/attr"
	"daddy-link.me/internal/page/el"
)

// A Page represents page content on the website.
type Page struct {
	ID         uuid.UUID        `json:"id" bson:"_id"`
	Attributes []attr.Attribute `json:"attributes" bson:"attributes"`
	Elements   []el.Element     `json:"elements" bson:"elements"`
}

func New(id uuid.UUID, opts ...Option) Page {

	p := Page{ID: id}

	for _, opt := range opts {
		opt.apply(&p)
	}

	return p
}
