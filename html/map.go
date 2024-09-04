package html

import (
	"fmt"

	"github.com/signaes/elements/elements"
)

func Map(attributes ...fmt.Stringer) func(children ...fmt.Stringer) elements.Element {
	return elements.New("map", attributes...)
}
