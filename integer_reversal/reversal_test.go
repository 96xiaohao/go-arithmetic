package integer_reversal

import (
	"go_arithmetic/integer_reversal"
	"testing"
)


func Test_001(t *testing.T)  {
	num := -123
	numRes := integer_reversal.Reversal(num)
	numTestRes := -321


	println(numRes==numTestRes)
	//assert.Equal(t,numRes,numTestRes)
}

func Test_002(t *testing.T)  {
	num := 12345
	numRes := integer_reversal.Reversal(num)
	numTestRes := 54321

	println(numRes==numTestRes)
}