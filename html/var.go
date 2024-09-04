package html

import (
	"fmt"

	"github.com/signaes/elements/elements"
)

func Var(attributes ...fmt.Stringer) func(children ...fmt.Stringer) elements.Element {
	return elements.New("var", attributes...)
}
