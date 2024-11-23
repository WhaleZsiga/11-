package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// 读取输入
	scanner := bufio.NewScanner(os.Stdin)

	// 读取第一行数据（10个苹果的高度）
	scanner.Scan()
	heights := strings.Split(scanner.Text(), " ")
	appleHeights := make([]int, 10)
	for i, heightStr := range heights {
		appleHeights[i], _ = strconv.Atoi(heightStr)
	}

	// 读取第二行数据（陶陶的最大伸手高度）
	scanner.Scan()
	taoTaoHeight, _ := strconv.Atoi(scanner.Text())

	// 加上板凳的高度（30厘米）
	taoTaoHeight += 30

	// 计算陶陶能摘到的苹果数量
	count := 0
	for _, height := range appleHeights {
		if height <= taoTaoHeight {
			count++
		}
	}

	// 输出结果
	fmt.Println(count)
}
