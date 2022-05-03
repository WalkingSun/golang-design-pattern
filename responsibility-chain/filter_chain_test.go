package responsibility_chain

import (
	"fmt"
	"testing"
)

func TestFilter(t *testing.T) {
	f := NewFilter(true, true, true)
	res := f.Filter(100000)
	if len(res) == 0 {
		t.Fatal("verify fail")
	}
	fmt.Println(res)
}
