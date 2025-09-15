package models

type AuthRequest struct {
	PartnerID string `json:"partner_id" binding:"required"`
	Timestamp string `json:"timestamp" binding:"required"`
	Signature string `json:"signature" binding:"required"`
	Body      string `json:"body" binding:"required"`
}

type AuthResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
}
