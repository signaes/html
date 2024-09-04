package html

import (
	"fmt"

	"github.com/signaes/elements/elements"
)

func Tbody(attributes ...fmt.Stringer) func(children ...fmt.Stringer) elements.Element {
	return elements.New("tbody", attributes...)
}
