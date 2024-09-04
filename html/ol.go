package html

import (
	"fmt"

	"github.com/signaes/elements/elements"
)

func Ol(attributes ...fmt.Stringer) func(children ...fmt.Stringer) elements.Element {
	return elements.New("ol", attributes...)
}
