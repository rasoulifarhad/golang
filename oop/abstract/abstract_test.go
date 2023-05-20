package abstract

import (
	"testing"
)

func Test_abstract(t *testing.T) {
	a := alpha{
		name: "test",
	}
	b := &beta{
		alpha: a,
	}
	b.work()
}
