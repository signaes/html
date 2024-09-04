package html

import (
	"fmt"

	"github.com/signaes/elements/elements"
)

func Span(attributes ...fmt.Stringer) func(children ...fmt.Stringer) elements.Element {
	return elements.New("span", attributes...)
}
