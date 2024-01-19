package main

import (
	"fmt"
	"sort"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func minDays(bloomDay []int, m int, k int) int {
	lg := len(bloomDay)
	M := make(map[int]int)
	res := [][]int{}
	if lg < m*k {
		return -1
	}
	for i := 0; i <= lg-k; i++ {
		tmp := 0
		for j := i; j < i+k; j++ {
			tmp = max(tmp, bloomDay[j])
		}
		res = append(res, []int{tmp, i})
	}
	sort.Slice(res, func(i, j int) bool { return res[i][0] < res[j][0] })
	cnt := 0
	for _, nums := range res {
		if M[nums[1]] == 0 {
			cnt++
			if cnt == m {
				return nums[0]
			}
			for i := 0; i < k; i++ {
				M[nums[1]+i]++
			}
		}
	}
	return -1
}

func main() {
	a := []int{56,89,16,25,33,49,6,95,16,8,63,53,38,1,91}
	fmt.Println(minDays(a,5,3))
}
