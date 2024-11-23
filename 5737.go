package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 判断是否为闰年
func isLeapYear(year int) bool {
	if year%400 == 0 {
		return true
	} else if year%100 == 0 {
		return false
	} else if year%4 == 0 {
		return true
	}
	return false
}

func main() {
	// 读取输入
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	// 分割输入并转换为整数
	parts := strings.Split(input, " ")
	x, _ := strconv.Atoi(parts[0])
	y, _ := strconv.Atoi(parts[1])

	// 初始化变量存储闰年
	leapYears := []int{}

	// 遍历区间内的年份并检查是否为闰年
	for i := x; i <= y; i++ {
		if isLeapYear(i) {
			leapYears = append(leapYears, i)
		}
	}

	// 将整数切片转换为字符串切片
	leapYearsStr := make([]string, len(leapYears))
	for i, year := range leapYears {
		leapYearsStr[i] = strconv.Itoa(year)
	}

	// 输出结果
	fmt.Println(len(leapYears))
	fmt.Println(strings.Join(leapYearsStr, " "))
}
