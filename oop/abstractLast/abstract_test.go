package abstractLast

import (
	"testing"
)

func TestAbstract(t *testing.T) {
	a := alpha{
		name: "test",
	}
	b := &beta{
		alpha: a,
	}
	b.common(b)
}
