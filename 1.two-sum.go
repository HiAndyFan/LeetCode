/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
    ret := [2]int{0,0}
	for i,v := range(nums){
		for j,b := range(nums){
			if i!=j && v+b==target{
                ret[0]=i
                ret[1]=j
                break
			}
		}
	}
    return ret[:];
}
// @lc code=end

