package main

import (
	"fmt"
	"math"
)

// 判断一个数是否为质数
func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	// 只需检查到平方根即可
	sqrtN := int(math.Sqrt(float64(n)))
	for i := 2; i <= sqrtN; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	var n int
	// 从标准输入读取一个数
	fmt.Scan(&n)

	// 判断是否为质数并输出结果
	if isPrime(n) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
