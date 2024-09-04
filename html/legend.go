package html

import (
	"fmt"

	"github.com/signaes/elements/elements"
)

func Legend(attributes ...fmt.Stringer) func(children ...fmt.Stringer) elements.Element {
	return elements.New("legend", attributes...)
}
