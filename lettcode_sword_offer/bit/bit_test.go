package bit

import (
	"testing"
)

// @Author XuPEngHao
// @DATE 2022/5/27 20:02

func Test_myPow(t *testing.T) {
	t.Log(myPowV2(2, -2))
	t.Log(myPowV2(2, 10))
}

func Test_hammingWeight(t *testing.T) {
	var num uint32 = 00000000000000000000000000001011
	num = 4294967293
	t.Log(hammingWeight(num))
}
