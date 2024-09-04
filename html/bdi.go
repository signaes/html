package html

import (
	"fmt"

	"github.com/signaes/elements/elements"
)

func Bdi(attributes ...fmt.Stringer) func(children ...fmt.Stringer) elements.Element {
	return elements.New("bdi", attributes...)
}
