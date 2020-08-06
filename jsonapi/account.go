package jsonapi

import (
	"context"
	"crypto/rand"
	"fmt"
	"net/http"
	"strings"

	"log"
)

// AccountClassification type
type AccountClassification string

// Status of account type
type Status string

// List of constants
const (
	DataVersion     int    = 0
	AccountDataType string = "accounts"

	AccountClassificationPersonal AccountClassification = "Personal"
	AccountClassificationBusiness AccountClassification = "Business"

	PendingStatus   Status = "Pending"
	ConfirmedStatus Status = "Confirmed"
	FailedStatus    Status = "Failed"
)

// AccountList represents list of accounts account from Form3 Fake API
// TODO: It also consists of the pagination details
type AccountList struct {
	Data []Data `json:"data"`
}

// Account represents an account from Form3 Fake API
//
// An Account represents a bank account that is registered with Form3
type Account struct {
	Data Data `json:"data"`
}

// Data struct represents one Account entity detail
type Data struct {
	Type           string     `json:"type"`
	ID             string     `json:"id"`
	OrganisationID string     `json:"organisation_id"`
	Version        int        `json:"version"`
	Attributes     Attributes `json:"attributes"`
}

// Attributes struct represents attributes from an account
type Attributes struct {
	Country                 string                `json:"country"`
	BaseCurrency            string                `json:"base_currency"`
	AccountNumber           string                `json:"account_number"`
	BankID                  string                `json:"bank_id"`
	BankIDCode              string                `json:"bank_id_code"`
	Bic                     string                `json:"bic"`
	Iban                    string                `json:"iban"`
	Name                    []string              `json:"name,omitempty"`
	AlternativeNames        []string              `json:"alternative_names,omitempty"`
	AccountClassification   AccountClassification `json:"account_classification"`
	JointAccount            bool                  `json:"joint_account,omitempty"`
	AccountMatchingOptOut   bool                  `json:"account_matching_opt_out,omitempty"`
	SecondaryIdentification string                `json:"secondary_identification,omitempty"`
	Switched                bool                  `json:"switched"`
	Status                  Status                `json:"status"`
}

// AccountsService conducts the conversation with account related
// resources of Form3
type AccountsService struct {
	client *Client
}

// AccountCreateRequest represents a request to create a Account.
type AccountCreateRequest struct {
	CountryCode           string
	BaseCurrency          string
	BankID                string
	BankIDCode            string
	Bic                   string
	AccountNumber         string
	Iban                  string
	AccountClassification string
}

// Create register an existing bank account with Form3 or create a new one
func (as *AccountsService) Create(ctx context.Context, acr *AccountCreateRequest) (*Account, error) {

	// create the resource
	u := fmt.Sprintf("v1/organisation/accounts")

	id, err := generateUUID()
	if err != nil {
		log.Fatalf("error generating uuid %+v", err)
		return nil, fmt.Errorf("account creation request")
	}

	var ac AccountClassification
	if acr.AccountClassification == "" || strings.EqualFold(acr.AccountClassification, "personal") {
		ac = AccountClassificationPersonal
	}
	if strings.EqualFold(acr.AccountClassification, "business") {
		ac = AccountClassificationBusiness
	}

	a := &Account{
		Data: Data{
			Type:           AccountDataType,
			Version:        DataVersion,
			ID:             id,
			OrganisationID: id,
			Attributes: Attributes{
				Country:               acr.CountryCode,
				BaseCurrency:          acr.BaseCurrency,
				BankID:                acr.BankID,
				BankIDCode:            acr.BankIDCode,
				Bic:                   acr.Bic,
				AccountNumber:         acr.AccountNumber,
				AccountClassification: ac,
			},
		},
	}
	// create a http request
	req, err := as.client.NewRequest(http.MethodPost, u, a)
	if err != nil {
		return nil, err
	}

	resp, err := as.client.Perform(ctx, req, a)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusCreated {
		return nil, &Error{Type: HttpErrorType, Code: resp.StatusCode, Message: "Status Code for Create response must be 201"}
	}

	return a, nil
}

// Fetch retrives the Form3 account information given an id
func (as *AccountsService) Fetch(ctx context.Context, id string) (*Account, error) {
	if id == "" {
		return nil, &Error{Type: UserErrorType, Message: "id can't be blank or empty"}
	}

	// create the resource
	u := fmt.Sprintf("v1/organisation/accounts/%s", id)

	// create a http request
	req, err := as.client.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}

	ac := new(Account)
	_, err = as.client.Perform(ctx, req, ac)
	if err != nil {
		return nil, err
	}

	return ac, nil
}

// List retrives all the Form3 accounts
// TODO: add ListOptions kind of thinhg with page and limit etc
// TODO: handle response here and get rid of response
func (as *AccountsService) List(ctx context.Context) (*AccountList, error) {

	// create the resource
	u := fmt.Sprintf("v1/organisation/accounts")

	// create a http request
	req, err := as.client.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}

	var acs *AccountList

	_, err = as.client.Perform(ctx, req, &acs)
	if err != nil {
		return nil, err
	}

	return acs, nil
}

// Delete removes the Form3 account information given an id and version of data
func (as *AccountsService) Delete(ctx context.Context, id string, version int) (*Account, error) {
	if id == "" {
		return nil, &Error{Type: UserErrorType, Message: "id can't be blank or empty"}
	}

	// create the resource
	u := fmt.Sprintf("v1/organisation/accounts/%s?version=%d", id, version)

	// create a http request
	req, err := as.client.NewRequest(http.MethodDelete, u, nil)
	if err != nil {
		return nil, err
	}

	ac := new(Account)
	_, err = as.client.Perform(ctx, req, ac)
	if err != nil {
		return nil, err
	}

	return ac, nil
}

// generate an UUID v4 using crypto/rand
func generateUUID() (string, error) {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	uuid := fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])

	return uuid, nil
}
