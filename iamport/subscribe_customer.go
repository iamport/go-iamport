package iamport

import (
	"errors"

	TypeSubscribe "github.com/iamport/interface/gen_src/go/v1/subscribe"
	TypeSubscribeCust "github.com/iamport/interface/gen_src/go/v1/subscribe_customers"

	subscribeCust "github.com/iamport/go-iamport/subscribe_customer"
	"github.com/iamport/go-iamport/util"
)

// GetMultipleBillingKeysByCustomer 여러개의 Customer UID를 통하여 결제내역을 가져오는 API
//
// GET /subscribe/customers
func (iamport *Iamport) GetMultipleBillingKeysByCustomer(customerUIDs []string) ([]*TypeSubscribeCust.CustomerBillingKey, error) {
	if customerUIDs == nil || len(customerUIDs) == 0 {
		return nil, errors.New(ErrMustExistCustomerUID)
	}

	req := &TypeSubscribeCust.GetMultipleCustomerBillingKeyRequest{
		CustomerUid: customerUIDs,
	}

	res, err := subscribeCust.GetMultipleBillingKeysByCustomer(iamport.Authenticate, req)
	if err != nil {
		return nil, err
	}

	if res.Code != util.CodeOK {
		return nil, errors.New(res.Message)
	}

	return res.Response, nil
}

// DeleteBillingKey Customer UID에 해당하는 빌링키 삭제
//
// DELETE /subscribe/customers/{customer_uid}
func (iamport *Iamport) DeleteBillingKey(customerUID, reason, requester string) (*TypeSubscribeCust.CustomerBillingKey, error) {
	if customerUID == "" {
		return nil, errors.New(ErrMustExistCustomerUID)
	}

	req := &TypeSubscribeCust.DeleteCustomerBillingKeyRequest{
		CustomerUid: customerUID,
		Reason:      reason,
		Requester:   requester,
	}

	res, err := subscribeCust.DeleteBillingKey(iamport.Authenticate, req)
	if err != nil {
		return nil, err
	}

	if res.Code != util.CodeOK {
		return nil, errors.New(res.Message)
	}

	return res.Response, nil
}

// GetBillingKeyByCustomer Customer UID에 해당하는 빌링키 정보 불러오기
//
// GET /subscribe/customers/{customer_uid}
func (iamport *Iamport) GetBillingKeyByCustomer(customerUID string) (*TypeSubscribeCust.CustomerBillingKey, error) {
	if customerUID == "" {
		return nil, errors.New(ErrMustExistCustomerUID)
	}

	req := &TypeSubscribeCust.GetCustomerBillingKeyRequest{
		CustomerUid: customerUID,
	}

	res, err := subscribeCust.GetBillingKeyByCustomer(iamport.Authenticate, req)
	if err != nil {
		return nil, err
	}

	if res.Code != util.CodeOK {
		return nil, errors.New(res.Message)
	}

	return res.Response, nil
}

// InsertBillingKeyByCustomer Customer UID에 빌링키 정보 집어넣기
//
// POST /subscribe/customers/{customer_uid}
func (iamport *Iamport) InsertBillingKeyByCustomer(
	customerUID, pg string,
	cardNumber, expiry, birth, pwd2Digit string,
	customerName, customerTel, customerEmail, customerAddr, customerPostcode string,
) (*TypeSubscribeCust.CustomerBillingKey, error) {
	if customerUID == "" {
		return nil, errors.New(ErrMustExistCustomerUID)
	}

	req := &TypeSubscribeCust.InsertCustomerBillingKeyRequest{
		CustomerUid:      customerUID,
		Pg:               pg,
		CardNumber:       cardNumber,
		Expiry:           expiry,
		Birth:            birth,
		Pwd_2Digit:       pwd2Digit,
		CustomerName:     customerName,
		CustomerTel:      customerTel,
		CustomerEmail:    customerEmail,
		CustomerAddr:     customerAddr,
		CustomerPostcode: customerPostcode,
	}

	res, err := subscribeCust.InsertBillingKeyByCustomer(iamport.Authenticate, req)
	if err != nil {
		return nil, err
	}

	if res.Code != util.CodeOK {
		return nil, errors.New(res.Message)
	}

	return res.Response, nil
}

// GetPaymentsByCustomer Customer UID로 결제한 내역 불러오기
//
// GET /subscribe/customers/{customer_uid}/payments
func (iamport *Iamport) GetPaymentsByCustomer(customerUID string, page int32) (*TypeSubscribeCust.NestedGetPaidByBillingKeyListData, error) {
	if customerUID == "" {
		return nil, errors.New(ErrMustExistCustomerUID)
	}

	req := &TypeSubscribeCust.GetPaidByBillingKeyListRequest{
		CustomerUid: customerUID,
		Page:        page,
	}

	res, err := subscribeCust.GetPaymentsByCustomer(iamport.Authenticate, req)
	if err != nil {
		return nil, err
	}

	if res.Code != util.CodeOK {
		return nil, errors.New(res.Message)
	}

	return res.Response, nil
}

// GetScheduledPaymentListByCustomerUID Customer UID로 결제한 내역 불러오기
//
// GET /subscribe/customers/{customer_uid}/schedule
func (iamport *Iamport) GetScheduledPaymentListByCustomerUID(customerUID string,
	page, from, to int32, scheduleStatus string,
) (*TypeSubscribe.NestedGetPaymentScheduleByCustomerData, error) {
	if customerUID == "" {
		return nil, errors.New(ErrMustExistCustomerUID)
	}

	var revisedPage int32 = 1
	if page != 0 {
		revisedPage = page
	}

	req := &TypeSubscribe.GetPaymentScheduleByCustomerRequest{
		CustomerUid:    customerUID,
		Page:           revisedPage,
		From:           from,
		To:             to,
		ScheduleStatus: scheduleStatus,
	}

	res, err := subscribeCust.GetScheduledPaymentByCustomerUID(iamport.Authenticate, req)
	if err != nil {
		return nil, err
	}

	if res.Code != util.CodeOK {
		return nil, errors.New(res.Message)
	}

	return res.Response, nil
}
