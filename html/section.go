package html

import (
	"fmt"

	"github.com/signaes/elements/elements"
)

func Section(attributes ...fmt.Stringer) func(children ...fmt.Stringer) elements.Element {
	return elements.New("section", attributes...)
}
