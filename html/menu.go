package html

import (
	"fmt"

	"github.com/signaes/elements/elements"
)

func Menu(attributes ...fmt.Stringer) func(children ...fmt.Stringer) elements.Element {
	return elements.New("menu", attributes...)
}
