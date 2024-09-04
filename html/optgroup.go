package html

import (
	"fmt"

	"github.com/signaes/elements/elements"
)

func Optgroup(attributes ...fmt.Stringer) func(children ...fmt.Stringer) elements.Element {
	return elements.New("optgroup", attributes...)
}
