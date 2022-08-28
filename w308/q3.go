package w308

import "strings"

/**
给你一个下标从 0开始的字符串数组garbage，其中garbage[i]表示第 i个房子的垃圾集合。garbage[i]只包含字符'M'，'P' 和'G'，但可能包含多个相同字符，每个字符分别表示一单位的金属、纸和玻璃。垃圾车收拾 一单位的任何一种垃圾都需要花费1分钟。

同时给你一个下标从 0开始的整数数组travel，其中travel[i]是垃圾车从房子 i行驶到房子 i + 1需要的分钟数。

城市里总共有三辆垃圾车，分别收拾三种垃圾。每辆垃圾车都从房子 0出发，按顺序到达每一栋房子。但它们 不是必须到达所有的房子。

任何时刻只有 一辆垃圾车处在使用状态。当一辆垃圾车在行驶或者收拾垃圾的时候，另外两辆车 不能做任何事情。

请你返回收拾完所有垃圾需要花费的 最少总分钟数。

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/minimum-amount-of-time-to-collect-garbage
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func garbageCollection(garbage []string, travel []int) int {
	res := 0
	cm, cp, cg := 0, 0, 0
	move := func(c, t int) int {
		cost := 0
		for c < t {
			cost += travel[c]
			c++
		}
		return cost
	}
	for i, g := range garbage {
		r := strings.Count(g, "M")
		if r > 0 {
			res += r
			res += move(cm, i)
			cm = i
		}
		p := strings.Count(g, "P")
		if p > 0 {
			res += p
			res += move(cp, i)
			cp = i
		}
		g := strings.Count(g, "G")
		if p > 0 {
			res += g
			res += move(cg, i)
			cg = i
		}
	}
	return res

}
