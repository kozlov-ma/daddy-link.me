package page

import (
	"github.com/google/uuid"
	"github.com/pkg/errors"

	"daddy-link.me/internal/page/attr"
	"daddy-link.me/internal/page/el"
)

// A Page represents page content on the website.
type Page struct {
	ID         uuid.UUID                 `json:"id" bson:"_id"`
	Attributes map[string]attr.Attribute `json:"attributes" bson:"attributes"`
	Elements   []el.Element              `json:"elements" bson:"elements"`
}

func New(id uuid.UUID, opts ...Option) (*Page, error) {

	p := &Page{ID: id}

	for _, opt := range opts {
		if err := opt.apply(p); err != nil {
			return nil, errors.Wrapf(
				err,
				"failed to apply option %#v when building a Page (id %d)",
				opt,
				id,
			)
		}
	}

	return p, nil
}
