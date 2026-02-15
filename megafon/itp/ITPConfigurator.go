package itp

import (
	"strings"
)

func GenerateITP(path string) string {
	var result strings.Builder
	extensionsConditions := getConditions(path)
	if len(extensionsConditions) == 0 {
		return "Проверьте наличие файла по указанному пути!"
	}
	result.WriteString("bossi product create ")
	result.WriteString(getBasePriceAndCountEmployees(extensionsConditions))
	result.WriteString(getConditionsExtensions(extensionsConditions))
	result.WriteString(getConditionsExtensionsQuant(extensionsConditions))
	resultString := result.String()
	resultString = strings.ReplaceAll(resultString, ",", "") //убираем лишние символы
	resultString = strings.ReplaceAll(resultString, ".00", "")
	return resultString
}

func getBasePriceAndCountEmployees(extensionsConditions [][]string) string {
	employee := getEmployees()
	for i := 0; i < len(extensionsConditions); i++ {
		if extensionsConditions[i][0] == employee[0] {
			var result strings.Builder
			result.WriteString(extensionsConditions[i][3]) //вытаскиваем цену за ТП
			result.WriteString(" ")
			result.WriteString(extensionsConditions[i][1]) // вытаскиваем кол-во сотрудников в ТП
			result.WriteString(" ")
			return result.String()
		}
	}
	return "ERROR!"
}

func getConditionsExtensions(extensionsConditions [][]string) string {
	var result strings.Builder
	listExtensions := getListExtensions()
	for i := 0; i < len(listExtensions); i++ {
		for j := 0; j < len(extensionsConditions); j++ {
			if listExtensions[i][0] == extensionsConditions[j][0] { //формируем строку для опции
				result.WriteString("--")
				result.WriteString(listExtensions[i][1])
				result.WriteString("=")
				result.WriteString(extensionsConditions[j][3])
				result.WriteString(":")
				if strings.Contains(extensionsConditions[j][2], "Можно отключить") {
					result.WriteString("false")
				} else if strings.Contains(extensionsConditions[j][2], "Нельзя отключить") {
					result.WriteString("true")
				}
				result.WriteString(" ")
			}
		}
	}
	return result.String()
}

func getConditionsExtensionsQuant(extensionsConditions [][]string) string {
	var result strings.Builder
	listExtensions := getListExtensionsQuant()
	for i := 0; i < len(listExtensions); i++ {
		for j := 0; j < len(extensionsConditions); j++ {
			if listExtensions[i][0] == extensionsConditions[j][0] { //формируем строку для опции
				result.WriteString("--")
				result.WriteString(listExtensions[i][1])
				result.WriteString("=")
				result.WriteString(extensionsConditions[j][3])
				result.WriteString(":")
				if extensionsConditions[j][3] != "0.00" { //если цена 0, то ограничиваем макс. кол-во опции
					result.WriteString(listExtensions[i][2])
				} else {
					result.WriteString(extensionsConditions[j][1])
				}
				result.WriteString(":")
				if strings.Contains(extensionsConditions[j][2], "Можно отключить") {
					result.WriteString("false")
				} else if strings.Contains(extensionsConditions[j][2], "Нельзя отключить") {
					result.WriteString("true")
				}
				result.WriteString(" ")
			}
		}
	}
	return result.String()
}

//func getEnableExtension()
