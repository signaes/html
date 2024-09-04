package html

import (
	"fmt"

	"github.com/signaes/elements/elements"
)

func Dl(attributes ...fmt.Stringer) func(children ...fmt.Stringer) elements.Element {
	return elements.New("dl", attributes...)
}
