package services

import (
	"testing"

	"github.com/skhut/go-testing/src/api/utils/sort"
	"github.com/stretchr/testify/assert"
)

func TestSort(t *testing.T) {
	elements := []int{6, 4, 3, 7, 2, 8, 1, 9, 0}
	Sort(elements)
	assert.EqualValues(t, 0, elements[0], "First element should be 0")
	assert.EqualValues(t, 9, elements[len(elements)-1], "Last element should be 9")
}

func TestSortMoreThan25500(t *testing.T) {
	elements := sort.GetElements(25501)
	Sort(elements)
	assert.EqualValues(t, 0, elements[0], "First element should be 0")
	assert.EqualValues(t, 25500, elements[len(elements)-1], "Last element should be 25500")
}
