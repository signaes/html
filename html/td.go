package html

import (
	"fmt"

	"github.com/signaes/elements/elements"
)

func Td(attributes ...fmt.Stringer) func(children ...fmt.Stringer) elements.Element {
	return elements.New("td", attributes...)
}
