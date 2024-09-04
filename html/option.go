package html

import (
	"fmt"

	"github.com/signaes/elements/elements"
)

func Option(attributes ...fmt.Stringer) func(children ...fmt.Stringer) elements.Element {
	return elements.New("option", attributes...)
}
