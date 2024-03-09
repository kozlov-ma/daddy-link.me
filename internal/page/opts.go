package page

import (
	"fmt"

	"daddy-link.me/internal/page/attr"
	"daddy-link.me/internal/page/el"
)

type Option interface {
	apply(*Page) error
}

type withAttribute struct {
	attr attr.Attribute
}

func (w withAttribute) apply(p *Page) error {
	if _, ok := p.Attributes[w.attr.DType]; ok {
		return fmt.Errorf(
			"a duplicate attribute with DType %q was specified",
			w.attr.DType,
		)
	}

	p.Attributes[w.attr.DType] = w.attr
	return nil
}

// WithAttribute allows you to specify an attr.Attribute for a page.
// page.New errors when a duplicate attr.Attribute.DType
// is specified with a WithAttribute option.
func WithAttribute(attr attr.Attribute) Option {
	return withAttribute{
		attr: attr,
	}
}

type withElement struct {
	element el.Element
}

func (w withElement) apply(p *Page) error {
	p.Elements = append(p.Elements, w.element)
	return nil
}

// WithElement allows you to provide an el.Element for a page.
// WithElement option never errors the page.New constructor.
func WithElement(element el.Element) Option {
	return withElement{
		element: element,
	}
}
