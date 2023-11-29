package leetcode

import (
	"slices"
	"sort"
)

type SmallestInfiniteSet struct {
	small     int
	smallList []int
	addList   []int
}

func Constructor() SmallestInfiniteSet {
	return SmallestInfiniteSet{
		small:     1,
		smallList: make([]int, 0),
		addList:   make([]int, 0),
	}
}

func (this *SmallestInfiniteSet) PopSmallest() int {
	var result int
	if len(this.addList) == 0 {
		result = this.small
		this.small++
		this.smallList = append(this.smallList, result)
	} else {
		sort.Ints(this.addList)
		result = this.addList[0]
		this.addList = append(this.addList[:0], this.addList[1:]...)
	}
	return result
}

func (this *SmallestInfiniteSet) AddBack(num int) {
	if slices.Contains(this.smallList, num) && !slices.Contains(this.addList, num) {
		this.addList = append(this.addList, num)
	}
}
