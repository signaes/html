package html

import (
	"fmt"

	"github.com/signaes/elements/elements"
)

func Em(attributes ...fmt.Stringer) func(children ...fmt.Stringer) elements.Element {
	return elements.New("em", attributes...)
}
