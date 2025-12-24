package domain

type CountResponse struct {
	Word   string `json:"word"`
	Count  int    `json:"count"`
	Amount string `json:"amount"`
	Error  string `json:"error,omitempty"`
}

type PDFProcessor interface {
	CountWordOccurrences(filePath, word string) (int, float64, error)
}

var InvalidPhrases = map[string]struct{}{
	"В Kaspi Банкомате":              {},
	"С карты другого банка":          {},
	"С Kaspi Депозита":               {},
	"Кредит Наличными":               {},
	"Оплата за проданный автомобиль": {},
}
