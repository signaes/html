package html

import (
	"fmt"

	"github.com/signaes/elements/elements"
)

func Strong(attributes ...fmt.Stringer) func(children ...fmt.Stringer) elements.Element {
	return elements.New("strong", attributes...)
}
