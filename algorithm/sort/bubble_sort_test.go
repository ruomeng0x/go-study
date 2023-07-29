package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBubbleSort(t *testing.T) {
	arr := []int{64, 34, 25, 12, 22, 11, 90}
	sortedArr := []int{11, 12, 22, 25, 34, 64, 90}
	BubbleSort(arr)
	assert.Equal(t, arr, sortedArr)
}
