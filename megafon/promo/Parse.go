package promo

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)

func openFile(path string) [][]string {
	file, err := excelize.OpenFile(path)
	if err != nil {
		fmt.Println(err)
		return [][]string{}
	}
	defer file.Close()
	table, err := file.GetRows("Лист1")
	if err != nil {
		fmt.Println(err)
		return [][]string{}
	}
	return table
}
