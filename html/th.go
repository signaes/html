package html

import (
	"fmt"

	"github.com/signaes/elements/elements"
)

func Th(attributes ...fmt.Stringer) func(children ...fmt.Stringer) elements.Element {
	return elements.New("th", attributes...)
}
