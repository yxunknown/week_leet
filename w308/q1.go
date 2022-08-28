package w308

import "sort"

/**
给你一个长度为 n 的整数数组 nums ，和一个长度为 m 的整数数组 queries 。

返回一个长度为 m 的数组 answer ，其中 answer[i] 是 nums 中 元素之和小于等于 queries[i] 的 子序列 的 最大 长度 。

子序列 是由一个数组删除某些元素（也可以不删除）但不改变剩余元素顺序得到的一个数组。

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/longest-subsequence-with-limited-sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func answerQueries(nums []int, queries []int) []int {
	sort.Ints(nums)
	preSum := make([]int, len(nums))
	preSum[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		preSum[i] = preSum[i-1] + nums[i]
	}

	search := func(t int) int {
		l, r := 0, len(preSum)-1
		for l <= r {
			mid := (l + r) / 2
			if preSum[mid] <= t {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
		return l
	}

	res := make([]int, len(queries))
	for i, q := range queries {
		res[i] = search(q)
	}
	return res
}
