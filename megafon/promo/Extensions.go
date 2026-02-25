package promo

// перечисление опций (не количественных)
func getListExtensions() [][]string {
	extensioins := [][]string{
		{"Аналитика", "analytics"},
		{"Эффективные продажи", "effectivesale"},
		{"Эффективное обслуживание", "effectiveservice"},
		{"Интеграция с CRM", "crm"},
		{"Автоинформирование", "autocaller"},
		{"Разговоры на повышенных тонах", "emotion"},
		{"Большой бизнес", "bigbusiness"},
	}
	return extensioins
}

// перечисление опций (количественных) с указанием на сколько ед. применяется промокод
func getListExtensionsQuant() [][]string {
	extensioinsQuant := [][]string{
		{"Виджет обратного звонка", "callback", "1"},
		{"Запись и хранение разговоров", "callsrecord", "1"},
		{"Запись разговоров", "callsrecord", "1"},
	}
	return extensioinsQuant
}
