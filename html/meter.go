package html

import (
	"fmt"

	"github.com/signaes/elements/elements"
)

func Meter(attributes ...fmt.Stringer) func(children ...fmt.Stringer) elements.Element {
	return elements.New("meter", attributes...)
}
