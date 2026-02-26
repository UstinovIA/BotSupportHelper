package promo

// перечисление опций (не количественных)
func getListExtensions() [][]string {
	extensioins := [][]string{
		{"аналитика", "analytics"},
		{"эффективные продажи", "effectivesale"},
		{"эффективное обслуживание", "effectiveservice"},
		{"интеграция с crm", "crm"},
		{"автоинформирование", "autocaller"},
		{"разговоры на повышенных тонах", "emotion"},
		{"большой бизнес", "bigbusiness"},
		{"эффектиное обслуживание", "effectiveservice"},
	}
	return extensioins
}

// перечисление опций (количественных) с указанием на сколько ед. применяется промокод
func getListExtensionsQuant() [][]string {
	extensioinsQuant := [][]string{
		{"виджет обратного звонка", "callback", "1"},
		{"запись и хранение разговоров", "callsrecord", "1"},
		{"запись разговоров", "callsrecord", "1"},
	}
	return extensioinsQuant
}
