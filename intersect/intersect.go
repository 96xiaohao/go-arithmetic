package main

import (
	"fmt"
	"sort"
)

/*
给定两个数组，编写一个函数来计算它们的交集。

示例 1：

输入：nums1 = [1,2,2,1], nums2 = [2,2]
输出：[2,2]
示例 2:

输入：nums1 = [4,9,5], nums2 = [9,4,9,8,4]
输出：[4,9]
 

说明：

输出结果中每个元素出现的次数，应与元素在两个数组中出现次数的最小值一致。
我们可以不考虑输出结果的顺序。
进阶：

如果给定的数组已经排好序呢？你将如何优化你的算法？
如果 nums1 的大小比 nums2 小很多，哪种方法更优？
如果 nums2 的元素存储在磁盘上，内存是有限的，并且你不能一次加载所有的元素到内存中，你该怎么办？

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/intersection-of-two-arrays-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */



/*
关键点：hashMap
思路：
1.先遍历长度较小的数组，将数组中的元素作为hashMap的key，value为1存储到map中
2.申请结果数组存储交集
3.遍历长度较长的数组，判断数组中的元素是否存在，如果存在，则把元素写入结果数组，
  同时把hashMap中key删除,当hashMap中没有元素时，返回结果集
返回去重后的交集
 */
func intersect(nums1 []int, nums2 []int) []int {
	minNums := nums1
	maxNums := nums2
	if len(nums1) > len(nums2) {
		minNums = nums2
		minNums = nums1
	}

	hashMap := map[int]int{}

	for _,num := range minNums{
		if _,ok := hashMap[num]; ok {

		}else {
			hashMap[num] = 1
		}
	}

	resultNums := []int{}

	for _,num := range maxNums{
		if len(hashMap) == 0 {
			return resultNums
		}else {
			if _,ok := hashMap[num];ok {
				resultNums = append(resultNums,num)
				delete(hashMap,num)
			}
		}
	}

	return resultNums
}


/*
关键点：hashMap
思路：
1.先遍历长度较小的数组，将数组中的元素作为hashMap的key，value为1存储到map中
2.申请结果数组存储交集
3.遍历长度较长的数组，判断数组中的元素是否存在，如果存在，则把元素写入结果数组，
  同时把hashMap中key删除,当hashMap中没有元素时，返回结果集
返回不去重的交集
*/
func intersect2(nums1 []int, nums2 []int) []int {
	minNums := nums1
	maxNums := nums2
	if len(nums1) > len(nums2) {
		minNums = nums2
		minNums = nums1
	}

	hashMap := map[int]int{}

	for _,num := range minNums{
		hashMap[num] ++
	}

	resultNums := []int{}

	for _,num := range maxNums{
		if hashMap[num] > 0{
			resultNums = append(resultNums,num)
			hashMap[num] --
		}
	}

	return resultNums
}

/*
关键点：数组有序；双指针
思路：先把两个数组排序，然后同时遍历两个数组求交集
 */
func intersect3(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)

	length1 := len(nums1)
	length2 := len(nums2)

	index1 := 0
	index2 := 0

	resultNums := []int{}
	for index1 < length1 && index2 < length2 {
		if nums1[index1] < nums2[index2] {
			index1 ++
		}else if nums1[index1] > nums2[index2] {
			index2 ++
		}else {
			resultNums = append(resultNums,nums1[index1])
			index1 ++
			index2 ++
		}
	}

	return resultNums
}




func main()  {
	num1 := []int{1,2,4,2,3,4}
	num2 := []int{3,2,2}
	fmt.Println(intersect3(num1,num2))
}