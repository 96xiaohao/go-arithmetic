package main

import "fmt"

/*
给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。

说明：

你的算法应该具有线性时间复杂度。 你可以不使用额外空间来实现吗？

示例 1:

输入: [2,2,1]
输出: 1
示例 2:

输入: [4,1,2,1,2]
输出: 4

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/single-number
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

/*
异或运算有以下三个性质。

任何数和 0 做异或运算，结果仍然是原来的数，即 a ^ 0 = a
任何数和其自身做异或运算，结果是 0,即 a ^ a=0
异或运算满足交换律和结合律
 */



// 本题使用异或算法可以很快解决 //todo：注意其他元素都出现了偶数次才可以用异或算法
func singleNumber(nums []int) int {
	singleNum := 0

	for _,num := range nums{
		singleNum ^= num
	}

	return singleNum
}

func main()  {

	fmt.Println(singleNumber([]int{2,2,3,3,1}))

}