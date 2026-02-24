package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// 引数チェック（渡せる文字列は一つのみ）
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run lowercase.go <text>")
		return
	}

	// 文字列の取得
	input := os.Args[1]

	// 小文字化
	output := strings.ToLower(input)

	// 出力
	fmt.Println(output)
}
