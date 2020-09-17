package iamport

import (
	"errors"
	"time"

	"github.com/joowonyun/go-iamport/payment"
	"github.com/joowonyun/go-iamport/util"
	TypePayment "github.com/joowonyun/interface/build/go/payment"
)

// GetPaymentImpUID imp_uid로 결제 정보 가져오기
//
// GET /payments/{imp_uid}
func (iamport *Iamport) GetPaymentImpUID(iuid string) (*TypePayment.Payment, error) {
	if iuid == "" {
		return nil, errors.New(ErrMustExistImpUID)
	}

	reqPaymentImpUID := &TypePayment.PaymentRequest{
		ImpUid: iuid,
	}

	res, err := payment.GetByImpUID(iamport.Authenticate, reqPaymentImpUID)
	if err != nil {
		return nil, err
	}

	if res.Code != util.CodeOK {
		return nil, errors.New(res.Message)
	}

	return res.Response, nil
}

// GetPaymentsImpUIDs 여러개의 imp_uid로 결제 정보 가져오기
//
// GET /payments
func (iamport *Iamport) GetPaymentsImpUIDs(iuids []string) ([]*TypePayment.Payment, error) {
	if len(iuids) < 0 {
		return nil, errors.New(ErrMustExistImpUID)
	}

	req := &TypePayment.PaymentsRequest{
		ImpUid: iuids,
	}

	res, err := payment.GetByImpUIDs(iamport.Authenticate, req)
	if err != nil {
		return nil, err
	}

	if res.Code != util.CodeOK {
		return nil, errors.New(res.Message)
	}

	return res.Response, nil
}

// GetPaymentMerchantUID merchant_uid로 결제 정보 가져오기
//
// GET /payments/find/{merchant_uid}
func (iamport *Iamport) GetPaymentMerchantUID(muid string, status string, sorting string) (*TypePayment.Payment, error) {
	if muid == "" {
		return nil, errors.New(ErrMustExistMerchantUID)
	}

	if !util.ValidateStatusParameter(status) {
		return nil, errors.New(ErrInvalidSortParam)
	}

	if !util.ValidateSortParameter(sorting) {
		return nil, errors.New(ErrInvalidSortParam)
	}

	merchantUIDPaymentReq := &TypePayment.PaymentMerchantUidRequest{
		MerchantUid: muid,
		Status:      status,
		Sorting:     sorting,
	}

	res, err := payment.GetByMerchantUID(iamport.Authenticate, merchantUIDPaymentReq)
	if err != nil {
		return nil, err
	}

	if res.Code != util.CodeOK {
		return nil, errors.New(res.Message)
	}

	return res.Response, nil
}

// GetPaymentsMerchantUID merchant_uid로 모든 결제 정보 가져오기
//
// GET /payments/find/{merchant_uid}
func (iamport *Iamport) GetPaymentsMerchantUID(muid string, status string, sorting string, page int) (*TypePayment.PaymentPage, error) {
	if muid == "" {
		return nil, errors.New(ErrMustExistMerchantUID)
	}

	if !util.ValidateStatusParameter(status) {
		return nil, errors.New(ErrInvalidSortParam)
	}

	if !util.ValidateSortParameter(sorting) {
		return nil, errors.New(ErrInvalidSortParam)
	}

	if page < 0 {
		return nil, errors.New(ErrInvalidPage)
	}

	merchantUIDPaymentReq := &TypePayment.PaymentsMerchantUidRequest{
		MerchantUid: muid,
		Status:      status,
		Sorting:     sorting,
		Page:        int32(page),
	}

	res, err := payment.GetByMerchantUIDs(iamport.Authenticate, merchantUIDPaymentReq)
	if err != nil {
		return nil, err
	}

	if res.Code != util.CodeOK {
		return nil, errors.New(res.Message)
	}

	return res.Response, nil
}

// GetPaymentsStatus 결제 상태에 따른 결제 정보들 가져오기
//
// GET /payments/status/{payment_status}
func (iamport *Iamport) GetPaymentsStatus(status string, page int, limit int, from time.Time, to time.Time, sorting string) (*TypePayment.PaymentPage, error) {
	if !util.ValidateSortParameter(sorting) {
		return nil, errors.New(ErrInvalidSortParam)
	}

	if !util.ValidateStatusParameter(status) {
		return nil, errors.New(ErrInvalidStatusParam)
	}

	if page < 0 {
		return nil, errors.New(ErrInvalidPage)
	}

	if limit < 0 {
		return nil, errors.New(ErrInvalidLimit)
	}

	if from.After(to) {
		return nil, errors.New(ErrInvalidFrom)
	}

	if from.AddDate(0, 3, 0).Before(to) {
		return nil, errors.New(ErrInvalidTo)
	}

	req := &TypePayment.PaymentStatusRequest{

		Status:  status,
		Page:    int32(page),
		From:    int32(from.Unix()),
		Limit:   int32(limit),
		Sorting: sorting,
		To:      int32(to.Unix()),
	}

	res, err := payment.GetByStatus(iamport.Authenticate, req)
	if err != nil {
		return nil, err
	}

	if res.Code != util.CodeOK {
		return nil, errors.New(res.Message)
	}

	return res.Response, nil
}

// GetPaymentBalanceImpUID imp_uid로 결제 정보 가져오기
//
// GET /payments/{imp_uid}
func (iamport *Iamport) GetPaymentBalanceImpUID(iuid string) (*TypePayment.PaymentBalance, error) {
	if iuid == "" {
		return nil, errors.New(ErrMustExistImpUID)
	}

	reqPaymentImpUID := &TypePayment.PaymentBalanceRequest{
		ImpUid: iuid,
	}

	res, err := payment.GetBalanceByImpUID(iamport.Authenticate, reqPaymentImpUID)
	if err != nil {
		return nil, err
	}

	if res.Code != util.CodeOK {
		return nil, errors.New(res.Message)
	}

	return res.Response, nil
}