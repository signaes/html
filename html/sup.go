package html

import (
	"fmt"

	"github.com/signaes/elements/elements"
)

func Sup(attributes ...fmt.Stringer) func(children ...fmt.Stringer) elements.Element {
	return elements.New("sup", attributes...)
}
