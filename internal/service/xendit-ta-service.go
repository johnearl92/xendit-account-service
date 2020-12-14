// Package service Official-Receipt API.
package service

import (
	"github.com/johnearl92/xendit-account-service/internal/model"
	"github.com/johnearl92/xendit-account-service/internal/store"
)

// xenditService EOR service definition
type xenditService struct {
	acctStore store.AccountStore
}

// NewXenditService xenditService provider
func NewXenditService(acctStore store.AccountStore) XenditService {
	return &xenditService{
		acctStore: acctStore,
	}
}

// XenditService Receiptservice interface
type XenditService interface {
	CreateAccount(account *model.Account, opts store.GetOpts) error
	UpdateAccount(account *model.Account) error
	GetAccount(id string, opts store.GetOpts) (*model.Account, error)
}

// Create saves receipts in database
func (p *xenditService) CreateAccount(acoount *model.Account, opts store.GetOpts) error {
	return p.acctStore.Create(acoount, opts)
}

// Update updates the specific receipt
func (p *xenditService) UpdateAccount(account *model.Account) error {
	return p.acctStore.Update(account)
}

// Get find a specific receipt in database
func (p *xenditService) GetAccount(id string, opts store.GetOpts) (*model.Account, error) {
	return p.acctStore.Get(id, opts)
}
