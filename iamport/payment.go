package iamport

import (
	"errors"
	"time"

	"github.com/iamport/go-iamport/payment"
	"github.com/iamport/go-iamport/util"
	TypePayment "github.com/iamport/interface/gen_src/go/v1/payment"
)

// GetPaymentImpUID imp_uid로 결제 정보 가져오기
//
// GET /payments/{imp_uid}
func (iamport *Iamport) GetPaymentImpUID(iuid string) (*TypePayment.Payment, error) {
	if iuid == "" {
		return nil, errors.New(ErrMustExistImpUID)
	}

	token, err := iamport.Authenticate.GetToken()
	if err != nil {
		return nil, err
	}

	reqPaymentImpUID := &TypePayment.PaymentRequest{
		ImpUid: iuid,
	}

	res, err := payment.GetByImpUID(
		iamport.Authenticate.Client, iamport.Authenticate.APIUrl,
		token, reqPaymentImpUID,
	)
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

	token, err := iamport.Authenticate.GetToken()
	if err != nil {
		return nil, err
	}

	req := &TypePayment.PaymentsRequest{
		ImpUid: iuids,
	}

	res, err := payment.GetByImpUIDs(
		iamport.Authenticate.Client, iamport.Authenticate.APIUrl,
		token, req,
	)
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

	token, err := iamport.Authenticate.GetToken()
	if err != nil {
		return nil, err
	}

	merchantUIDPaymentReq := &TypePayment.PaymentMerchantUidRequest{
		MerchantUid: muid,
		Status:      status,
		Sorting:     sorting,
	}

	res, err := payment.GetByMerchantUID(
		iamport.Authenticate.Client, iamport.Authenticate.APIUrl,
		token, merchantUIDPaymentReq,
	)

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

	token, err := iamport.Authenticate.GetToken()
	if err != nil {
		return nil, err
	}

	merchantUIDPaymentReq := &TypePayment.PaymentsMerchantUidRequest{
		MerchantUid: muid,
		Status:      status,
		Sorting:     sorting,
		Page:        int32(page),
	}

	res, err := payment.GetByMerchantUIDs(
		iamport.Authenticate.Client, iamport.Authenticate.APIUrl,
		token, merchantUIDPaymentReq,
	)
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

	token, err := iamport.Authenticate.GetToken()
	if err != nil {
		return nil, err
	}

	req := &TypePayment.PaymentStatusRequest{

		Status:  status,
		Page:    int32(page),
		From:    int32(from.Unix()),
		Limit:   int32(limit),
		Sorting: sorting,
		To:      int32(to.Unix()),
	}

	res, err := payment.GetByStatus(
		iamport.Authenticate.Client, iamport.Authenticate.APIUrl,
		token, req,
	)

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

	token, err := iamport.Authenticate.GetToken()
	if err != nil {
		return nil, err
	}

	reqPaymentImpUID := &TypePayment.PaymentBalanceRequest{
		ImpUid: iuid,
	}

	res, err := payment.GetBalanceByImpUID(
		iamport.Authenticate.Client, iamport.Authenticate.APIUrl,
		token, reqPaymentImpUID,
	)
	if err != nil {
		return nil, err
	}

	if res.Code != util.CodeOK {
		return nil, errors.New(res.Message)
	}

	return res.Response, nil
}

// CancelPaymentImpUID imp_uid로 결제 취소하기
//
// POST /payments/cancel
func (iamport *Iamport) CancelPaymentImpUID(iuid string, merchantUID string, amount float64, taxFree float64, checkSum float64, reason string, refundHolder string, refundBank string, refundAccount string) (*TypePayment.Payment, error) {
	if iuid == "" && merchantUID == "" {
		return nil, errors.New(ErrMustExistImpUIDorMerchantUID)
	}

	if amount < 0 {
		return nil, errors.New(ErrInvalidAmount)
	}

	token, err := iamport.Authenticate.GetToken()
	if err != nil {
		return nil, err
	}

	req := &TypePayment.PaymentCancelRequest{
		ImpUid:        iuid,
		MerchantUid:   merchantUID,
		Amount:        amount,
		TaxFree:       taxFree,
		Checksum:      checkSum,
		Reason:        reason,
		RefundHolder:  refundHolder,
		RefundBank:    refundBank,
		RefundAccount: refundAccount,
	}

	res, err := payment.Cancel(
		iamport.Authenticate.Client, iamport.Authenticate.APIUrl,
		token, req,
	)
	if err != nil {
		return nil, err
	}

	if res.Code != util.CodeOK {
		return nil, errors.New(res.Message)
	}

	return res.Response, nil
}

// PreparePayment 결제 정보 사전 등록하기
//
// POST /payments/prepare
func (iamport *Iamport) PreparePayment(merchantUID string, amount float64) (*TypePayment.Prepare, error) {
	if merchantUID == "" {
		return nil, errors.New(ErrMustExistMerchantUID)
	}

	if amount < 0 {
		return nil, errors.New(ErrInvalidAmount)
	}

	token, err := iamport.Authenticate.GetToken()
	if err != nil {
		return nil, err
	}

	req := &TypePayment.PaymentPrepareRequest{
		MerchantUid: merchantUID,
		Amount:      amount,
	}

	res, err := payment.Prepare(
		iamport.Authenticate.Client, iamport.Authenticate.APIUrl,
		token, req,
	)
	if err != nil {
		return nil, err
	}

	if res.Code != util.CodeOK {
		return nil, errors.New(res.Message)
	}

	return res.Response, nil
}

// GetPreparePayment 사전 등록된 결제 정보 보기
//
// GET /payments/prepare/{merchant_uid}
func (iamport *Iamport) GetPreparePayment(merchantUID string) (*TypePayment.Prepare, error) {
	if merchantUID == "" {
		return nil, errors.New(ErrMustExistMerchantUID)
	}

	token, err := iamport.Authenticate.GetToken()
	if err != nil {
		return nil, err
	}

	req := &TypePayment.PaymentGetPrepareRequest{
		MerchantUid: merchantUID,
	}

	res, err := payment.GetPrepareByMerchantUID(
		iamport.Authenticate.Client, iamport.Authenticate.APIUrl,
		token, req,
	)
	if err != nil {
		return nil, err
	}

	if res.Code != util.CodeOK {
		return nil, errors.New(res.Message)
	}

	return res.Response, nil
}
