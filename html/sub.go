package html

import (
	"fmt"

	"github.com/signaes/elements/elements"
)

func Sub(attributes ...fmt.Stringer) func(children ...fmt.Stringer) elements.Element {
	return elements.New("sub", attributes...)
}
