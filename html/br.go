package html

import (
	"fmt"

	"github.com/signaes/elements/elements"
)

func Br(attributes ...fmt.Stringer) func(children ...fmt.Stringer) elements.Element {
	return elements.New("br", attributes...)
}
