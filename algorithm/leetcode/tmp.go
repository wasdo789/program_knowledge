package leetcode

import (
	"fmt"
	"sort"
	"strings"
)

func merge(nums1 []int, m int, nums2 []int, n int) {
	if len(nums2) == 0 {
		return
	}
	p1 := m - 1
	p2 := n - 1
	insertIdx := m + n - 1
	for p1 >= 0 && p2 >= 0 {
		if nums1[p1] > nums2[p2] {
			nums1[insertIdx] = nums1[p1]
			p1--
		} else {
			nums1[insertIdx] = nums2[p2]
			p2--
		}
		insertIdx--
	}
	if p2 >= 0 {
		for i := 0; i <= p2; i++ {
			nums1[i] = nums2[i]
		}
	}

}

func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}
	lastIdx := len(nums) - 1
	beginIdx := 0
	for beginIdx <= lastIdx {
		if nums[beginIdx] == val {
			nums[beginIdx], nums[lastIdx] = nums[lastIdx], nums[beginIdx]
			lastIdx--
		} else {
			beginIdx++
		}
	}
	return beginIdx
}

func removeDuplicates(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	//找到首个重复位置
	duplicateIdx := 1
	for ; duplicateIdx < len(nums); duplicateIdx++ {
		if nums[duplicateIdx] == nums[duplicateIdx-1] {
			break
		}
	}
	if duplicateIdx == len(nums) {
		return len(nums)
	}
	for idx := duplicateIdx + 1; idx < len(nums); {
		if nums[idx] == nums[duplicateIdx-1] {
			idx++
		} else {
			nums[duplicateIdx] = nums[idx]
			duplicateIdx++
			idx++
		}
	}
	return duplicateIdx
}

func removeDuplicates2(nums []int) int {
	if len(nums) <= 2 {
		return len(nums)
	}
	//找到首个重复位置
	duplicateIdx := 1
	duplicate := 1
	for ; duplicateIdx < len(nums); duplicateIdx++ {
		if nums[duplicateIdx] == nums[duplicateIdx-1] {
			duplicate++
		} else {
			duplicate = 1
		}
		if duplicate == 3 {
			break
		}
	}
	if duplicateIdx >= len(nums) {
		return len(nums)
	}
	for idx := duplicateIdx + 1; idx < len(nums); {
		if nums[idx] != nums[duplicateIdx-1] {
			duplicate = 1
		} else {
			duplicate++
		}
		if duplicate < 3 {
			nums[duplicateIdx] = nums[idx]
			duplicateIdx++
		}
		idx++
	}
	return duplicateIdx
}

func majorityElement(nums []int) int {
	if len(nums) <= 2 {
		return nums[0]
	}
	last := len(nums) - 1
	pre := last - 1

	for pre > 0 {
		for ; pre >= 0 && nums[pre] == nums[last]; pre-- {
		}
		if pre < 0 {
			break
		}
		nums[pre], nums[last-1] = nums[last-1], nums[pre]
		last -= 2
		if pre-1 <= last-1 {
			pre = pre - 1
		} else {
			pre = last - 1
		}

	}
	return nums[last]
}

func rotate(nums []int, k int) {
	l := len(nums)
	if l == 0 || k == 0 || k%l == 0 {
		return
	}
	//简化操作
	k = k % l
	//三次翻转法
	for i := 0; i < l/2; i++ {
		nums[i], nums[l-1-i] = nums[l-1-i], nums[i]
	}
	//反转前k个
	for i := 0; i <= (k-1)/2; i++ {
		nums[i], nums[k-1-i] = nums[k-1-i], nums[i]
	}
	//反转后l-k个
	for i := k; i <= (l+k-1)/2; i++ {
		nums[i], nums[k+l-1-i] = nums[k+l-1-i], nums[i]
	}
}

func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	profit := 0
	minPrices := prices[0]
	for i := 1; i < len(prices); i++ {
		curProfit := prices[i] - minPrices
		if curProfit > profit {
			profit = curProfit
		}
		if prices[i] < minPrices {
			minPrices = prices[i]
		}
	}
	return profit
}

func maxProfit2(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	profit := 0
	buyPrices := prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i] < buyPrices {
			buyPrices = prices[i]
		} else if prices[i] > buyPrices {
			profit += prices[i] - buyPrices
			buyPrices = prices[i]
		}
	}
	return profit

}

