package w308

/**
给你一个包含若干星号 * 的字符串 s 。

在一步操作中，你可以：

选中 s 中的一个星号。
移除星号 左侧 最近的那个 非星号 字符，并移除该星号自身。
返回移除 所有 星号之后的字符串。

注意：

生成的输入保证总是可以执行题面中描述的操作。
可以证明结果字符串是唯一的。

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/removing-stars-from-a-string
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func removeStars(s string) string {
	bytes := []byte(s)
	res := make([]byte, 0)
	for _, c := range bytes {
		if c != '*' {
			res = append(bytes, c)
		} else {
			res = append(res[:len(res)-1])
		}
	}
	return string(res)
}
