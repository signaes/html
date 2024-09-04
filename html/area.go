package html

import (
	"fmt"

	"github.com/signaes/elements/elements"
)

func Area(attributes ...fmt.Stringer) func(children ...fmt.Stringer) elements.Element {
	return elements.New("area", attributes...)
}
