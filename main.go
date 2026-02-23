package main

import (
	itp "botHelper/Megafon/ITP"
	"fmt"
)

func main() {
	pathFile := "<путь_к_файлу>"
	var result1, result2 = itp.GetCmd(pathFile)
	fmt.Println(result1)
	fmt.Println(result2)
}
