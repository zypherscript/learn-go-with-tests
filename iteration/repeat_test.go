package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("Repeat 'a' five times", func(t *testing.T) {
		got := Repeat("a", 5)
		want := "aaaaa"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Repeat 'a' five times with build-in method from strings", func(t *testing.T) {
		got := BuiltInRepeat("a", 5)
		want := "aaaaa"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func BenchmarkBuiltInRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BuiltInRepeat("a", 5)
	}
}

func ExampleRepeat() {
	repeated := Repeat("a", 5)
	fmt.Println(repeated)
	// Output: aaaaa
}

func ExampleBuiltInRepeat() {
	repeated := BuiltInRepeat("a", 5)
	fmt.Println(repeated)
	// Output: aaaaa
}
