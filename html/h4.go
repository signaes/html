package html

import (
	"fmt"

	"github.com/signaes/elements/elements"
)

func H4(attributes ...fmt.Stringer) func(children ...fmt.Stringer) elements.Element {
	return elements.New("h4", attributes...)
}
