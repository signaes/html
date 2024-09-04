package html

import (
	"fmt"

	"github.com/signaes/elements/elements"
)

func I(attributes ...fmt.Stringer) func(children ...fmt.Stringer) elements.Element {
	return elements.New("i", attributes...)
}
