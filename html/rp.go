package html

import (
	"fmt"

	"github.com/signaes/elements/elements"
)

func Rp(attributes ...fmt.Stringer) func(children ...fmt.Stringer) elements.Element {
	return elements.New("rp", attributes...)
}
