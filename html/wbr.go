package html

import (
	"fmt"

	"github.com/signaes/elements/elements"
)

func Wbr(attributes ...fmt.Stringer) func(children ...fmt.Stringer) elements.Element {
	return elements.New("wbr", attributes...)
}
