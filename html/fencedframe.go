package html

import (
	"fmt"

	"github.com/signaes/elements/elements"
)

func Fencedframe(attributes ...fmt.Stringer) func(children ...fmt.Stringer) elements.Element {
	return elements.New("fencedframe", attributes...)
}
