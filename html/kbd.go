package html

import (
	"fmt"

	"github.com/signaes/elements/elements"
)

func Kbd(attributes ...fmt.Stringer) func(children ...fmt.Stringer) elements.Element {
	return elements.New("kbd", attributes...)
}
