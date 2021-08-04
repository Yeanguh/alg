package main

import (
	"fmt"
	"sort"
)

func main() {

	fmt.Println(permute([]int{1,2,3}))
	fmt.Println(permute1([]int{1,1,2,3}))
	fmt.Println(permuteSet([]int{1,2,3}))
	fmt.Println(maxSubArray([]int{5,4,-1,7,8}))
}


func permute(nums []int) [][]int {
	var res [][]int
	var path []int
	used := make(map[int]bool)
	var dfs func([]int)
	dfs = func(path []int) {
		if len(path) == len(nums) {
			tmp := make([]int,len(path))
			copy(tmp,path)
			res = append(res, tmp)
			return
		}
		for _,v := range nums {
			if used[v] {
				continue
			}
			used[v] = true
			path = append(path,v)
			dfs(path)
			path = path[:len(path)-1]
			used[v] = false
		}
	}
	dfs(path)
	return res
}

func permute1(nums []int) [][]int {
	var res [][]int
	used := make(map[int]bool)
	var path []int
	var dfs func([]int)
	dfs = func(path []int) {
		if len(path) == len(nums){
			tmp := make([]int,len(path))
			copy(tmp,path)
			res = append(res,tmp)
			return
		}
		for i,v := range nums{
			if i>0 && used[i-1] && nums[i] == nums[i-1] {
				continue
			}
			if used[i] {
				continue
			}
			used[i] = true
			path = append(path,v)
			dfs(path)
			path = path[:len(path)-1]
			used[i] = false
		}
	}
	dfs(path)
	return res
}

func permuteSet(nums []int) [][]int {

	var res [][]int
	var path []int
	var dfs func([]int)
	sort.Ints(nums)

	dfs = func(path []int) {
		tmp := make([]int,len(path))
		copy(tmp,path)
		res = append(res,tmp)
		for _,v := range nums{
			path = append(path,v)
			dfs(path)
			path = path[:len(path)-1]
		}
	}
	dfs(path)
	return res
}

func maxSubArray(nums []int) int {
	var res int
	for i:=0;i<len(nums);i++{
		if nums[i]>0 {
			res +=nums[i]
		}
		fmt.Println(res)
	}
	return res
}

func max(a,b int) int {
	if a>b {
		return a
	} else{
		return b
	}
}
