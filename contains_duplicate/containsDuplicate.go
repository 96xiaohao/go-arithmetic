package main

import (
	"fmt"
	"sort"
)

// 先对nums进行排序，然后判断相邻的两个元素是否相等
func containsDuplicate(nums []int) bool {
	sort.Ints(nums)

	for i:=0; i<len(nums)-1; i++{
		if nums[i] == nums[i+1] {
			return true
		}

	}
	return false
}

// 利用map存储遍历过的元素
func containsDuplicate2(nums []int) bool{
	numsMap := map[int]int{}

	for _, num := range nums {
		if _,ok := numsMap[num]; ok {
			return true
		}else {
			numsMap[num] = 1
		}
	}

	return false
}


func main()  {
	fmt.Println(containsDuplicate([]int{2,3,4,5,1,3}))
	fmt.Println(containsDuplicate2([]int{2,3,4}))
}