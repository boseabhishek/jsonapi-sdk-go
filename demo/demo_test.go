package demo

import "testing"

func TestEqual(t *testing.T) {
	p := NewPoint(12, 12, 12)

	want := true
	got := p.Equal()

	if want != got {
		t.Errorf("test failed")
	}

}