func canJump(nums []int) bool {
	if len(nums) <= 1 {
		return true
	}
	maxArrive := 0
	for idx := 0; idx <= maxArrive; idx++ {
		if idx+nums[idx] > maxArrive {
			maxArrive = idx + nums[idx]
			if maxArrive >= len(nums)-1 {
				return true
			}
		}

	}
	return false
}

// [2,3,1,1,4]
func jump(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}
	maxArrive := 0
	jumpTimes := 0
	jumpEnd := -1
	for idx := 0; idx < len(nums); idx++ {
		if idx > jumpEnd {
			jumpTimes++
			jumpEnd = maxArrive
		}
		if nums[idx]+idx > maxArrive {
			maxArrive = nums[idx] + idx
			if maxArrive >= len(nums)-1 {
				break
			}
		}

	}
	return jumpTimes
}
func canCompleteCircuit(gas []int, cost []int) int {
	if len(gas) == 0 {
		return -1
	}

	for i := 0; i < len(gas); {
		j := i
		restGas := 0
		for {
			restGas += gas[j] - cost[j]
			if restGas < 0 {
				break
			}
			j = (j + 1) % len(gas)
			if j == i {
				break
			}
		}
		if restGas >= 0 {
			return i
		}
		if j < i {
			return -1
		} else {
			i = j + 1
		}
	}

	return -1

}

func intToRoman(num int) string {
	if num >= 4000 {
		return ""
	}
	m := map[int]map[int]string{
		1: map[int]string{
			1: "I",
			5: "V",
			9: "IX",
		},
		10: map[int]string{
			1: "X",
			5: "L",
			9: "XC",
		},
		100: map[int]string{
			1: "C",
			5: "D",
			9: "CM",
		},
		1000: map[int]string{
			1: "M",
		},
	}

	res := ""
	level := 1
	for num > 0 {
		n := num % 10
		if n > 0 && n <= 3 {
			res = strings.Repeat(m[level][1], n) + res
		} else if n == 4 {
			res = m[level][1] + m[level][5] + res
		} else if n == 5 {
			res = m[level][5] + res
		} else if n > 0 && n < 9 {
			res = m[level][5] + strings.Repeat(m[level][1], n-5) + res
		} else if n == 9 {
			res = m[level][9] + res
		}
		num /= 10
		level *= 10
	}
	return res
}

func lengthOfLastWord(s string) int {
	end := len(s) - 1
	i := len(s) - 1
	flag := false
	for ; i >= 0; i-- {
		if s[i] != ' ' && !flag {
			end = i
			flag = true
		}
		if s[i] == ' ' && flag {
			break
		}
	}
	return end - i
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
	i := 0
	for ; i < len(strs[0]); i++ {
		for j := 1; j < len(strs); j++ {
			if strs[0][i] != strs[j][i] {
				return strs[0][:i]
			}
		}
	}
	return strs[0][:i]
}

func reverseWords(s string) string {
	strList := strings.Split(s, " ")
	res := ""

	for i := len(strList) - 1; i >= 0; i-- {
		if strList[i] != "" {
			res += strList[i] + " "
		}

	}
	return strings.TrimRight(res, " ")

}

// func convert(s string, numRows int) string {
// 	if len(s) <= 1 || numRows <= 1 {
// 		return s
// 	}
// 	res := make([]byte, len(s))
// 	maxInterval := 2*numRows - 2
// 	idx := 0
// 	for i := 0; i < numRows; i++ {
// 		rowInterval := maxInterval - i*2
// 		if rowInterval == 0 {
// 			rowInterval = maxInterval
// 		}
// 		for j := i; j < len(s); {
// 			res[idx] = s[j]
// 			idx++
// 			j = j + rowInterval
// 			if maxInterval-rowInterval > 0 {
// 				rowInterval = maxInterval - rowInterval
// 			}
// 		}
// 	}
// 	return string(res)
// }

func convert(s string, numRows int) string {
	if len(s) <= 1 || numRows <= 1 {
		return s
	}
	strList := make([][]byte, numRows)
	rowIdx := 0
	flag := 1
	for i := 0; i < len(s); i++ {
		strList[rowIdx] = append(strList[rowIdx], s[i])
		rowIdx = rowIdx + flag
		if rowIdx == numRows-1 || rowIdx == 0 {
			flag = -flag
		}
	}
	res := ""
	for _, str := range strList {
		res += string(str)
	}
	return res
}

