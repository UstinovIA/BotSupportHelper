package main

import (
	"botHelper/megafon/promo"
	"fmt"
)

func main() {
	/* 	pathFile := "<путь_к_файлу>"
	   	var result1, result2 = itp.GetCmd(pathFile)
	   	fmt.Println(result1)
	   	fmt.Println(result2) */
	pathFile := "<путь_к_файлу>"
	listPromo := promo.GetCmd(pathFile)
	for i := 0; i < len(listPromo); i++ {
		fmt.Println(listPromo[i])
	}
}
