package _015_three_sum
/*
 * @lc app=leetcode.cn id=15 lang=golang
 *
 * [15] 三数之和
 *
 * https://leetcode-cn.com/problems/3sum/description/
 *
 * algorithms
 * Medium (22.51%)
 * Likes:    1015
 * Dislikes: 0
 * Total Accepted:    58.1K
 * Total Submissions: 258.1K
 * Testcase Example:  '[-1,0,1,2,-1,-4]'
 *
 * 给定一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0
 * ？找出所有满足条件且不重复的三元组。
 *
 * 注意：答案中不可以包含重复的三元组。
 *
 * 例如, 给定数组 nums = [-1, 0, 1, 2, -1, -4]，
 *
 * 满足要求的三元组集合为：
 * [
 * 	[-1, 0, 1],
 * 	[-1, -1, 2]
 * ]
 *
 *
 */

import (
	"sort"
)

// time complexity: O(n^2)
func threeSum(nums []int) [][]int {
	// 正序排序
	sort.Ints(nums)
	ret := [][]int{}

	// 三元组，只需遍历到倒数第三个
	for i :=0;i < len(nums)-2;i++{
		// 当前移动元素已经大于0，则退出
		if nums[i] > 0 {
			break
		}

		// 当前移动元素和上一个移动元素相等，则不再搜索
		if i> 0 && nums[i] == nums[i-1] {
			continue
		}

		target := 0- nums[i]
		// j，k 分别指向两侧
		j,k := i+1,len(nums) -1
        for j < k {
        	if nums[j] + nums[k] == target {
				// 去重
				if j > i+1 && nums[j] == nums[j-1] {
					j++
					continue
				}
				// 去重
				if k < len(nums)-1 && nums[k] == nums[k +1] {
					k--
					continue
				}
        		ret = append(ret,[]int{nums[i],nums[j],nums[k]})
        		j++
        		k--
			}else if nums[j] + nums[k]  > target {
				k--
			}else {
				j++
			}
		}


	}

	return ret
}


