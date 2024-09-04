package html

import (
	"fmt"

	"github.com/signaes/elements/elements"
)

func Header(attributes ...fmt.Stringer) func(children ...fmt.Stringer) elements.Element {
	return elements.New("header", attributes...)
}
