package Day1

import (
	"testing"
)

func Test_ReviewSequence(t *testing.T) {
	s := ReviewSequence("hello")
	if s != "hello" {
		t.Error("output not as expected")
	}
}