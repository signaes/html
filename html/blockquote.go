package html

import (
	"fmt"

	"github.com/signaes/elements/elements"
)

func Blockquote(attributes ...fmt.Stringer) func(children ...fmt.Stringer) elements.Element {
	return elements.New("blockquote", attributes...)
}
