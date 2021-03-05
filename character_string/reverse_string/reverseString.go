package main


/*
头尾双指针交换法
 */
func reverseString(s []byte)  {
	left := 0
	right := len(s)-1
	
	for left < right {
		temp := s[left]
		s[left] = s[right]
		s[right] = temp
		left ++ 
		right --
	}
}

func main()  {
	reverseString([]byte{'h','e','l','l','o'})
}