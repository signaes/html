package html

import (
	"fmt"

	"github.com/signaes/elements/elements"
)

func U(attributes ...fmt.Stringer) func(children ...fmt.Stringer) elements.Element {
	return elements.New("u", attributes...)
}
