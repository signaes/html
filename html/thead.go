package html

import (
	"fmt"

	"github.com/signaes/elements/elements"
)

func Thead(attributes ...fmt.Stringer) func(children ...fmt.Stringer) elements.Element {
	return elements.New("thead", attributes...)
}
