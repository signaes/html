package html

import (
	"fmt"

	"github.com/signaes/elements/elements"
)

func Ruby(attributes ...fmt.Stringer) func(children ...fmt.Stringer) elements.Element {
	return elements.New("ruby", attributes...)
}
