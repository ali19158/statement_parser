package domain

type CountResponse struct {
	Count  int    `json:"count"`
	Amount string `json:"amount"`
	Error  string `json:"error,omitempty"`
}

type PDFProcessor interface {
	CountWordOccurrences(filePath string) (int, float64, error)
}

/*Пополнение Толықтыру Replenishment*/
var InvalidPhrases = map[string]struct{}{
	"В Kaspi Банкомате":                                {},
	"С карты другого банка":                            {},
	"С Kaspi Депозита":                                 {},
	"Кредит Наличными":                                 {},
	"Оплата за проданный автомобиль":                   {},
	"Kaspi банкоматында":                               {},
	"Kaspi Gold-ты басқа банктің картасынан толықтыру": {},
	"Kaspi Депозиттен":                                 {},
	"Ақшалай Кредит":                                   {},
	"Сатылған автомобиль үшін ақы төлеу":               {},
	"At Kaspi ATM":             {},
	"From card of other banks": {},
	"From Kaspi Deposit":       {},
	"Сash Loan":                {},
	"Payment for a car sold":   {},
}

var LangWords = map[string]string{
	"ВЫПИСКА":        "Пополнение",
	"ҮЗІНДІ КӨШІРМЕ": "Толықтыру",
	"Kaspi Gold":     "Replenishment",
}
