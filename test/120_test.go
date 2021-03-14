package test

import (
	"kevinlmll.github.com/leetcode/src"
	"testing"
)

func Test120(t *testing.T) {
	a := [][]int{
		{2},
		{3, 4},
		{6, 5, 7},
		{4, 1, 8, 3},
	}
	b := [][]int{
		{-10},
	}
	t.Logf("result:%v", src.ResultOfminimumTotal(a))
	t.Logf("result:%v", src.ResultOfminimumTotal(b))
}
