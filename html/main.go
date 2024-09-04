package html

import (
	"fmt"

	"github.com/signaes/elements/elements"
)

func Main(attributes ...fmt.Stringer) func(children ...fmt.Stringer) elements.Element {
	return elements.New("main", attributes...)
}
