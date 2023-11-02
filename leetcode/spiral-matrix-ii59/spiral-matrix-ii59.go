package main

import "fmt"

func main() {
	n := 3
	res := generateMatrix(n)
	for _, re := range res {
		for _, i2 := range re {
			fmt.Printf("%d", i2)
		}
		fmt.Println()
	}
}
func generateMatrix(n int) [][]int {
	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, n)
	}
	left, right := 0, n-1 //left 为左边界，right为右边界
	top, bottom := 0, n-1 //top 为上边界，bottom为下边界
	count := 1            //count为计数
	for count <= n*n {
		//最上边 从左到右
		for i := left; i <= right; i++ {
			res[top][i] = count
			count++
		}
		top++

		//最右边 从上到下
		for i := top; i <= bottom; i++ {
			res[i][right] = count
			count++
		}
		right--

		//最下边 从右到左
		for i := right; i >= left; i-- {
			res[bottom][i] = count
			count++
		}
		bottom--

		//最左边 从下到上
		for i := bottom; i >= top; i-- {
			res[i][left] = count
			count++
		}
		left++

	}

	return res
}
