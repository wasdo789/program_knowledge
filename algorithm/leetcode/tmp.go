package leetcode

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
