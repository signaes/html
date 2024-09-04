package html

import (
	"fmt"

	"github.com/signaes/elements/elements"
)

func Summary(attributes ...fmt.Stringer) func(children ...fmt.Stringer) elements.Element {
	return elements.New("summary", attributes...)
}
