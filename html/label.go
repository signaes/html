package html

import (
	"fmt"

	"github.com/signaes/elements/elements"
)

func Label(attributes ...fmt.Stringer) func(children ...fmt.Stringer) elements.Element {
	return elements.New("label", attributes...)
}
