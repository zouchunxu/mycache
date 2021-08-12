package tests

import (
	"fmt"
	"sort"
	"testing"
)

func TestHash(t *testing.T) {
	keys := []int{1, 10, 20, 30}
	ma := make(map[int]int)
	var arr []int

	for _, key := range keys {
		for i := 0; i < 10; i++ {
			h := key + i
			arr = append(arr, h)
			ma[h] = key
		}
	}

	sort.Ints(arr)

	hash := 23
	ind := sort.Search(len(arr), func(i int) bool {
		return arr[i] >= hash
	})
	fmt.Println(ma[arr[ind%len(arr)]])
}
