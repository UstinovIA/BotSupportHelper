package itp

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
	table, err := file.GetRows("Калькулятор")
	if err != nil {
		fmt.Println(err)
		return [][]string{}
	}
	return table
}

func getConditions(path string) [][]string {
	var calcITP [][]string
	table := openFile(path)
	if len(table) == 0 {
		return [][]string{}
	}
	for i := 0; i < len(table); i++ {
		for j := 0; j < len(table[i]); j++ {
			if table[i][j] == "Абонентские платежи" {
				for k := i + 1; k < len(table); k++ {
					if table[k][j] == "" {
						return calcITP
					}
					extensionConditions := []string{table[k][j], table[k][j+1], table[k][j+2], table[k][j+3]}
					calcITP = append(calcITP, extensionConditions)
				}
			}
		}
	}
	return calcITP
}
