package html

import (
	"fmt"

	"github.com/signaes/elements/elements"
)

func Nav(attributes ...fmt.Stringer) func(children ...fmt.Stringer) elements.Element {
	return elements.New("nav", attributes...)
}
