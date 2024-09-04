package html

import (
	"fmt"

	"github.com/signaes/elements/elements"
)

func Article(attributes ...fmt.Stringer) func(children ...fmt.Stringer) elements.Element {
	return elements.New("article", attributes...)
}
