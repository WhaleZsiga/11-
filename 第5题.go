package main

import "fmt"

func main() {
	// 创建一个切片，元素为1到50
	numbers := make([]int, 50)
	for i := range numbers {
		numbers[i] = i + 1
	}

	// 创建一个新的切片来存储非3倍数的元素
	var result []int
	for _, num := range numbers {
		if num%3 != 0 {
			result = append(result, num)
		}
	}

	// 在结果切片的末尾添加数字114514
	result = append(result, 114514)

	// 输出结果切片
	fmt.Println(result)
}
