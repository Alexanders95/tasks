package tasks

import (
	"sort"
)

//Tasks 1
	/*
	Given an array A of strings made only from lowercase letters, return a list of all characters
	that show up in all strings within the list (including duplicates).  For example, if a character
	occurs 3 times in all strings but not 4 times, you need to include that character three times in the final answer.
	You may return the answer in any order.
	*/

func findDuplicates(a []byte) []byte {
	for i := 0; i < len(a); i++ {
		for a[i] != a[a[i]-1] {
			a[i], a[a[i]-1] = a[a[i]-1], a[i]
		}
	}

	res := make([]interface{}, 0, len(a))

	for i, n := range a {
		if byte(i) != n-1 {
			res = append(res, n)
		}
	}

	
	//...non-ended yet
	interface(res)

	var itemInterface []interface{}
	itemInterface = sort.Sort(res)
	byte(itemInterface)

	return itemInterface
}