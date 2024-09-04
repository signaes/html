package html

import (
	"fmt"

	"github.com/signaes/elements/elements"
)

func Base(attributes ...fmt.Stringer) func(children ...fmt.Stringer) elements.Element {
	return elements.New("base", attributes...)
}
