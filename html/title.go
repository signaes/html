package html

import (
	"fmt"

	"github.com/signaes/elements/elements"
)

func Title(attributes ...fmt.Stringer) func(children ...fmt.Stringer) elements.Element {
	return elements.New("title", attributes...)
}
