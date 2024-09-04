package html

import (
	"fmt"

	"github.com/signaes/elements/elements"
)

func Fieldset(attributes ...fmt.Stringer) func(children ...fmt.Stringer) elements.Element {
	return elements.New("fieldset", attributes...)
}
