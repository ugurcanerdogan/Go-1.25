package tests

import "testing"

func TestAttrs(t *testing.T) {
	t.Attr("issue", "sellable contents")
	t.Attr("description", "stock test for sellable content")

	if 21*2 != 42 {
		t.Fatal("Are you sure?")
	}
}
