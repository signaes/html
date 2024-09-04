package html

import (
	"fmt"

	"github.com/signaes/elements/elements"
)

func Meta(attributes ...fmt.Stringer) func(children ...fmt.Stringer) elements.Element {
	return elements.New("meta", attributes...)
}
