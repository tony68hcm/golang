package model

type Response struct {
	Status        string `json:"status"`
	ReasonCode    uint16 `json:"reason_code"`
	ReasonMessage string `json:"reason_message"`
}
