package html

import (
	"fmt"

	"github.com/signaes/elements/elements"
)

func Table(attributes ...fmt.Stringer) func(children ...fmt.Stringer) elements.Element {
	return elements.New("table", attributes...)
}
