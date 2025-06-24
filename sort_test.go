package slicesortlib

import (
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	t.Run("Empty slice", func(t *testing.T) {
		input := []int{}
		expected := []int{}
		result := QuickSort(input)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("QuickSort(%v) = %v, want %v", input, result, expected)
		}
	})

	t.Run("Single element", func(t *testing.T) {
		input := []int{42}
		expected := []int{42}
		result := QuickSort(input)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("QuickSort(%v) = %v, want %v", input, result, expected)
		}
	})

	t.Run("Already sorted", func(t *testing.T) {
		input := []int{1, 2, 3, 4, 5}
		expected := []int{1, 2, 3, 4, 5}
		result := QuickSort(input)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("QuickSort(%v) = %v, want %v", input, result, expected)
		}
	})

	t.Run("Reverse", func(t *testing.T) {
		input := []int{5, 4, 3, 2, 1}
		expected := []int{1, 2, 3, 4, 5}
		result := QuickSort(input)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("QuickSort(%v) = %v, want %v", input, result, expected)
		}
	})

	t.Run("With duplicates", func(t *testing.T) {
		input := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3}
		expected := []int{1, 1, 2, 3, 3, 4, 5, 5, 6, 9}
		result := QuickSort(input)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("QuickSort(%v) = %v, want %v", input, result, expected)
		}
	})
}

func TestInsertionSort(t *testing.T) {
	t.Run("Empty slice", func(t *testing.T) {
		input := []int{}
		expected := []int{}
		result := InsertionSort(input)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("InsertionSort(%v) = %v, want %v", input, result, expected)
		}
	})

	t.Run("Single element", func(t *testing.T) {
		input := []int{7}
		expected := []int{7}
		result := InsertionSort(input)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("InsertionSort(%v) = %v, want %v", input, result, expected)
		}
	})

	t.Run("Random", func(t *testing.T) {
		input := []int{9, 7, 5, 11, 12, 2, 14, 3, 10, 6}
		expected := []int{2, 3, 5, 6, 7, 9, 10, 11, 12, 14}
		result := InsertionSort(input)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("InsertionSort(%v) = %v, want %v", input, result, expected)
		}
	})

	t.Run("All equal", func(t *testing.T) {
		input := []int{8, 8, 8, 8, 8}
		expected := []int{8, 8, 8, 8, 8}
		result := InsertionSort(input)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("InsertionSort(%v) = %v, want %v", input, result, expected)
		}
	})

	t.Run("Negative", func(t *testing.T) {
		input := []int{-5, -1, -8, -3, 0, 2, -4}
		expected := []int{-8, -5, -4, -3, -1, 0, 2}
		result := InsertionSort(input)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("InsertionSort(%v) = %v, want %v", input, result, expected)
		}
	})
}
