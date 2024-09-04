package html

import (
	"fmt"

	"github.com/signaes/elements/elements"
)

func Figure(attributes ...fmt.Stringer) func(children ...fmt.Stringer) elements.Element {
	return elements.New("figure", attributes...)
}
