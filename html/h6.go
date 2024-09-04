package html

import (
	"fmt"

	"github.com/signaes/elements/elements"
)

func H6(attributes ...fmt.Stringer) func(children ...fmt.Stringer) elements.Element {
	return elements.New("h6", attributes...)
}
