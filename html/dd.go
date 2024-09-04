package html

import (
	"fmt"

	"github.com/signaes/elements/elements"
)

func Dd(attributes ...fmt.Stringer) func(children ...fmt.Stringer) elements.Element {
	return elements.New("dd", attributes...)
}
