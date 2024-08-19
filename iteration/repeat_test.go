package iteration

import (
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a")
	expected := "aaaaa"

	assertCorrectRepetitions(t, repeated, expected)	
}

func assertCorrectRepetitions(t testing.TB, repeated, expected string) {
	t.Helper()
	if repeated != expected {
		t.Errorf("got %q expected %q", repeated, expected)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}
