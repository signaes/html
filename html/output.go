package html

import (
	"fmt"

	"github.com/signaes/elements/elements"
)

func Output(attributes ...fmt.Stringer) func(children ...fmt.Stringer) elements.Element {
	return elements.New("output", attributes...)
}
