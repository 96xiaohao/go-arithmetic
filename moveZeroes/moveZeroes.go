package main

import "fmt"

/*
给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。

示例:

输入: [0,1,0,3,12]
输出: [1,3,12,0,0]
说明:

必须在原数组上操作，不能拷贝额外的数组。
尽量减少操作次数。


来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/move-zeroes
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */


/*
关键点：双指针；一次循环；交换0元素交换的时机
思路：
1.循环数组，当元素不等于0时就交换，右指针一直++探路
2.当发生元素交换时，左指针++

 */
func moveZeroes(nums []int) []int {
	left := 0
	right := 0

	for right < len(nums) {
		if nums[right] != 0{
			temp := nums[left]
			nums[left] =  nums[right]
			nums[right] = temp
			left ++
		}
		right++
	}

	return nums
}

/*
关键点：双指针；一次循环；交换0元素交换的时机
思路：
1.首先确定元素交换时机，
2.只有左指针对应的元素=0，且右元素对应的元素不等于0时才交换元素
3.其他情况控制指针++即可

 */
func moveZeroes2(nums []int) []int {
	left := 0
	right := 0

	for right < len(nums) {
		if nums[right] != 0 && nums[left] !=0 {
			left++
			right++
		}else if nums[right] != 0 && nums[left] ==0 {
			temp := nums[left]
			nums[left] =  nums[right]
			nums[right] = temp
			left ++
			right ++
		}else if nums[left] ==0 && nums[right] == 0{
			right ++
		}
	}

	return nums
}

func main()  {
	fmt.Println(moveZeroes2([]int{2,3,4,0,0,6,7,0,8}))

}