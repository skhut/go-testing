package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBubblesortOrderAsc(t *testing.T) {
	//Init
	elements := []int{9, 5, 7, 2, 8, 4, 6, 0, 1, 3}

	//Execution
	Bubblesort(elements)

	//Validation
	assert.EqualValues(t, 0, elements[0], "First element should be 0")
	assert.EqualValues(t, 9, elements[len(elements)-1], "Last element should be 9")
}

func TestSort(t *testing.T) {
	//Init
	elements := []int{9, 5, 7, 2, 8, 4, 6, 0, 1, 3}

	//Execution
	Sort(elements)

	//Validation
	assert.EqualValues(t, 0, elements[0], "First element should be 0")
	assert.EqualValues(t, 9, elements[len(elements)-1], "Last element should be 9")
}

func BenchmarkBubblesort(b *testing.B) {
	//Init
	//elements := []int{9, 5, 7, 2, 8, 4, 6, 0, 3}
	elements := GetElements(25490)
	//Execution
	for i := 0; i < b.N; i++ {
		Bubblesort(elements)
	}
}

func BenchmarkSort(b *testing.B) {
	//Init
	//elements := []int{9, 5, 7, 2, 8, 4, 6, 0, 3}
	elements := GetElements(25500)
	//Execution
	for i := 0; i < b.N; i++ {
		Sort(elements)
	}
}
