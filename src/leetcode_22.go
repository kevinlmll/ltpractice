package src

import (
	"strings"
)

func loopValidParenthesis(left_p_num, right_p_num, uncloused int, parenthesis []string) []string {
	if left_p_num == 0 && right_p_num == 0 {
		return []string{strings.Join(parenthesis, "")}
	}

	final_parenthesis := make([]string, 0)
	if uncloused == 0 {
		if left_p_num == 0 {
			return []string{}
		}

		parenthesis_new := make([]string, len(parenthesis))
		copy(parenthesis_new, parenthesis)
		parenthesis_new = append(parenthesis_new, "(")
		left_p_num_ := left_p_num - 1
		uncloused1_ := uncloused + 1
		final_parenthesis = append(final_parenthesis, loopValidParenthesis(left_p_num_, right_p_num, uncloused1_, parenthesis_new)...)
	} else {
		if left_p_num > 0 {
			parenthesis_new1 := make([]string, len(parenthesis))
			copy(parenthesis_new1, parenthesis)
			parenthesis_new1 = append(parenthesis_new1, "(")
			left_p_num_ := left_p_num - 1
			uncloused2_ := uncloused + 1
			final_parenthesis = append(final_parenthesis, loopValidParenthesis(left_p_num_, right_p_num, uncloused2_, parenthesis_new1)...)
		}
		if right_p_num > 0 {
			parenthesis_new2 := make([]string, len(parenthesis))
			copy(parenthesis_new2, parenthesis)
			parenthesis_new2 = append(parenthesis_new2, ")")
			right_p_num_ := right_p_num - 1
			uncloused3_ := uncloused - 1
			final_parenthesis = append(final_parenthesis, loopValidParenthesis(left_p_num, right_p_num_, uncloused3_, parenthesis_new2)...)
		}
	}
	return final_parenthesis
}

func generateParenthesis(n int) []string {
	initArray := make([]string, 0)
	return loopValidParenthesis(n, n, 0, initArray)
}
