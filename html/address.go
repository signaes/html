package html

import (
	"fmt"

	"github.com/signaes/elements/elements"
)

func Address(attributes ...fmt.Stringer) func(children ...fmt.Stringer) elements.Element {
	return elements.New("address", attributes...)
}
