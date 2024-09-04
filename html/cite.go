package html

import (
	"fmt"

	"github.com/signaes/elements/elements"
)

func Cite(attributes ...fmt.Stringer) func(children ...fmt.Stringer) elements.Element {
	return elements.New("cite", attributes...)
}
