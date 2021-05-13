package main

import "testing"

func TestInsertionSort(t *testing.T) {
	t.Run("Testing insertion sort", func(t *testing.T) {
		values := []int {3, 5, 7, 1, 2}
		ans := insertionSort(values)
		expected := []int {1, 2, 3, 5, 7}
		for i := 0; i < len(ans); i++ {
			if ans[i] != expected[i] {
				t.Errorf("got %d, wanted %d", ans[i], expected[i])
				return
			}
		}
	})

	t.Run("Testing reverse insertion sort", func(t *testing.T) {
		values := []int {3, 5, 7, 1, 2}
		ans := reverseInsertionSort(values)
		expected := []int {7, 5, 3, 2, 1}
		for i := 0; i < len(ans); i++ {
			if ans[i] != expected[i] {
				t.Errorf("got %d, wanted %d", ans[i], expected[i])
				return
			}
		}
	})

}
