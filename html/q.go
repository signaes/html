package html

import (
	"fmt"

	"github.com/signaes/elements/elements"
)

func Q(attributes ...fmt.Stringer) func(children ...fmt.Stringer) elements.Element {
	return elements.New("q", attributes...)
}
