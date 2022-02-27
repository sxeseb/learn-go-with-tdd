package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	assertResult := func(t testing.TB, expected, result string) {
		t.Helper()
		if result != expected {
			t.Errorf("expected %q, got %q", expected, result)
		}
	}

	t.Run("5 iteration test", func(t *testing.T) {
		repeated := Repeat("a", 5)
		expected := "aaaaa"

		assertResult(t, expected, repeated)
	})
}

func ExampleRepeat() {
	r := Repeat("B", 3)
	fmt.Println(r)
	// Output: BBB
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}
