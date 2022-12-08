package quicksort_test

import (
	"testing"

	"github.com/AxelUser/algo/quicksort"
	"github.com/stretchr/testify/assert"
)

func TestSort_ShouldSortArray_ArrayIsNotSorted(t *testing.T) {
	arr := []int{3, 6, 5, 1, 2, 8}
	quicksort.Sort(arr)

	assert.Equal(t, []int{1, 2, 3, 5, 6, 8}, arr)
}

func TestSort_ShouldLeaveAsIs_ArrayIsSorted(t *testing.T) {
	arr := []int{1, 2, 3, 5, 6, 8}
	quicksort.Sort(arr)

	assert.Equal(t, []int{1, 2, 3, 5, 6, 8}, arr)
}
