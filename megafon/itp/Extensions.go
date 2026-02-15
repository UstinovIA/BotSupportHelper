package itp

func getEmployees() []string {
	employee := []string{
		"Количество сотрудников в пакете",
		"Количество доп. сотрудников",
	}
	return employee
}

// перечисление опций (не количественных)
func getListExtensions() [][]string {
	extensioins := [][]string{
		{"Аналитика", "analytics"},
		{"Эффективные продажи", "effectivesale"},
		{"Эффективное обслуживание", "effectiveservice"},
		{"Интеграция (API или прямая с CRM)", "crm"},
		{"Автоинформирование", "autocaller"},
		{"Раговоры на повышенных тонах", "emotion"},
		{"Большой бизнес", "bigbusiness"},
	}
	return extensioins
}

// перечисление опций (количественных) с указанием дефолтного макс. значения
func getListExtensionsQuant() [][]string {
	extensioinsQuant := [][]string{
		{"Количество доп. сотрудников", "addseat", "9999"},
		{"Виджет, шт", "callback", "50"},
		{"Запись разговоров, лет", "callsrecord", "3"},
	}
	return extensioinsQuant
}
