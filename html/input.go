package html

import (
	"fmt"

	"github.com/signaes/elements/elements"
)

func Input(attributes ...fmt.Stringer) func(children ...fmt.Stringer) elements.Element {
	return elements.New("input", attributes...)
}
