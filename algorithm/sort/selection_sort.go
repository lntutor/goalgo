package sort

func SelectionSort(array []int) []int {
	for i := 0; i < len(array)-1; i++ {
		min := i
		for j := i + 1; j < len(array); j++ {
			if array[j] < array[min] {
				min = j
			}
		}

		array[min], array[i] = array[i], array[min]
	}
	return array
}
