package html

import (
	"fmt"

	"github.com/signaes/elements/elements"
)

func B(attributes ...fmt.Stringer) func(children ...fmt.Stringer) elements.Element {
	return elements.New("b", attributes...)
}
