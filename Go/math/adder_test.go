package math

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	sum := Add(420, 69)
	expected := 489

	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}

func ExampleAdd() {
	sum := Add(99, 2)
	fmt.Println(sum)
	// Output: 101
}