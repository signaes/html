package html

import (
	"fmt"

	"github.com/signaes/elements/elements"
)

func Dialog(attributes ...fmt.Stringer) func(children ...fmt.Stringer) elements.Element {
	return elements.New("dialog", attributes...)
}
