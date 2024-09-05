package html

import (
	"fmt"

	"github.com/signaes/elements/attributes"
	"github.com/signaes/elements/elements"
)

func InputDate(a ...fmt.Stringer) func(children ...fmt.Stringer) elements.Element {
	attrs := append(a, attributes.Attributes{"type": "date"})

	return elements.New("input", attrs...)
}
