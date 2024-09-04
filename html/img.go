package html

import (
	"fmt"

	"github.com/signaes/elements/elements"
)

func Img(attributes ...fmt.Stringer) func(children ...fmt.Stringer) elements.Element {
	return elements.New("img", attributes...)
}
