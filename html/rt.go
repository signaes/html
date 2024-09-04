package html

import (
	"fmt"

	"github.com/signaes/elements/elements"
)

func Rt(attributes ...fmt.Stringer) func(children ...fmt.Stringer) elements.Element {
	return elements.New("rt", attributes...)
}
