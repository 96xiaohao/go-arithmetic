package is_valid_bracket

func isValidBracket(str string)  bool{

	strLen := len(str)

	if strLen % 2 != 0 {
		return false
	}

	stark := []byte{}

	bracketMap := map[byte]byte{
		')':'(',
		']':'[',
		'}':'{',
	}

	for i:=0; i<strLen; i++{
		// stark[len(stark)-1] 栈顶元素
		// bracketMap[str[i]] map中的左括号
		if bracketMap[str[i]] >0 {
			if len(stark) == 0 || stark[len(stark)-1] != bracketMap[str[i]] {
				return false
			}else {
				// 弹出栈顶
				stark = stark[:len(stark)-1]
			}
		}else {
			stark = append(stark,str[i])
		}
	}


	return len(stark) == 0

}
