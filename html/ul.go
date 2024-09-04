package html

import (
	"fmt"

	"github.com/signaes/elements/elements"
)

func Ul(attributes ...fmt.Stringer) func(children ...fmt.Stringer) elements.Element {
	return elements.New("ul", attributes...)
}
