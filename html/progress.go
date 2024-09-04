package html

import (
	"fmt"

	"github.com/signaes/elements/elements"
)

func Progress(attributes ...fmt.Stringer) func(children ...fmt.Stringer) elements.Element {
	return elements.New("progress", attributes...)
}
