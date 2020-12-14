// Package handler Official-Receipt API.
//
// Official-Receipt service API endpoints
//
//     Version: 1.0.0
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
// swagger:meta
package handler

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/johnearl92/xendit-account-service/internal/model"
	"github.com/johnearl92/xendit-account-service/internal/service"
	"github.com/johnearl92/xendit-account-service/internal/utils"
	log "github.com/sirupsen/logrus"
	"net/http"
)

//XenditHandler main handler for the eor
type XenditHandler struct {
	xenditService service.XenditService
}

// NewORHandler provides XenditHandler definition
func NewXenditHandler(p service.XenditService) *XenditHandler {
	return &XenditHandler{
		xenditService: p,
	}
}

// GetHeartBeat check if the server is up
func (h *XenditHandler) GetHeartBeat(res http.ResponseWriter, req *http.Request) {
	log.Debugln("invoke getHeartBeat")
	resp := &model.GenericResponse{
		Success: true,
	}

	utils.WriteEntity(res, http.StatusOK, resp)
	log.Debugln("end getHeartBeat")
}

func (h *XenditHandler) AddAccount(res http.ResponseWriter, req *http.Request) {
	log.Debugln("invoke AddAccount")
	var data model.AccountRequest
	log.Infoln("Reading Request")
	if err := utils.ReadEntity(req, &data); err != nil {
		log.Errorln(err)
		utils.WriteError(res, http.StatusBadRequest, err)
		return
	}
	var account = &model.Account{
		Name:         data.Name,
		CallbackUrl:  data.CallbackUrl,
		MerchantCode: data.MerchantCode,
	}

	if err := h.xenditService.CreateAccount(account, nil); err != nil {
		log.Error(err.Error())
		utils.WriteServerError(res, "/account", "Failed to save account",
			fmt.Sprintf("Failed to save account. Please contact the administrator. Error: %s", err.Error()))
		return
	}

	resp := &model.AccountResponse{
		MerchantCode: account.MerchantCode,
		MerchantKey:  account.MerchantKey,
	}

	log.Debugln("end addComment")
	utils.WriteEntity(res, http.StatusOK, resp)

}

func (h *XenditHandler) GetAccount(res http.ResponseWriter, req *http.Request) {
	log.Debugln("invoke GetAccount")
	vars := mux.Vars(req)

	account, err := h.xenditService.GetAccount(vars["id"], nil)

	if err != nil {
		log.Error(err.Error())
		utils.WriteServerError(res, "/accounts/{id}", "Failed to get Account",
			fmt.Sprintf("Failed to get Account. Please check organization name or contact the administrator. Error: %s", err.Error()))
		return
	}

	accountResponse := &model.AccountResponse{
		MerchantKey:  account.MerchantKey,
		MerchantCode: account.MerchantCode,
		Name:         account.Name,
		CallbackUrl:  account.CallbackUrl,
	}

	log.Debugln("end getComments")
	utils.WriteEntity(res, http.StatusOK, accountResponse)

}

func (h *XenditHandler) UpdateAccount(res http.ResponseWriter, req *http.Request) {
	log.Debugln("invoke UpdateAccount")
	var data model.AccountRequest

	log.Infoln("Reading Request")
	if err := utils.ReadEntity(req, &data); err != nil {
		log.Errorln(err)
		utils.WriteError(res, http.StatusBadRequest, err)
		return
	}

	vars := mux.Vars(req)

	log.Infoln("Retreiving Account")
	account, err := h.xenditService.GetAccount(vars["id"], nil)

	if err != nil {
		log.Error(err.Error())
		utils.WriteServerError(res, "/accounts/{id}", "Failed to get Account",
			fmt.Sprintf("Failed to get Account. Please check organization name or contact the administrator. Error: %s", err.Error()))
		return
	}

	if len(data.Name) > 0 {
		account.Name = data.Name
	}
	if len(data.CallbackUrl) > 0 {
		account.CallbackUrl = data.CallbackUrl
	}

	if len(data.MerchantCode) > 0 {
		account.MerchantCode = data.MerchantCode
	}

	log.Infoln("Updating Account")
	if err := h.xenditService.UpdateAccount(account); err != nil {
		log.Error(err.Error())
		utils.WriteServerError(res, "/account/{id}", "Failed to update account",
			fmt.Sprintf("Failed to update account. Please contact the administrator. Error: %s", err.Error()))
		return
	}

	resp := &model.GenericResponse{
		Success: true,
	}

	log.Debugln("end Update Account")
	utils.WriteEntity(res, http.StatusOK, resp)

}

// Register registers the route
func (h *XenditHandler) Register(router *mux.Router) {

	// swagger:operation GET  /health GenericRes
	// ---
	// summary: This will check if the server is up
	// responses:
	//   "200":
	//     "$ref": "#/responses/GenericRes"
	//   "400":
	//     "$ref": "#/responses/JSONErrors"
	//   "500":
	//     "$ref": "#/responses/JSONErrors"
	router.HandleFunc("/health", h.GetHeartBeat).Methods(http.MethodGet)
	log.Info("[GET] /health registered")

	// swagger:operation POST /accounts acct AccountReq
	// Add account
	// ---
	// responses:
	//   "200":
	//     "$ref": "#/responses/AccountRes"
	//   "400":
	//     "$ref": "#/responses/JSONErrors"
	//   "500":
	//     "$ref": "#/responses/JSONErrors"
	router.HandleFunc("/accounts", h.AddAccount).Methods(http.MethodPost)
	log.Info("[POST] /accouunts registered")

	// swagger:operation GET  /accounts/{id} acct
	// ---
	// summary: This will get the account via id
	// parameters:
	// - name: id
	//   in: path
	//   required: true
	//   type: string
	// responses:
	//   "200":
	//     "$ref": "#/responses/AccountRes"
	//   "400":
	//     "$ref": "#/responses/JSONErrors"
	//   "500":
	//     "$ref": "#/responses/JSONErrors"
	router.HandleFunc("/accounts/{id}", h.GetAccount).Methods(http.MethodGet)

	// swagger:operation PUT /accounts/{id} acct
	// ---
	// summary: This will update the account via id
	// parameters:
	// - name: id
	//   in: path
	//   required: true
	//   type: string
	// - name: AccountReq
	//   in: body
	//   required: true
	//   schema:
	//     $ref: "#/definitions/AccountReq"
	// responses:
	//   "200":
	//     "$ref": "#/responses/AccountRes"
	//   "400":
	//     "$ref": "#/responses/JSONErrors"
	//   "500":
	//     "$ref": "#/responses/JSONErrors"
	router.HandleFunc("/accounts/{id}", h.UpdateAccount).Methods(http.MethodPut)
	log.Info("[PUT] /accouunts/{id} registered")

}
