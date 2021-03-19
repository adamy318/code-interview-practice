package main

import (
	"fmt"
	"sort"
)


func backtrack(index int, arr []int, curr []int, out *[][]int) {
	*out = append(*out, curr)
	for i := index; i < len(arr); i++ {
		tmp2 := append(curr, arr[i])

		backtrack(index + 1, arr, tmp2, out)

		//curr = curr[:len(curr)-1]
	}
}


func findComboSum(index int, arr []int, curr []int, out *[][]int, target int) {
	if target == 0 {
		*out = append(*out, curr)
		return
	}
	if target < 0 {
		return
	}

	tmp2 := []int{}
	for i := index; i < len(arr); i++ {
		if arr[i] > target {
			return
		} else {
			tmp2 = append(curr, arr[i])
			fmt.Println(tmp2)
		}
		findComboSum(index, arr, tmp2, out, target-arr[i])

	}
}

func subsets(nums []int) [][]int {
	out := [][]int{}

	backtrack(0, nums, make([]int, 0), &out)

	return out

}

func comboSum(nums []int, target int) [][]int {
	out := [][]int{}

	findComboSum(0, nums, make([]int, 0), &out, target)

	return out
}


func main() {
	//a := []int{1, 2, 3}
	//a := []int{8, 10, 6, 11, 1, 16, 8}
	a := []int{2,3,6,7}
	b := 7
	c := []int{}

	sort.Ints(a)
	for i := 0; i < len(a)-1; i++ {
		if a[i] == a[i+1] {
			continue
		} else {
			c = append(c, a[i])
		}
	}
	c = append(c, a[len(a)-1])

	fmt.Println(c)

	//fmt.Println(subsets(a))
	fmt.Println(comboSum(c, b))

}
