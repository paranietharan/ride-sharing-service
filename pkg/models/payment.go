package models

var paymentTokens = map[string]float64{
	"token123": 74.00, // Example token with associated fare amount
	"token456": 87.5,
	"token789": 100.00,
}

func GetHardcodedPaymentToken(token string) (float64, bool) {
	amount, exists := paymentTokens[token]
	return amount, exists
}
