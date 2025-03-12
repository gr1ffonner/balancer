package models

type RedirectRequest struct {
	Video string `json:"video"`
}

type RedirectResponse struct {
	RedirectURL string `json:"redirect_url"`
}
