package structure

// manacher 算法
// 求最大回文子串的长度
func manacher(s string) int {
	str := fillStr(s)
	p := make([]int, len(str)) //回文半径
	c := 0                     //中心位置
	r := -1                    //右边界加1
	maxLen := -1
	for i := 0; i < len(str); i++ {
		if i > r || 2*c-i < 0 {
			p[i] = 1
		} else {
			if r-i < p[2*c-i] {
				p[i] = p[2*c-i]
			} else {
				p[i] = r - i
			}
		}
		//继续扩展
		for i+p[i] < len(str) && i-p[i] > -1 {
			if str[i+p[i]] != str[i-p[i]] {
				break
			} else {
				p[i]++
			}
		}
		if i+p[i] > r {
			r = p[i]
			c = i
		}
		if maxLen < r {
			maxLen = r
		}
	}
	return maxLen - 1
}

func fillStr(s string) string {
	res := make([]byte, 2*len(s)+1)
	for i := 0; i < len(res); i++ {
		if i%2 == 0 {
			res[i] = '#'
		} else {
			res[i] = s[i/2]
		}
	}
	return string(res)
}

//kmp 算法
