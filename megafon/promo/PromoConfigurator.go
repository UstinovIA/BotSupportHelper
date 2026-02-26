package promo

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// Функция для генерации списка команд на создание промо. На вход принимает путь к файлу, на выходе - список команд в массиве.
func GetCmd(path string) []string {
	var listCmd []string
	table := openFile(path)
	if len(table) == 0 {
		return []string{"Файл не найден!"}
	}
	for i := 1; i < len(table); i++ {
		if strings.Contains(table[i][0], "демо") || strings.Contains(table[i][1], "демо") ||
			strings.Contains(table[i][0], "тестов") || strings.Contains(table[i][1], "тестов") {
			listCmd = append(listCmd, generatePromoForDemo(table[i]))
		} else if strings.Contains(table[i][0], "опци") || strings.Contains(table[i][1], "опци") {
			listCmd = append(listCmd, generatePromoForExtension(table[i]))
		} else {
			listCmd = append(listCmd, fmt.Sprintf("Я не смог обработать строку %d, извини. Покажи это разработчику", i+1))
		}
	}
	return listCmd
}

// генерация промо для продления демо
func generatePromoForDemo(row []string) string {
	var result strings.Builder
	result.WriteString("bossi promo create ")
	date := time.Now()
	dateStr := date.Format("2006-01-02")
	countDay := returnCountDayDemo(row[0], row[1])
	result.WriteString(fmt.Sprintf("promo_%s_demo_%sd ", dateStr, countDay))
	formatedNameCompany := strings.ReplaceAll(row[0], "«", `\"`)
	formatedNameCompany = strings.ReplaceAll(formatedNameCompany, "»", `\"`)
	result.WriteString(fmt.Sprintf(`"%s"`, formatedNameCompany))
	result.WriteString(fmt.Sprintf(" %s ", normalizeDate(row[2])))
	result.WriteString(fmt.Sprintf("%s ", row[3]))
	result.WriteString("prolongation ")
	result.WriteString(fmt.Sprintf("%s ", countDay))
	if row[4] != "" {
		result.WriteString(fmt.Sprintf("%s ", row[4]))
	}
	result.WriteString(getTarget(row[5]))
	if strings.ToLower(row[6]) == "да" {
		result.WriteString(fmt.Sprintf(" %s", "--offline"))
	}
	return result.String()
}

func returnCountDayDemo(nameCompany string, descriptionCompany string) string {
	re := regexp.MustCompile(`(\d+)\s*(день|дня|дней|месяц|месяца|месяцев)`)
	match := re.FindStringSubmatch(nameCompany)
	if match == nil {
		match = re.FindStringSubmatch(descriptionCompany)
	}
	switch match[2] {
	case "месяц", "месяца", "месяцев":
		countMonth, _ := strconv.Atoi(match[1])
		countMonthStr := strconv.Itoa(countMonth * 30)
		return countMonthStr
	case "день", "дня", "дней":
		countDaysStr := match[1]
		return countDaysStr
	default:
		return "!ERROR!"
	}
}

func getTarget(row string) string {
	row = strings.ToLower(row)
	result := "--target="
	if strings.Contains(row, "всех") {
		result = result + "commercial,demo,blocked,freeze"
		return result
	}
	if strings.Contains(row, "действующих") {
		result = result + "commercial,"
	}
	if strings.Contains(row, "новых") {
		result = result + "demo,"
	}
	if strings.Contains(row, "демо") {
		result = result + "demo,"
	}
	result = strings.TrimSuffix(result, ",")
	return result
}

func normalizeDate(intputDate string) string {
	layouts := []string{
		"2006-01-02",
		"02.01.2006",
		"01/02/2006",
	}
	for _, layout := range layouts {
		t, err := time.Parse(layout, intputDate)
		if err == nil {
			return t.Format("2006-01-02")
		}
	}
	return "!ERROR!"
}

