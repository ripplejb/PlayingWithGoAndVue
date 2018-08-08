package models

type Dividends struct {
	ExDate       string  `json:"exDate"`
	PaymentDate  string  `json:"paymentDate"`
	Amount       float64 `json:"amount"`
	Indicated    string  `json:"indicated"`
	Flag         string  `json:"flag"`
	RecordDate   string  `json:"recordDate"`
	DeclaredDate string  `json:"declaredDate"`
	Type         string  `json:"type"`
	Qualified    string  `json:"qualified"`
}
