//给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target 的那 两个 整数，并返回它们的数组下标。
//
// 你可以假设每种输入只会对应一个答案，并且你不能使用两次相同的元素。
//
// 你可以按任意顺序返回答案。
//
//
//
// 示例 1：
//
//
//输入：nums = [2,7,11,15], target = 9
//输出：[0,1]
//解释：因为 nums[0] + nums[1] == 9 ，返回 [0, 1] 。
//
//
// 示例 2：
//
//
//输入：nums = [3,2,4], target = 6
//输出：[1,2]
//
//
// 示例 3：
//
//
//输入：nums = [3,3], target = 6
//输出：[0,1]
//
//
//
//
// 提示：
//
//
// 2 <= nums.length <= 10⁴
// -10⁹ <= nums[i] <= 10⁹
// -10⁹ <= target <= 10⁹
// 只会存在一个有效答案
//
//
//
//
// 进阶：你可以想出一个时间复杂度小于 O(n²) 的算法吗？
//
// Related Topics 数组 哈希表 👍 19403 👎 0

package leetcode

import (
	"slices"
	"testing"
)

func TestTwoSum(t *testing.T) {
	cases := []struct {
		Case   string
		Input1 []int
		Input2 int
		Output []int
	}{
		{"1", []int{2, 7, 11, 15}, 9, []int{0, 1}},
		{"2", []int{3, 2, 4}, 6, []int{1, 2}},
		{"3", []int{3, 3}, 6, []int{0, 1}},
	}

	for _, c := range cases {
		t.Run(c.Case, func(t *testing.T) {
			output := twoSum(c.Input1, c.Input2)
			slices.Sort(output)
			slices.Sort(c.Output)
			if !slices.Equal(output, c.Output) {
				t.Fatalf("测试用例%v不通过，期望：%v，得到：%v", c.Case, c.Output, output)
			} else {
				t.Logf("测试用例%v通过了，期望：%v，得到：%v", c.Case, c.Output, output)
			}
		})
	}
}

// leetcode submit region begin(Prohibit modification and deletion)
func twoSum(nums []int, target int) []int {
	numberMaps := make(map[int]int)
	for index1, num := range nums {
		if index2, ok := numberMaps[target-num]; ok {
			return []int{index1, index2}
		}
		numberMaps[num] = index1
	}
	return []int{}
}

//leetcode submit region end(Prohibit modification and deletion)
