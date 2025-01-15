package dto

type (
	PurchaseBody struct {
		Duration int `json:"duration"`
	}

	PurchasePayload struct {
		UserId string `json:"user_id"`
		Duration int `json:"duration"`
	}
)