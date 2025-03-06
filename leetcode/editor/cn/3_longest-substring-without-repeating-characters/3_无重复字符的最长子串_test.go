//给定一个字符串 s ，请你找出其中不含有重复字符的 最长 子串 的长度。
//
//
//
// 示例 1:
//
//
//输入: s = "abcabcbb"
//输出: 3
//解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
//
//
// 示例 2:
//
//
//输入: s = "bbbbb"
//输出: 1
//解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
//
//
// 示例 3:
//
//
//输入: s = "pwwkew"
//输出: 3
//解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
//     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
//
//
//
//
// 提示：
//
//
// 0 <= s.length <= 5 * 10⁴
// s 由英文字母、数字、符号和空格组成
//
//
// Related Topics 哈希表 字符串 滑动窗口 👍 10657 👎 0

package leetcode

import (
	"testing"
)

func TestLongestSubstringWithoutRepeatingCharacters(t *testing.T) {
	cases := []struct {
		Case   string
		Input  string
		Output int
	}{
		{"1", "abcabcbb", 3},
		{"2", "bbbbb", 1},
		{"3", "pwwkew", 3},
		{"4", "djiqhfjweniofjqipdjfqwinmfipqwjpofjqwpfd", 9},
	}

	for _, c := range cases {
		t.Run(c.Case, func(t *testing.T) {
			if output := lengthOfLongestSubstring(c.Input); output != c.Output {
				t.Fatalf("测试用例%v不通过，期望：%v，得到：%v", c.Case, c.Output, output)
			} else {
				t.Logf("测试用例%v通过了，期望：%v，得到：%v", c.Case, c.Output, output)
			}
		})
	}
}

//  a b c a b c b b
//  0 1 2 3 4 5 6 7

// max 0
// curr 3
// index 0

// leetcode submit region begin(Prohibit modification and deletion)
func lengthOfLongestSubstring(s string) int {
	s1 := []byte(s)
	n := len(s1)
	if n == 0 {
		return 0
	}
	maxLen := 0
	for index, char := range s1 {
		currentLen := 1
		strMap := map[string]int{
			string(char): index,
		}
		for index2 := index + 1; index2 < n; index2++ {
			i, ok := strMap[string(s1[index2])]
			if ok {
				index2 = i
				break
			} else {
				strMap[string(s1[index2])] = index2
				currentLen++
			}
		}
		if currentLen > maxLen {
			maxLen = currentLen
		}
	}
	return maxLen
}

//leetcode submit region end(Prohibit modification and deletion)
