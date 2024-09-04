package html

import (
	"fmt"

	"github.com/signaes/elements/elements"
)

func H3(attributes ...fmt.Stringer) func(children ...fmt.Stringer) elements.Element {
	return elements.New("h3", attributes...)
}
