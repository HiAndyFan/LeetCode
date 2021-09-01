/*
 * @lc app=leetcode id=70 lang=golang
 *
 * [70] Climbing Stairs
 */

// @lc code=start
func climbStairs(n int) int {
	rt := make([]int, n+1)
	rt[1] = 1
	rt[2] = 2
	for i := 3; i <= n; i++ {
		rt[i] = rt[i-1] + rt[i-2]
	}
	return rt[n]
}

// @lc code=end

