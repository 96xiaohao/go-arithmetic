package is_Palindrome

func IsPalindrome(num int)  bool{
	// 负数及个位为0的整数不是回文数
	if num < 0 || (num %10) ==0  {
		return false
	}

	if num ==0 {
		return true
	}
	revertedNumber := 0

	for num > revertedNumber {
		revertedNumber = revertedNumber*10 + num % 10
		num /= 10
	}

	return num == revertedNumber || num == revertedNumber/10
}