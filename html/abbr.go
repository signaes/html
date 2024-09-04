package html

import (
	"fmt"

	"github.com/signaes/elements/elements"
)

func Abbr(attributes ...fmt.Stringer) func(children ...fmt.Stringer) elements.Element {
	return elements.New("abbr", attributes...)
}
