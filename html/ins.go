package html

import (
	"fmt"

	"github.com/signaes/elements/elements"
)

func Ins(attributes ...fmt.Stringer) func(children ...fmt.Stringer) elements.Element {
	return elements.New("ins", attributes...)
}
