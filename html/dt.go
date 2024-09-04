package html

import (
	"fmt"

	"github.com/signaes/elements/elements"
)

func Dt(attributes ...fmt.Stringer) func(children ...fmt.Stringer) elements.Element {
	return elements.New("dt", attributes...)
}
