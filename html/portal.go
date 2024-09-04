package html

import (
	"fmt"

	"github.com/signaes/elements/elements"
)

func Portal(attributes ...fmt.Stringer) func(children ...fmt.Stringer) elements.Element {
	return elements.New("portal", attributes...)
}
