package html

import (
	"fmt"

	"github.com/signaes/elements/elements"
)

func Details(attributes ...fmt.Stringer) func(children ...fmt.Stringer) elements.Element {
	return elements.New("details", attributes...)
}
