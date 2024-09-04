package html

import (
	"fmt"

	"github.com/signaes/elements/elements"
)

func Hr(attributes ...fmt.Stringer) func(children ...fmt.Stringer) elements.Element {
	return elements.New("hr", attributes...)
}
