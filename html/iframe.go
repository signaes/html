package html

import (
	"fmt"

	"github.com/signaes/elements/elements"
)

func Iframe(attributes ...fmt.Stringer) func(children ...fmt.Stringer) elements.Element {
	return elements.New("iframe", attributes...)
}
