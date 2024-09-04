package html

import (
	"fmt"

	"github.com/signaes/elements/elements"
)

func Footer(attributes ...fmt.Stringer) func(children ...fmt.Stringer) elements.Element {
	return elements.New("footer", attributes...)
}
