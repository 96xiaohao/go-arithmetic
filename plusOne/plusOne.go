package main

import "fmt"

/*
给定一个由 整数 组成的 非空 数组所表示的非负整数，在该数的基础上加一。

最高位数字存放在数组的首位， 数组中每个元素只存储单个数字。

你可以假设除了整数 0 之外，这个整数不会以零开头。

示例 1：

输入：digits = [1,2,3]
输出：[1,2,4]
解释：输入数组表示数字 123。
示例 2：

输入：digits = [4,3,2,1]
输出：[4,3,2,2]
解释：输入数组表示数字 4321。
示例 3：

输入：digits = [0]
输出：[1]
 

提示：

1 <= digits.length <= 100
0 <= digits[i] <= 9

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/plus-one
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */


/*
关键点：加一后进位；对10求余进位
思路：
1.从数组末端加一，判断加一后对10求余，求余=表示要进位
2.从右到左循环判断数组中的每个元素是否需要进位，若不进位则跳出循环返回数组
3.如果从右到左完成循环后，数组中的0下表未知也需要进位，则需要扩展数组长度一位，0位置至为1
 */
func plusOne(digits []int) []int {

	digitsLen := len(digits)
	for i := digitsLen-1; i>=0; i-- {
		digits[i] ++
		digits[i] %= 10
		if digits[i] != 0 {
			return digits
		}

	}

	digits = append([]int{1},digits...)

	return digits
}

func main()  {
	fmt.Println(plusOne([]int{1,2,9,9}))
}