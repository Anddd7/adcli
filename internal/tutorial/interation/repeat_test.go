package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 5)
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func ExampleRepeat() {
	repeated := Repeat("a", 5)
	fmt.Println(repeated)
	// Output: aaaaa
}

// another type of test, will record the performance informantion
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func BenchmarkRepeatFast(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RepeatFast("a", 5)
	}
}
