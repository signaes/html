package html

import (
	"fmt"

	"github.com/signaes/elements/elements"
)

func Search(attributes ...fmt.Stringer) func(children ...fmt.Stringer) elements.Element {
	return elements.New("search", attributes...)
}
