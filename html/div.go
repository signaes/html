package html

import (
	"fmt"

	"github.com/signaes/elements/elements"
)

func Div(attributes ...fmt.Stringer) func(children ...fmt.Stringer) elements.Element {
	return elements.New("div", attributes...)
}
