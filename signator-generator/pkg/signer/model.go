package signer

type SignRequest struct {
	PartnerID string `json:"partner_id"`
	Body      string `json:"body"`
}

type SignResponse struct {
	Timestamp string `json:"X-TIMESTAMP"`
	Signature string `json:"X-SIGNATURE"`
}
