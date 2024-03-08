package page

import (
	"changemedaddy/internal/page/attr"
	"changemedaddy/internal/page/el"
)

type Option interface {
	apply(*Page)
}

type withAttribute struct {
	attr attr.Attribute
}

func (w withAttribute) apply(p *Page) {
	p.Attributes = append(p.Attributes, w.attr)
}

func WithAttribute(attr attr.Attribute) Option {
	return withAttribute{
		attr: attr,
	}
}

type withElement struct {
	element el.Element
}

func (w withElement) apply(p *Page) {
	p.Elements = append(p.Elements, w.element)
}

func WithElement(element el.Element) Option {
	return withElement{
		element: element,
	}
}
