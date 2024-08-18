package integers

import (
	"testing"
	"fmt"
)

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	assertCorrectAdd(t, sum, expected)
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
func assertCorrectAdd(t testing.TB, sum, expected int) {
	t.Helper()
	if sum != expected {
		t.Errorf("got %d wanted %d", sum, expected)
	}
}