func maxArea(height []int) int {
	if len(height) <= 1 {
		return 0
	}
	var area int
	left := 0
	leftMax := 0
	right := len(height) - 1
	rightMax := 0
	for left < right {
		var curArea int
		if height[left] > leftMax {
			leftMax = height[left]
		}
		if height[right] > rightMax {
			rightMax = height[right]
		}
		if height[left] <= height[right] {
			curArea = height[left] * (right - left)
			for ; left < right && height[left] <= leftMax; left++ {
			}
		} else {
			curArea = height[right] * (right - left)
			for ; left < right && height[right] <= rightMax; right-- {
			}
		}
		if curArea > area {
			area = curArea
		}
	}
	return area
}

func threeSum(nums []int) [][]int {
	if len(nums) == 0 {
		return nil
	}
	res := make([][]int, 0)
	//排序
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left := i + 1
		right := len(nums) - 1
		for left < right {
			if left > i+1 && left < right && nums[left] == nums[left-1] {
				left++
				continue
			}
			if right < len(nums)-1 && left < right && nums[right] == nums[right+1] {
				right--
				continue
			}
			if nums[left]+nums[right]+nums[i] < 0 {
				left++

			} else if nums[left]+nums[right]+nums[i] > 0 {
				right--
			} else {
				res = append(res, []int{nums[i], nums[left], nums[right]})
				left++
				right--
			}

		}
	}
	return res
}

func minSubArrayLen(target int, nums []int) int {
	left := 0
	right := 0
	sum := 0
	minLen := len(nums) + 1
	for right < len(nums) {
		if nums[right] >= target {
			return 1
		}
		//右边进
		sum += nums[right]
		if sum < target {
			right++
		} else {
			//左边出
			for ; left <= right && sum >= target; left++ {
				sum -= nums[left]
			}
			curLen := right - left + 2
			if minLen > curLen {
				minLen = curLen
			}
			right++

		}
	}
	if minLen <= len(nums) {
		return minLen
	}
	return 0
}

// kmp方法找到s中所有p出现的位置
func prefixStrIdx(s string, t string) []int {
	if len(s) == 0 || len(t) == 0 {
		return nil
	}
	sp := t + "#" + s
	pi := make([]int, len(s)+1+len(t))
	ocr := make([]int, 0, len(s))
	for i := 1; i < len(sp); i++ {
		j := i - 1
		for ; j >= 0 && sp[pi[j]] != sp[i]; j = pi[j] - 1 {
		}
		if j >= 0 {
			pi[i] = pi[j] + 1
		} else {
			pi[i] = 0
		}
		if pi[i] == len(t) {
			ocr = append(ocr, i-2*len(t))
		}
	}
	return ocr
}

func findSubstring(s string, words []string) []int {
	if len(s) == 0 || len(words) == 0 || len(words[0]) == 0 {
		return nil
	}
	if len(words) == 1 && s == words[0] {
		return []int{0}
	}
	//
	allWordMap := map[string]int{}
	pos2WordType := map[int]string{}
	occrPosArr := make([]int, 0, len(s))
	for _, pat := range words {
		if _, ok := allWordMap[pat]; ok {
			allWordMap[pat] = allWordMap[pat] + 1
			continue
		} else {
			occr := prefixStrIdx(s, pat)
			for _, pos := range occr {
				pos2WordType[pos] = pat
				occrPosArr = append(occrPosArr, pos)
			}
			allWordMap[pat] = 1
		}

	}
	//sort occrArr
	sort.Ints(occrPosArr)
	//
	curWordType2Idx := map[string][]int{}
	beginIdx := 0
	lastIdx := 0
	curWordType2Idx[pos2WordType[occrPosArr[0]]] = []int{0}
	num := 1
	res := make([]int, 0, len(s))
	for i := 1; i < len(occrPosArr); i++ {
		if occrPosArr[i]-occrPosArr[lastIdx] < len(words[0]) {
			continue
		} else if occrPosArr[i]-occrPosArr[lastIdx] > len(words[0]) {
			curWordType2Idx = map[string][]int{
				pos2WordType[occrPosArr[i]]: []int{i},
			}
			lastIdx = i
			beginIdx = i
			num = 1
			continue
		} else {
			lastIdx = i
			wordType := pos2WordType[occrPosArr[i]]
			idxList, ok := curWordType2Idx[wordType]
			if !ok || len(idxList) < allWordMap[wordType] {
				curWordType2Idx[wordType] = append(curWordType2Idx[wordType], i)
				num++
			} else {
				//curWordType2Idx 弹出小于等于idx的所有元素
				curWordType2Idx = map[string][]int{}
				num = 0
				for j := idxList[0] + 1; j <= i; j++ {
					curWordType2Idx[pos2WordType[occrPosArr[j]]] = append(curWordType2Idx[pos2WordType[occrPosArr[j]]], j)
					num++
				}
				beginIdx = idxList[0] + 1
			}
		}
		if num == len(words) {
			res = append(res, occrPosArr[beginIdx])
			pos2WordType[occrPosArr[beginIdx]] = pos2WordType[occrPosArr[beginIdx]][1:]
			beginIdx = beginIdx + 1
		}
	}

	return res
}

