package html

import (
	"fmt"

	"github.com/signaes/elements/elements"
)

func S(attributes ...fmt.Stringer) func(children ...fmt.Stringer) elements.Element {
	return elements.New("s", attributes...)
}
