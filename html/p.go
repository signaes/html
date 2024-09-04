package html

import (
	"fmt"

	"github.com/signaes/elements/elements"
)

func P(attributes ...fmt.Stringer) func(children ...fmt.Stringer) elements.Element {
	return elements.New("p", attributes...)
}
