package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	want := 4
	got := Add(2, 2)
	if got != want {
		t.Errorf("want: %d, got: %d", want, got)
	}

	want = 40
	got = Add(22, 18)
	if got != want {
		t.Errorf("want: %d, got: %d", want, got)
	}
}

func ExampleAdd() {
	sum := Add(1, 4)
	fmt.Println(sum)
	// Output: 5
}
