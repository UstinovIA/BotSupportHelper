package main

import (
	itp "botHelper/Megafon/ITP"
	"fmt"
)

func main() {
	pathFile := "D:/Projects/ItoolabsBotSupportHelper/megafon/itp/ИТП пример 2.xlsx"
	fmt.Println(itp.GenerateITP(pathFile))
}
