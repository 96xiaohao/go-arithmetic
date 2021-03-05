package main

/*
字符串中的第一个唯一字符

给定一个字符串，找到它的第一个不重复的字符，并返回它的索引。如果不存在，则返回 -1。



示例：

s = "leetcode"
返回 0

s = "loveleetcode"
返回 2



提示：你可以假定该字符串只包含小写字母。

作者：力扣 (LeetCode)
链接：https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/xn5z8r/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
 */


/*
关键点：哈希map
思路：两次遍历，第一次遍历统计出所有字符出现的次数，第二次遍历找到出现了一次的第一个字符，返回下标
注意：字符串只包含小写字母，那么声明hashMap的长度<=26
 */

func firstUniqChar(s string)  int{
	chn := [26]int{}

	for _,char := range s{
		chn[char-'a'] ++
	}

	for i,char := range s{
		if chn[char-'a'] ==1 {
			return i
		}
	}

	return -1
}

func main() {
	println(firstUniqChar("loveleetcode"))
}