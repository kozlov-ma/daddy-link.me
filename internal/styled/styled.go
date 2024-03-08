package styled

// Styled represents a styled component of a page, like an Attribute(Data) or an Element(Content).
type Styled interface {
	// CSS class string for a Styled.
	// Must return a constant string.
	CSS() string
}
