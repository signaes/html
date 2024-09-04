package html

import (
	"fmt"

	"github.com/signaes/elements/elements"
)

func Form(attributes ...fmt.Stringer) func(children ...fmt.Stringer) elements.Element {
	return elements.New("form", attributes...)
}
