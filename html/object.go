package html

import (
	"fmt"

	"github.com/signaes/elements/elements"
)

func Object(attributes ...fmt.Stringer) func(children ...fmt.Stringer) elements.Element {
	return elements.New("object", attributes...)
}
