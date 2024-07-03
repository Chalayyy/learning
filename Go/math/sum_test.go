package math

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	values := []int{0,1,2,3,4,5,6,7,8,9,10}
	got := Sum(values)
	expected := 55

	if got != expected {
		t.Errorf("expected %d, got %d given %v", expected, got, values)
	}
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1,10,100}, []int{1,2,3,4})
	expected := []int{111,10}

	if !reflect.DeepEqual(got,expected) {
		t.Errorf("expected %v, got %v", expected, got)
	}
}

func TestSumAllTails(t *testing.T) {

	checkSums := func(t testing.TB, got, expected []int) {
		t.Helper()
		if !reflect.DeepEqual(got,expected) {
			t.Errorf("expected %v, got %v", expected, got)
		}
	}
	t.Run("Sum some slices", func(t *testing.T){
		got := SumAllTails([]int{1,10,100}, []int{1,2,3,4})
		expected := []int{110, 9}

		checkSums(t, got, expected)

	})

	t.Run("Safely sum some empty slices", func(t *testing.T){
		got := SumAllTails([]int{}, []int{1,2,3,4})
		expected := []int{0, 9}

		checkSums(t, got, expected)
	})

}