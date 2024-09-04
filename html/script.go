package html

import (
	"fmt"

	"github.com/signaes/elements/elements"
)

func Script(attributes ...fmt.Stringer) func(children ...fmt.Stringer) elements.Element {
	return elements.New("script", attributes...)
}
