package html

import (
	"fmt"

	"github.com/signaes/elements/elements"
)

func Li(attributes ...fmt.Stringer) func(children ...fmt.Stringer) elements.Element {
	return elements.New("li", attributes...)
}
