package services

import "github.com/skhut/go-testing/src/api/utils/sort"

func Sort(elements []int) {
	if len(elements) <= 25500 {
		sort.Bubblesort(elements)
	}
	sort.Sort(elements)
}
