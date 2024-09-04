package html

import (
	"fmt"

	"github.com/signaes/elements/elements"
)

func Pre(attributes ...fmt.Stringer) func(children ...fmt.Stringer) elements.Element {
	return elements.New("pre", attributes...)
}
