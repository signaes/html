package html

import (
	"fmt"

	"github.com/signaes/elements/elements"
)

func Select(attributes ...fmt.Stringer) func(children ...fmt.Stringer) elements.Element {
	return elements.New("select", attributes...)
}
