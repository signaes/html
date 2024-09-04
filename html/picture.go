package html

import (
	"fmt"

	"github.com/signaes/elements/elements"
)

func Picture(attributes ...fmt.Stringer) func(children ...fmt.Stringer) elements.Element {
	return elements.New("picture", attributes...)
}
