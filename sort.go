package slicesortlib

func QuickSort(s []int) []int {
	if len(s) < 2 {
		return s
	}
	sorted := make([]int, len(s))
	copy(sorted, s)
	quickSort(sorted, 0, len(sorted)-1)
	return sorted
}

func quickSort(s []int, l, r int) {
	if l < r {
		p := partition(s, l, r)
		quickSort(s, l, p-1)
		quickSort(s, p+1, r)
	}
}

func partition(s []int, l, r int) int {
	pivot := s[r]
	i := l
	for j := l; j < r; j++ {
		if s[j] < pivot {
			s[i], s[j] = s[j], s[i]
			i++
		}
	}
	s[i], s[r] = s[r], s[i]
	return i
}

func InsertionSort(s []int) []int {
	if len(s) < 2 {
		return s
	}
	sorted := make([]int, len(s))
	copy(sorted, s)
	for i := 1; i < len(sorted); i++ {
		key := sorted[i]
		j := i - 1
		for j >= 0 && sorted[j] > key {
			sorted[j+1] = sorted[j]
			j--
		}
		sorted[j+1] = key
	}
	return sorted
}
