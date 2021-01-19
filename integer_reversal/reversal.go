package integer_reversal

// 32位有符号整数的取值范围 -2^31-2^31-1
const (
	MAX_INTEGER = 2147483647    // 2^31-1
	MIN_INTEGER = -2147483648   // -2^31
)

func Reversal(num int ) int  {
	res :=0

	for num !=0 {
		last_num := num % 10

		if (res > MAX_INTEGER / 10) || ((res == MAX_INTEGER / 10) && last_num > 7) {
			return 0
		}

		if (res < MIN_INTEGER / 10) || ((res == MIN_INTEGER / 10) && last_num < -8) {
			return  0
		}

		res = res * 10 + last_num

		num /= 10
	}

	return res
}