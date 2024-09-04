package html

import (
	"fmt"

	"github.com/signaes/elements/elements"
)

func Tr(attributes ...fmt.Stringer) func(children ...fmt.Stringer) elements.Element {
	return elements.New("tr", attributes...)
}
