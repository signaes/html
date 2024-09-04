package html

import (
	"fmt"

	"github.com/signaes/elements/elements"
)

func Html(attributes ...fmt.Stringer) func(children ...fmt.Stringer) elements.Element {
	return elements.New("html", attributes...)
}
