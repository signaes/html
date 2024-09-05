package html

import (
	"fmt"

	"github.com/signaes/elements/attributes"
	"github.com/signaes/elements/elements"
)

func InputNumber(a ...fmt.Stringer) func(children ...fmt.Stringer) elements.Element {
	attrs := append(a, attributes.Attributes{"type": "number"})

	return elements.New("input", attrs...)
}
