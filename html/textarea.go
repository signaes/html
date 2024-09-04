package html

import (
	"fmt"

	"github.com/signaes/elements/elements"
)

func Textarea(attributes ...fmt.Stringer) func(children ...fmt.Stringer) elements.Element {
	return elements.New("textarea", attributes...)
}
