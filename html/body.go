package html

import (
	"fmt"

	"github.com/signaes/elements/elements"
)

func Body(attributes ...fmt.Stringer) func(children ...fmt.Stringer) elements.Element {
	return elements.New("body", attributes...)
}
