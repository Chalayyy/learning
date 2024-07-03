package outputters

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("x", 7)
	expected := "xxxxxxx"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("o", 5)
	}
}

func ExampleRepeat() {
	result := Repeat("A", 10)
	fmt.Println(result)
	// Output: AAAAAAAAAA
}