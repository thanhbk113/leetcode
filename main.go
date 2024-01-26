package main

import "fmt"

func dailyTemperatures(temperatures []int) []int {
	size := len(temperatures)
	ans := make([]int, size)
	stack := make([]int, 0) // Sử dụng stack để lưu trữ index của các ngày chưa tìm được kết quả

	for i := 0; i < size; i++ {
		for len(stack) > 0 && temperatures[i] > temperatures[stack[len(stack)-1]] {
			fmt.Println("~ temperatures[stack[len(stack)-1]] : ", temperatures[stack[len(stack)-1]])
			fmt.Println("~  temperatures[i] : ", temperatures[i])
			fmt.Println("~ stack before: ", stack)

			idx := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			ans[idx] = i - idx
			fmt.Println("~ ans : ", ans)
		}
		stack = append(stack, i)
		fmt.Println("~ stack after: ", stack)
	}

	return ans
}

func main() {

	temperatures := []int{73, 74, 75, 71, 69, 72, 76, 73}
	fmt.Println(dailyTemperatures(temperatures))
}
