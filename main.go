package main

func insertionSort(a []int) []int {
	for j := 1; j < len(a); j++ {
		for i := j - 1; i > -1 && a[i] > a[i+1]; i-- {
			a[i+1], a[i] = a[i], a[i+1]
		}
	}
	return a
}

func reverseInsertionSort(a []int) []int {
	for j := 1; j < len(a); j++ {
		for i := j - 1; i > -1 && a[i] < a[i+1]; i-- {
			a[i+1], a[i] = a[i], a[i+1]
		}
	}
	return a
}


func main() {}