package iamport

import (
	"errors"

	Typepayment "github.com/iamport/interface/gen_src/go/v1/payment"
	TypeSubscribe "github.com/iamport/interface/gen_src/go/v1/subscribe"

	"github.com/iamport/go-iamport/subscribe"
	"github.com/iamport/go-iamport/util"
)

// OnetimePayment ActiveX 없는 비인증결제
//
// POST /subscribe/payments/onetime
func (iamport *Iamport) OnetimePayment(
	merchantUID string,
	amount, taxFree int32,
	cardNumber, expiry, birth, pwd_2digit string,
	customerUid, pg, name string,
	buyerName, buyerEmail, buyerTel, buyerAddr, buyerPostcode string,
	cardQuota int32, interestFreeByMerchant bool,
	customData, noticeUrl string,
) (*Typepayment.Payment, error) {

	if merchantUID == "" {
		return nil, errors.New(ErrMustExistMerchantUID)
	}

	if amount < 0 {
		return nil, errors.New(ErrInvalidAmount)
	}

	req := &TypeSubscribe.OnetimePaymentRequest{
		MerchantUid:            merchantUID,
		Amount:                 amount,
		TaxFree:                taxFree,
		CardNumber:             cardNumber,
		Expiry:                 expiry,
		Birth:                  birth,
		Pwd_2Digit:             pwd_2digit,
		CustomerUid:            customerUid,
		Pg:                     pg,
		Name:                   name,
		BuyerName:              buyerName,
		BuyerEmail:             buyerEmail,
		BuyerTel:               buyerTel,
		BuyerAddr:              buyerAddr,
		BuyerPostcode:          buyerPostcode,
		CardQuota:              cardQuota,
		InterestFreeByMerchant: interestFreeByMerchant,
		CustomData:             customData,
		NoticeUrl:              noticeUrl,
	}

	res, err := subscribe.Onetime(iamport.Authenticate, req)
	if err != nil {
		return nil, err
	}

	if res.Code != util.CodeOK {
		return nil, errors.New(res.Message)
	}

	return res.Response, nil
}

// AgainPayment 이전 결제 데이터를 이용한 재결제
//
// POST /subscribe/payments/again
func (iamport *Iamport) AgainPayment(
	customerUID, merchantUID string,
	amount, taxFree int32,
	name string,
	buyerName, buyerEmail, buyerTel, buyerAddr, buyerPostcode string,
	cardQuota int32, interestFreeByMerchant bool,
	customData, noticeUrl string,
) (*Typepayment.Payment, error) {

	if merchantUID == "" || customerUID == "" {
		return nil, errors.New(ErrMustExistImpUIDorMerchantUID)
	}

	if amount < 0 {
		return nil, errors.New(ErrInvalidAmount)
	}

	req := &TypeSubscribe.AgainPaymentRequest{
		CustomerUid:            customerUID,
		MerchantUid:            merchantUID,
		Amount:                 amount,
		TaxFree:                taxFree,
		Name:                   name,
		BuyerName:              buyerName,
		BuyerEmail:             buyerEmail,
		BuyerTel:               buyerTel,
		BuyerAddr:              buyerAddr,
		BuyerPostcode:          buyerPostcode,
		CardQuota:              cardQuota,
		InterestFreeByMerchant: interestFreeByMerchant,
		CustomData:             customData,
		NoticeUrl:              noticeUrl,
	}

	res, err := subscribe.Again(iamport.Authenticate, req)
	if err != nil {
		return nil, err
	}

	if res.Code != util.CodeOK {
		return nil, errors.New(res.Message)
	}

	return res.Response, nil
}

// SchedulePayment 예약결제
//
// POST /subscribe/payments/schedule
func (iamport *Iamport) SchedulePayment(
	customerUID string, checkingAmount int32,
	cardNumber, expiry, birth, pwd2Digit, pg string,
	schedules []*TypeSubscribe.PaymentScheduleParam,
) ([]*TypeSubscribe.UnitSchedulePaymentResponse, error) {

	if customerUID == "" {
		return nil, errors.New(ErrMustExistCustomerUID)
	}

	req := &TypeSubscribe.SchedulePayemntRequest{
		CustomerUid:    customerUID,
		CheckingAmount: checkingAmount,
		CardNumber:     cardNumber,
		Expiry:         expiry,
		Birth:          birth,
		Pwd_2Digit:     pwd2Digit,
		Pg:             pg,
		Schedules:      schedules,
	}

	res, err := subscribe.Schedule(iamport.Authenticate, req)
	if err != nil {
		return nil, err
	}

	if res.Code != util.CodeOK {
		return nil, errors.New(res.Message)
	}

	return res.Response, nil
}

// UnschedulePayment 예약결제 취소
//
// POST /subscribe/payments/unschedule
func (iamport *Iamport) UnschedulePayment(customerUID string, merchantUID []string) ([]*TypeSubscribe.UnitSchedulePaymentResponse, error) {
	if customerUID == "" {
		return nil, errors.New(ErrMustExistCustomerUID)
	}

	req := &TypeSubscribe.UnschedulePaymentRequest{
		CustomerUid: customerUID,
		MerchantUid: merchantUID,
	}

	res, err := subscribe.Unschedule(iamport.Authenticate, req)
	if err != nil {
		return nil, err
	}

	if res.Code != util.CodeOK {
		return nil, errors.New(res.Message)
	}

	return res.Response, nil
}

// GetScheduledPaymentByMerchantUID Merchant UID별로 예약결제내역을 가져오는 API
//
// GET /subscribe/payments/schedule/{merchant_uid}
func (iamport *Iamport) GetScheduledPaymentByMerchantUID(merchantUID string) (*TypeSubscribe.UnitSchedulePaymentResponse, error) {
	if merchantUID == "" {
		return nil, errors.New(ErrMustExistMerchantUID)
	}

	req := &TypeSubscribe.GetPaymentScheduleRequest{
		MerchantUid: merchantUID,
	}

	res, err := subscribe.GetScheduledPaymentByMerchantUID(iamport.Authenticate, req)
	if err != nil {
		return nil, err
	}

	if res.Code != util.CodeOK {
		return nil, errors.New(res.Message)
	}

	return res.Response, nil
}

// GetScheduledPaymentByCustomerUID Customer UID별로 예약결제내역을 가져오는 API
//
// GET /subscribe/payments/schedule/customer/{customer_uid}
func (iamport *Iamport) GetScheduledPaymentByCustomerUID(
	customerUID string,
	page, from, to int32,
	scheduleStatus string,
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

	res, err := subscribe.GetScheduledPaymentByCustomerUID(iamport.Authenticate, req)
	if err != nil {
		return nil, err
	}

	if res.Code != util.CodeOK {
		return nil, errors.New(res.Message)
	}

	return res.GetResponse(), nil
}
