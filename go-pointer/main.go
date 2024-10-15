package main

import (
	"flag"
	"fmt"
	"strings"
)

// 创建两个指针
var n = flag.Bool("n", false, "omit trailing newline") // 忽略换行符
var sep = flag.String("s", " ", "separator") // 指定分割符

func main() {
    flag.Parse() // flag.Parse() 会更新指针指向内存的值
    fmt.Print(strings.Join(flag.Args(), *sep)) // 解饮用获取到最新的值
    if !*n {
        fmt.Println()
    }
}

// test
// $ go build main.go
// ./main -s - a bc def
// ./main -n a bc def