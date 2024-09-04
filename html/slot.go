package html

import (
	"fmt"

	"github.com/signaes/elements/elements"
)

func Slot(attributes ...fmt.Stringer) func(children ...fmt.Stringer) elements.Element {
	return elements.New("slot", attributes...)
}
