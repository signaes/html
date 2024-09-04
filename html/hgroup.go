package html

import (
	"fmt"

	"github.com/signaes/elements/elements"
)

func Hgroup(attributes ...fmt.Stringer) func(children ...fmt.Stringer) elements.Element {
	return elements.New("hgroup", attributes...)
}