// генерация промо на скидку для опций
func generatePromoForExtension(row []string) string {
	var result strings.Builder
	result.WriteString("bossi promo create ")
	date := time.Now()
	dateStr := date.Format("2006-01-02")
	idCompany := fmt.Sprintf("promo_%s_", dateStr)
	currentExtensions, currentExtensionsQuant := getListExtensionsPromo(row[0], row[1])
	for i := 0; i < len(currentExtensions); i++ {
		idCompany = idCompany + currentExtensions[i] + "_"
	}
	for i := 0; i < len(currentExtensionsQuant); i++ {
		idCompany = idCompany + currentExtensionsQuant[i][0] + "_"
	}
	period := getPeriodLenghtForExtensions(row[0], row[1])
	idCompany = idCompany + period + "m"
	result.WriteString(idCompany)
	result.WriteString(" ")
	formatedNameCompany := strings.ReplaceAll(row[0], "«", `\"`)
	formatedNameCompany = strings.ReplaceAll(formatedNameCompany, "»", `\"`)
	result.WriteString(fmt.Sprintf(`"%s"`, formatedNameCompany))
	result.WriteString(fmt.Sprintf(" %s ", normalizeDate(row[2])))
	result.WriteString(fmt.Sprintf("%s ", row[3]))
	result.WriteString("services ")
	result.WriteString(period)
	result.WriteString(" ")
	if row[4] != "" {
		result.WriteString(fmt.Sprintf("%s ", row[4]))
	}
	discount := getDiscount(row[0], row[1])
	for i := 0; i < len(currentExtensions); i++ {
		result.WriteString(fmt.Sprintf("--%s=%s:false ", currentExtensions[i], discount))
	}
	for i := 0; i < len(currentExtensionsQuant); i++ {
		result.WriteString(fmt.Sprintf("--%s=%s:false:%s ", currentExtensionsQuant[i][0], discount, currentExtensionsQuant[i][1]))
	}
	result.WriteString(getTarget(row[5]))
	if strings.ToLower(row[6]) == "да" {
		result.WriteString(fmt.Sprintf(" %s", "--offline"))
	}
	return result.String()
}

// функция для анализа списка опций
func getListExtensionsPromo(nameCompany string, descriptionCompany string) ([]string, [][]string) {
	listExtensions := getListExtensions()
	listExtensionsQuant := getListExtensionsQuant()
	var resultExtensions []string
	var resultExtensionsQuant [][]string
	for i := 0; i < len(listExtensions); i++ {
		if strings.Contains(strings.ToLower(nameCompany), listExtensions[i][0]) {
			resultExtensions = append(resultExtensions, listExtensions[i][1])
		}
	}
	for i := 0; i < len(listExtensionsQuant); i++ {
		if strings.Contains(strings.ToLower(nameCompany), listExtensionsQuant[i][0]) {
			resultExtensionsQuant = append(resultExtensionsQuant, []string{listExtensionsQuant[i][1], listExtensionsQuant[i][2]})
		}
	}
	if resultExtensions == nil && resultExtensionsQuant == nil {
		for i := 0; i < len(listExtensions); i++ {
			if strings.Contains(strings.ToLower(descriptionCompany), listExtensions[i][0]) {
				resultExtensions = append(resultExtensions, listExtensions[i][1])
			}
		}
		for i := 0; i < len(listExtensionsQuant); i++ {
			if strings.Contains(strings.ToLower(descriptionCompany), listExtensionsQuant[i][0]) {
				resultExtensionsQuant = append(resultExtensionsQuant, []string{listExtensionsQuant[i][1], listExtensionsQuant[i][2]})
			}
		}
	}
	return resultExtensions, resultExtensionsQuant
}

func getPeriodLenghtForExtensions(nameCompany string, descriptionCompany string) string {
	re := regexp.MustCompile(`(\d+)\s*(?:месяц|месяца|месяцев)`)
	period := re.FindStringSubmatch(nameCompany)
	if period == nil {
		period = re.FindStringSubmatch(descriptionCompany)
	}
	if period != nil {
		return period[1]
	}
	return "!ERROR!"
}

func getDiscount(nameCompany string, descriptionCompany string) string {
	re := regexp.MustCompile(`(\d+)%`)
	discount := re.FindStringSubmatch(nameCompany)
	if discount == nil {
		discount = re.FindStringSubmatch(descriptionCompany)
	}
	if discount != nil {
		return discount[1]
	}
	return "100"
}
