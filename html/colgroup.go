package html

import (
	"fmt"

	"github.com/signaes/elements/elements"
)

func Colgroup(attributes ...fmt.Stringer) func(children ...fmt.Stringer) elements.Element {
	return elements.New("colgroup", attributes...)
}
