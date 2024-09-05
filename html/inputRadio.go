package html

import (
	"fmt"

	"github.com/signaes/elements/attributes"
	"github.com/signaes/elements/elements"
)

func InputRadio(a ...fmt.Stringer) func(children ...fmt.Stringer) elements.Element {
	attrs := append(a, attributes.Attributes{"type": "radio"})

	return elements.New("input", attrs...)
}
