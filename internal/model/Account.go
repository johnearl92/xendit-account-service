package model

type Account struct {
	BaseModel
	MerchantKey  string
	MerchantCode string
	Name         string
	CallbackUrl  string
}

type AccountRequest struct {
	Name         string `json:"name,omitempty"`
	CallbackUrl  string `json:"callbackUrl,omitempty"`
	MerchantCode string `json:"merchantCode,omitempty"`
}

type AccountResponse struct {
	MerchantKey  string `json:"merchantKey"`
	MerchantCode string `json:"merchantCode"`
	Name         string `json:"name,omitempty"`
	CallbackUrl  string `json:"callbackUrl,omitempty"`
}

// AccountReq parameters model
//
// swagger:parameters AccountReq
type AccountReq struct {
	// in: body
	// required: true
	Account AccountRequest
}

// AccountResponse comments response
type AccountResponseList struct {
	Accounts []Account `json:"accounts"`
}

// AccountResWrapperv wrapper struct for AccountResponse
//
// swagger:response AccountResponse
type AccountResWrapper struct {
	// in: body
	// required: true
	AccountResponseList AccountResponseList
}

// AccountRes AccountResponse
//
// swagger:response AccountRes
type AccountRes struct {
	// in: body
	// required: true
	AccountResponse AccountResponse
}