func copyMap(src, dst map[byte]int) {
	for k, v := range src {
		dst[k] = v
	}
}
func minWindow(s string, t string) string {
	if s == "" || t == "" || len(s) < len(t) {
		return ""
	}
	allCharMap := map[byte]int{}
	for i := 0; i < len(t); i++ {
		allCharMap[t[i]] = allCharMap[t[i]] + 1
	}
	curWindowMap := map[byte]int{}
	copyMap(allCharMap, curWindowMap)
	wordSta := map[byte]int{}
	curIdxArr := []int{}
	for i := 0; i < len(s); i++ {
		if allCharMap[s[i]] == 0 {
			continue
		} else {
			curIdxArr = append(curIdxArr, i)
			wordSta[s[i]] = wordSta[s[i]] + 1
			v, ok := curWindowMap[s[i]]
			if ok {
				if v > 1 {
					curWindowMap[s[i]] = v - 1
				} else {
					delete(curWindowMap, s[i])
				}
			}
			if len(curWindowMap) == 0 {
				break
			}
		}
	}
	if len(curWindowMap) != 0 {
		return ""
	}
	if len(curIdxArr) == 1 {
		return t
	}

	minLen := curIdxArr[len(curIdxArr)-1] - curIdxArr[0] + 1
	k := curIdxArr[len(curIdxArr)-1] + 1
	res := s[curIdxArr[0]:k]
	for j := 1; j < len(curIdxArr); j++ {
		w := s[curIdxArr[j-1]]
		//fmt.Printf("begin %d %c\n", j, s[curIdxArr[j]])
		wordSta[w] = wordSta[w] - 1
		if wordSta[w] < allCharMap[w] {
			for ; k < len(s); k++ {
				//fmt.Printf("end %d %c, str %s\n", k, s[k], s[curIdxArr[j]:k+1])
				if allCharMap[s[k]] == 0 {
					continue
				} else {
					curIdxArr = append(curIdxArr, k)
					wordSta[s[k]] = wordSta[s[k]] + 1
					if wordSta[w] >= allCharMap[w] {
						k++
						break
					}
				}
			}
		}
		if wordSta[w] >= allCharMap[w] {
			curLen := curIdxArr[len(curIdxArr)-1] - curIdxArr[j] + 1
			if curLen < minLen {
				minLen = curLen
				res = s[curIdxArr[j] : curIdxArr[j]+minLen]
			}
		} else {
			break
		}

	}

	return res
}

func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return nil
	}
	res := []string{}
	begin := 0
	i := 1
	for i < len(nums) {
		if nums[i]-nums[i-1] == 1 {
			i++
		} else {
			if i-1 == begin {
				res = append(res, fmt.Sprint(nums[begin]))
			} else {
				res = append(res, fmt.Sprintf("%d->%d", nums[begin], nums[i-1]))
			}
			begin = i
			i++
		}
	}
	if i-1 == begin {
		res = append(res, fmt.Sprint(nums[begin]))
	} else {
		res = append(res, fmt.Sprintf("%d->%d", nums[begin], nums[i-1]))
	}
	return res
}

func myMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// return 包含end的最大值，和截止到end的最大值
func maxSubArraySum(nums []int, end int) (int, int) {
	if end == 0 {
		//fmt.Printf("end:%d, sum:%d\n", 0, nums[end])
		return nums[0], nums[0]
	}
	sum, maxNum := maxSubArraySum(nums, end-1)

	fmt.Printf("end:%d, sum:%d\n", end-1, sum)
	var nextSum int
	if sum <= 0 {
		nextSum = nums[end]
	} else {
		nextSum = sum + nums[end]
	}
	if nextSum > maxNum {
		return nextSum, nextSum
	} else {
		return nextSum, maxNum
	}

}

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	_, sum := maxSubArraySum(nums, len(nums)-1)
	//fmt.Printf("end:%d, endPos:%d, sum:%d, aftersum:%d\n", len(nums)-1, endPos, sum, aftersum)
	return sum
}
