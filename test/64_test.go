package test

import (
	"kevinlmll.github.com/leetcode/src"
	"testing"
)

func Test64(t *testing.T) {
	ma1 := [][]int{
		[]int{1, 3, 1},
		[]int{1, 5, 1},
		[]int{4, 2, 1},
	}

	ma2 := [][]int{
		[]int{1, 2, 3},
		[]int{4, 5, 6},
	}

	t.Logf("result:%v", src.ResultOfminPathSum(ma1))
	t.Logf("result:%v", src.ResultOfminPathSum(ma2))
}
