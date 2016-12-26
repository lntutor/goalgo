package sort

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	array := []int{10, 3, 2, 4, 1}
	array = InsertionSort(array)
	assert.Equal(t, []int{1, 2, 3, 4, 10}, array)
}

// func TestInsertionSort(t *testing.T) {
// 	array := []int{9, 10, 3, 4, 5}
// 	array = InsertionSort(array)
// 	assert.Equal(t, array, []int{3, 4, 5, 9, 10})
// }
