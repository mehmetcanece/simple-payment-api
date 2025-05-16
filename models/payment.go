package models

type PaymentRequest struct {
	Amount float64 `json:"amount"` //jsondan gelen alanın bunla eşleşmesi için bu sondakini yazdık
	Method string  `json:"method"`
}