package html

import (
	"fmt"

	"github.com/signaes/elements/elements"
)

func Col(attributes ...fmt.Stringer) func(children ...fmt.Stringer) elements.Element {
	return elements.New("col", attributes...)
}
