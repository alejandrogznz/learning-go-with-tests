package array_slice

import (
	"testing"
	"reflect"
	)

func TestSum(t *testing.T) {
	t.Run("Collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		assertCorrectSum(t, got, want, numbers)
	})
	
}

func TestSumAll(t *testing.T){
	t.Run("Sum an array of slices", func(t *testing.T){
		got := SumAll([]int{1, 2}, []int{0, 9})
		want := []int{3, 9}
		assertCorrectSumAll(t, got, want)
	})
}

func TestSumAllTails(t *testing.T){
	assertCorrectSumAllTail := func(t testing.TB, got, want []int){	
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v want: %v", got, want)
		}
	}
	t.Run("Sum Tails of 2 slices", func(t *testing.T){
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}
		assertCorrectSumAllTail(t, got, want)
	})
	t.Run("Sum Tails of an empty slice", func(t *testing.T){
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}
		assertCorrectSumAllTail(t, got, want)
	})
}

func assertCorrectSum(t testing.TB, got, want int, numbers []int) {
	t.Helper()
	if got != want {
		t.Errorf("got: %d; wanted: %d given %v", got, want, numbers)
	}
}

func assertCorrectSumAll(t testing.TB, got, want []int) {
	t.Helper()
	if !reflect.DeepEqual(got, want){
		t.Errorf("got: %v, want: %v", got, want)
	}
}
