package sort

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNumberToTelco(t *testing.T) {
	array := []int{9, 10, 3, 4, 5}
	array = SelectionSort(array)
	assert.Equal(t, array, []int{3, 4, 5, 9, 10})
}
