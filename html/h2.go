package html

import (
	"fmt"

	"github.com/signaes/elements/elements"
)

func H2(attributes ...fmt.Stringer) func(children ...fmt.Stringer) elements.Element {
	return elements.New("h2", attributes...)
}
