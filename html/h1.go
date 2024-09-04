package html

import (
	"fmt"

	"github.com/signaes/elements/elements"
)

func H1(attributes ...fmt.Stringer) func(children ...fmt.Stringer) elements.Element {
	return elements.New("h1", attributes...)
}
