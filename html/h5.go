package html

import (
	"fmt"

	"github.com/signaes/elements/elements"
)

func H5(attributes ...fmt.Stringer) func(children ...fmt.Stringer) elements.Element {
	return elements.New("h5", attributes...)
}
