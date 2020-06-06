package sort

import (
	"testing"
)

func TestBubblesortOrderDesc(t *testing.T) {
	//Init
	elements := []int{9, 8, 5, 3, 7, 4, 2, 0, 1}

	//Execution
	Bubblesort(elements)

	//Validation
	if elements[0] != 9 {
		t.Error("First element should be 9")
	}

	if elements[len(elements)-1] != 0 {
		t.Error("Last element should be 0")
	}
}

func TestBubblesortAlreadySorted(t *testing.T) {
	//Init
	elements := []int{9, 8, 7, 6, 5, 4}

	//Execution
	Bubblesort(elements)

	//Validation
}
