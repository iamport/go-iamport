package iamport

import (
	"net/http"

	"github.com/iamport/go-iamport/authenticate"
)

const (
	DefaultURL = "https://api.iamport.kr"

	ErrMustExistImpUID              = "iamport: imp_uid must be exist"
	ErrMustExistMerchantUID         = "iamport: Merchant UID must be exist"
	ErrMustExistImpUIDorMerchantUID = "iamport: imp_uid or Merchant UID must be exist"
	ErrMustExistCustomerUID         = "iamport: customer_uid must be exist"
	ErrInvalidStatusParam           = "iamport: status parmeter is invalid. must be all, ready, paid, failed and cancelled"
	ErrInvalidSortParam             = "iamport: sort parmeter is invalid. must be -started, started, -paid, paid, -updated and updated"
	ErrInvalidPage                  = "iamport: page is more than 1"
	ErrInvalidLimit                 = "iamport: limit is more than 0"
	ErrInvalidFrom                  = "iamport: 'from' cannot be more future than 'to'"
	ErrInvalidTo                    = "iamport: 'to' date cannot be more than 3 months."
	ErrInvalidAmount                = "iamport: amount is more than 0"
)

type Iamport struct {
	Authenticate *authenticate.Authenticate
}

func NewIamport(apiURL string, restAPIKey string, restAPISecret string) (*Iamport, error) {
	client := &http.Client{}

	auth, err := authenticate.NewAuthenticate(apiURL, client, restAPIKey, restAPISecret)
	if err != nil {
		return nil, err
	}

	iamport := &Iamport{
		Authenticate: auth,
	}

	return iamport, nil
}
