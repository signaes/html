package html

import (
	"fmt"

	"github.com/signaes/elements/elements"
)

func Data(attributes ...fmt.Stringer) func(children ...fmt.Stringer) elements.Element {
	return elements.New("data", attributes...)
}
