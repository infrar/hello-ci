package foo

import "testing"

func TestFoo(t *testing.T) {
	got := 1
	if got != 1 {
		t.Errorf("ans = %d; want 1", got)
	}
}
