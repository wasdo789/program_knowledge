package leetcode

/*
位运算
*/

/*
&按位与
1. 提取数字二进制形式最右侧的1
2. 数字二进制形式1的个数
*/

// 提取数字二进制形式最右侧的1
func LowBitOne(a int) int {
	return a & (^a + 1)
}

// 数字二进制形式1的个数
func OneCount(a int) int {
	count := 0
	for a != 0 {
		x := a & (^a + 1)
		count++
		a = a ^ x
	}
	return count
}

/*
异或应用，异或理解为无进位加法
1. 不用额外变量交换a,b a=a^b,b=a^b,a=a^b
2. 找唯一一个出现奇数次的数字
3. 找出一堆数里唯二的，出现奇数次的数字。
*/

func FindOddTimesNumber(data []int) int {
	eor := 0
	for _, d := range data {
		eor = eor ^ d
	}
	return eor
}

func Find2OddTimesNumber(data []int) (int, int) {
	//假设两个奇数次出现的数字是a，b，首先得到a异或b
	eor := 0
	for _, d := range data {
		eor = eor ^ d
	}
	//因为a!=b，那么eor必然有一位是1，取得最低位1的，假设在第n位
	lowbit1 := eor & (^eor + 1)
	//把d中所有数按照第n位是1和0分位两组，相当于转换成了找1个奇数次的数字
	eor2 := 0
	for _, d := range data {
		if d&lowbit1 != 0 {
			eor2 = eor2 ^ d
		}
	}
	return eor2, eor ^ eor2

}
