package jsonapi

import (
	"context"
	"crypto/rand"
	"fmt"
	"net/http"

	"github.com/prometheus/common/log"
)

// Account represents an account from JsonAPI
//
// An Account represents a bank account that is registered with Form3

type AccountClassification string

type Status string

const (
	AccountResource string = "v1/organisation/accounts"

	DataVersion int    = 0
	DataType    string = "accounts"

	AccountClassificationPersonal AccountClassification = "Personal"
	AccountClassificationBusiness AccountClassification = "Business"

	PendingStatus   Status = "Pending"
	ConfirmedStatus Status = "Confirmed"
	FailedStatus    Status = "Failed"
)

type AccountList struct {
	Data []Data `json:"data"`
}

type Account struct {
	Data Data `json:"data"`
}

type Data struct {
	Type           string     `json:"type"`
	ID             string     `json:"id"`
	OrganisationID string     `json:"organisation_id"`
	Version        int        `json:"version"`
	Attributes     Attributes `json:"attributes"`
}

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
// resources of JsonAPI
type AccountsService struct {
	client *Client
}

// AccountCreateRequest represents a request to create a Droplet.
type AccountCreateRequest struct {
}

func generateUUID() (string, error) {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	uuid := fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])

	return uuid, nil
}

// Create retrives the account information given an id
// TODO: handle response here and get rid of response
func (as *AccountsService) Create(ctx context.Context, acr *AccountCreateRequest) (*Account, *http.Response, error) {

	// create the resource
	u := fmt.Sprintf(AccountResource)

	id, err := generateUUID()
	if err != nil {
		log.Errorf("error generating uuid %+v", err)
		return nil, nil, fmt.Errorf("account creation request")
	}

	ac := &Account{
		Data: Data{
			Type:           DataType,
			Version:        DataVersion,
			ID:             id,
			OrganisationID: id,
			Attributes: Attributes{
				Country:               "HK",
				BankID:                "111111",
				BankIDCode:            "HKNCC",
				Bic:                   "hvh",
				AccountClassification: AccountClassificationPersonal,
			},
		},
	}
	// create a http request
	req, err := as.client.NewRequest(http.MethodPost, u, ac)
	if err != nil {
		return nil, nil, err
	}

	resp, err := as.client.Perform(ctx, req, ac)
	if err != nil {
		return nil, resp, err
	}

	return ac, resp, nil
}

// Fetch retrives the account information given an id
// TODO: handle response here and get rid of response
func (as *AccountsService) Fetch(ctx context.Context, id string) (*Account, *http.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id can't be blank or empty")
	}

	// create the resource
	u := fmt.Sprintf(AccountResource+"/%s", id)

	// create a http request
	req, err := as.client.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	ac := new(Account)
	resp, err := as.client.Perform(ctx, req, ac)
	if err != nil {
		return nil, resp, err
	}

	return ac, resp, nil
}

// List retrives all the accounts
// TODO: add ListOptions kind of thinhg with page and limit etc
// TODO: handle response here and get rid of response
/* func (as *AccountsService) List(ctx context.Context) ([]*Account, *http.Response, error) {

	// create the resource
	u := fmt.Sprintf(AccountResource)

	// create a http request
	req, err := as.client.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, nil, err
	}

	var acs []*Account

	resp, err := as.client.Perform(ctx, req, &acs)
	if err != nil {
		return nil, resp, err
	}

	return acs, resp, nil
} */
