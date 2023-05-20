package abstractFinal

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
	b.alpha.work = b.work
	b.common()
}
