package main

import (
	"reflect"
	"testing"
)

// Test UpperCase
func TestToUpperCase(t *testing.T){

	t.Run("correct index", func(t *testing.T) {
		got, _:= ToUpperCase("javascript", 3)
		want:= "javAscript"
		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})

	t.Run("invalid index", func(t *testing.T) {
		_, err:= ToUpperCase("go", 4)
		want:= "Invalid index"
		if err == nil {
			t.Fatal("didn't get an error but wanted one")
		}
		if err.Error() != want {
			t.Errorf("got %s want %s", err.Error() , want)
		}

	})
}
// Test Range
func TestFilterRange(t *testing.T){
	helper:= func(t *testing.T, arr, candidate  []int, min, max int ) {
		got:= FilterRange(arr, min,max)
		want:= candidate
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	}

	t.Run("correct range", func(t *testing.T) {
		array:= []int{1,2,3,4,5,6,7,8,9}
		candidate:= []int{3,4,5,6,7}
		helper(t,array, candidate, 3, 7 )
	})
	t.Run("with negative argument", func(t *testing.T) {
		array:= []int{-5,-4,-3,-2, 10,11,12}
		candidate := []int{-4 ,-3, -2, 10}
		helper(t, array, candidate  , -4, 10 )
	})
}

func TestRemoveItemInArray(t *testing.T) {
	t.Run("", func(t *testing.T) {
		sliceOfStrings:= []string{"javaScript", "Go", "React"}
		got:= RemoveItemInArray(sliceOfStrings, "a")
		want:= []string{"Go"}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}