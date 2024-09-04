package html

import (
	"fmt"

	"github.com/signaes/elements/elements"
)

func Code(attributes ...fmt.Stringer) func(children ...fmt.Stringer) elements.Element {
	return elements.New("code", attributes...)
}
