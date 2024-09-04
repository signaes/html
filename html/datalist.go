package html

import (
	"fmt"

	"github.com/signaes/elements/elements"
)

func Datalist(attributes ...fmt.Stringer) func(children ...fmt.Stringer) elements.Element {
	return elements.New("datalist", attributes...)
}
