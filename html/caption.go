package html

import (
	"fmt"

	"github.com/signaes/elements/elements"
)

func Caption(attributes ...fmt.Stringer) func(children ...fmt.Stringer) elements.Element {
	return elements.New("caption", attributes...)
}
