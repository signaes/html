package html

import (
	"fmt"

	"github.com/signaes/elements/elements"
)

func Audio(attributes ...fmt.Stringer) func(children ...fmt.Stringer) elements.Element {
	return elements.New("audio", attributes...)
}
