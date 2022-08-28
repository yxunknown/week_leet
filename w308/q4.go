package w308

/**
给你一个 正整数k，同时给你：

一个大小为 n的二维整数数组rowConditions，其中rowConditions[i] = [abovei, belowi]和
一个大小为 m的二维整数数组colConditions，其中colConditions[i] = [lefti, righti]。
两个数组里的整数都是1到k之间的数字。

你需要构造一个k x k的矩阵，1到k每个数字需要恰好出现一次。剩余的数字都是0。

矩阵还需要满足以下条件：

对于所有 0到n - 1之间的下标i，数字abovei所在的 行必须在数字belowi所在行的上面。
对于所有 0到 m - 1之间的下标i，数字lefti所在的 列必须在数字righti所在列的左边。

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/build-a-matrix-with-conditions
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func buildMatrix(k int, rowConditions [][]int, colConditions [][]int) [][]int {
	topV := make([]int, k)
	topH := make([]int, k)
	for _, con := range rowConditions {

	}
}
