package html

import (
	"fmt"

	"github.com/signaes/elements/elements"
)

func Aside(attributes ...fmt.Stringer) func(children ...fmt.Stringer) elements.Element {
	return elements.New("aside", attributes...)
}
