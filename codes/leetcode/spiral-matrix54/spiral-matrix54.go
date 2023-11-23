package main

import "fmt"

func main() {
	n, m := 3, 3

	//创建数组
	nums := make([][]int, n)
	for i := 0; i < n; i++ {
		nums[i] = make([]int, m)
	}

	//赋值
	count := 1
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			nums[i][j] = count
			count++
		}
	}

	res := spiralOrder(nums)

	for _, re := range res {
		fmt.Println(re)
	}

}

func spiralOrder(matrix [][]int) []int {
	left, right := 0, len(matrix[0])-1
	top, bottom := 0, len(matrix)-1
	max := (right + 1) * (bottom + 1)
	res := make([]int, 0)

	x, y := 0, 0 //移动的横纵坐标

	for i := 1; i < max; {

		//最上边  从左到右
		for ; x <= right; x++ {
			res = append(res, matrix[y][x])
			i++
		}
		x--
		top++
		y++

		//最右边  从上到下
		for ; y <= bottom; y++ {
			res = append(res, matrix[y][x])
			i++
		}
		y--
		right--
		x--

		//最下边  从右到左
		for ; x >= left; x-- {
			res = append(res, matrix[y][x])
			i++
		}
		x++
		bottom--
		y--

		//最左边  从下到上
		for ; y >= top; y-- {
			res = append(res, matrix[y][x])
			i++
		}
		y++
		left++
		x++

	}

	return res
}
