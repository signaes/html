package html

import (
	"fmt"

	"github.com/signaes/elements/elements"
)

func Time(attributes ...fmt.Stringer) func(children ...fmt.Stringer) elements.Element {
	return elements.New("time", attributes...)
}
