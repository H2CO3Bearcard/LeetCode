//给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。
//
// 请你将两个数相加，并以相同形式返回一个表示和的链表。
//
// 你可以假设除了数字 0 之外，这两个数都不会以 0 开头。
//
//
//
// 示例 1：
//
//
//输入：l1 = [2,4,3], l2 = [5,6,4]
//输出：[7,0,8]
//解释：342 + 465 = 807.
//
//
// 示例 2：
//
//
//输入：l1 = [0], l2 = [0]
//输出：[0]
//
//
// 示例 3：
//
//
//输入：l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
//输出：[8,9,9,9,0,0,0,1]
//
//
//
//
// 提示：
//
//
// 每个链表中的节点数在范围 [1, 100] 内
// 0 <= Node.val <= 9
// 题目数据保证列表表示的数字不含前导零
//
//
// Related Topics 递归 链表 数学 👍 11115 👎 0

package leetcode

import (
	"strconv"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	cases := []struct {
		Case   string
		Input1 *ListNode
		Input2 *ListNode
		Output *ListNode
	}{
		{
			"1",
			&ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3, Next: nil}}},
			&ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4, Next: nil}}},
			&ListNode{Val: 7, Next: &ListNode{Val: 0, Next: &ListNode{Val: 8, Next: nil}}},
		}, {
			"2",
			&ListNode{Val: 0, Next: nil},
			&ListNode{Val: 0, Next: nil},
			&ListNode{Val: 0, Next: nil},
		}, {
			"3",
			&ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: nil}}}}}}},
			&ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: nil}}}},
			&ListNode{Val: 8, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 0, Next: &ListNode{Val: 0, Next: &ListNode{Val: 0, Next: &ListNode{Val: 1, Next: nil}}}}}}}},
		},
	}

	for _, c := range cases {
		t.Run(c.Case, func(t *testing.T) {
			if output := addTwoNumbers(c.Input1, c.Input2); output.ToString() != c.Output.ToString() {
				t.Fatalf("测试用例%v不通过，期望：%v，得到：%v", c.Case, c.Output.ToString(), output.ToString())
			} else {
				t.Logf("测试用例%v通过了，期望：%v，得到：%v", c.Case, c.Output.ToString(), output.ToString())
			}
		})
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) ToString() string {
	if l.Next == nil {
		return strconv.Itoa(l.Val)
	} else {
		return strconv.Itoa(l.Val) + l.Next.ToString()
	}
}

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1.Val+l2.Val >= 10 {
		if l1.Next == nil && l2.Next == nil {
			return &ListNode{Val: (l1.Val + l2.Val) % 10, Next: &ListNode{Val: (l1.Val + l2.Val) / 10, Next: nil}}
		}
		if l1.Next == nil {
			return &ListNode{Val: (l1.Val + l2.Val) % 10, Next: addTwoNumbers(l2.Next, &ListNode{Val: (l1.Val + l2.Val) / 10, Next: nil})}
		}
		if l2.Next == nil {
			return &ListNode{Val: (l1.Val + l2.Val) % 10, Next: addTwoNumbers(l1.Next, &ListNode{Val: (l1.Val + l2.Val) / 10, Next: nil})}
		}
		return &ListNode{Val: (l1.Val + l2.Val) % 10, Next: addTwoNumbers(addTwoNumbers(l1.Next, l2.Next), &ListNode{Val: (l1.Val + l2.Val) / 10, Next: nil})}
	} else {
		if l1.Next == nil && l2.Next == nil {
			return &ListNode{Val: l1.Val + l2.Val, Next: nil}
		}
		if l1.Next == nil {
			return &ListNode{Val: l1.Val + l2.Val, Next: l2.Next}
		}
		if l2.Next == nil {
			return &ListNode{Val: l1.Val + l2.Val, Next: l1.Next}
		}
		return &ListNode{Val: l1.Val + l2.Val, Next: addTwoNumbers(l1.Next, l2.Next)}
	}

}

//leetcode submit region end(Prohibit modification and deletion)
