package payment

import (
	"net/url"
	"strconv"
	"strings"

	"google.golang.org/protobuf/encoding/protojson"

	"github.com/joowonyun/go-iamport/util"

	"github.com/joowonyun/go-iamport/authenticate"
	"github.com/joowonyun/interface/build/go/payment"
)

const (
	URLPayments = "/payments"
	URLFind     = "/find"
	URLFindAll  = "/findAll"
	URLBalance  = "/balance"
	URLStatus   = "/status"
	URLCancle   = "/cancle"
	URLPrepare  = "/prepare"

	URLParamSorting = "sorting="
	URLParamPage    = "page="
	URLParamLimit   = "limit="
	URLParamFrom    = "from="
	URLParamTo      = "to="
	URLParamImpUids = "imp_uid[]="

	ImpUID        = "imp_uid"
	MerchantUID   = "merchant_uid"
	Amount        = "amount"
	TaxFree       = "tax_free"
	Checksum      = "checksum"
	Reason        = "reason"
	RefundHolder  = "refund_holder"
	RefundBank    = "refund_bank"
	RefundAccount = "refund_account"
)

// GetByImpUID - GET /payments/{imp_uid}
// 아임포트 고유번호로 결제내역을 확인합니다
func GetByImpUID(auth *authenticate.Authenticate, params *payment.PaymentRequest) (*payment.PaymentResponse, error) {
	urls := []string{auth.APIUrl, URLPayments, "/", params.ImpUid}
	urlPayment := strings.Join(urls, "")

	token, err := auth.GetToken()
	if err != nil {
		return nil, err
	}

	res, err := util.CallGet(auth.Client, token, urlPayment)
	if err != nil {
		return nil, err
	}

	paymentRes := payment.PaymentResponse{}
	err = protojson.Unmarshal(res, &paymentRes)
	if err != nil {
		return nil, err
	}

	return &paymentRes, nil
}

// GetByImpUIDs - GET /payments
// 여러 개의 아임포트 고유번호로 결제내역을 한 번에 조회합니다.(최대 100개)
// (예시) /payments?imp_uid[]=imp_448280090638&imp_uid[]=imp_448280090639
func GetByImpUIDs(auth *authenticate.Authenticate, params *payment.PaymentsRequest) (*payment.PaymentsResponse, error) {
	urls := []string{auth.APIUrl, URLPayments}

	isFirstQuery := true
	for _, impUID := range params.GetImpUid() {
		urls = append(urls, []string{util.GetQueryPrefix(&isFirstQuery), URLParamImpUids, impUID}...)
	}
	urlPayment := strings.Join(urls, "")

	token, err := auth.GetToken()
	if err != nil {
		return nil, err
	}

	res, err := util.CallGet(auth.Client, token, urlPayment)
	if err != nil {
		return nil, err
	}

	paymentsRes := payment.PaymentsResponse{}
	err = protojson.Unmarshal(res, &paymentsRes)
	if err != nil {
		return nil, err
	}

	return &paymentsRes, nil
}

// GetByMerchantUID - GET /payments/find/{merchant_uid}/{payment_status}
// 동일한 merchant_uid가 여러 건 존재하는 경우, 정렬 기준에 따라 가장 첫 번째 해당되는 건을 반환합니다. (모든 내역에 대한 조회가 필요하시면 /payments/findAll/{merchant_uid}를 사용해주세요.)
// payment_status를 추가로 지정하시면, 해당 status에 해당하는 가장 최신 데이터를 반환합니다.
func GetByMerchantUID(auth *authenticate.Authenticate, params *payment.PaymentMerchantUidRequest) (*payment.PaymentMerchantUidResponse, error) {
	urls := []string{auth.APIUrl, URLPayments, URLFind, "/", params.MerchantUid, "/"}

	if params.Status != "" {
		urls = append(urls, params.Status)
	}

	isFirstQuery := true
	if params.Sorting != "" {
		urls = append(urls, []string{util.GetQueryPrefix(&isFirstQuery), URLParamSorting, params.Sorting}...)
	}

	urlPayment := strings.Join(urls, "")

	token, err := auth.GetToken()
	if err != nil {
		return nil, err
	}

	res, err := util.CallGet(auth.Client, token, urlPayment)
	if err != nil {
		return nil, err
	}

	paymentRes := payment.PaymentMerchantUidResponse{}
	err = protojson.Unmarshal(res, &paymentRes)
	if err != nil {
		return nil, err
	}

	return &paymentRes, nil
}

// GetByMerchantUIDs - GET /payments/findAll/{merchant_uid}/{payment_status}
// 동일한 merchant_uid의 모든 내역에 대한 조회
func GetByMerchantUIDs(auth *authenticate.Authenticate, params *payment.PaymentsMerchantUidRequest) (*payment.PaymentsMerchantUidResponse, error) {
	urls := []string{auth.APIUrl, URLPayments, URLFindAll, "/", params.MerchantUid, "/"}

	if params.Status != "" {
		urls = append(urls, params.Status)
	}

	isFirstQuery := true
	if params.Sorting != "" {
		urls = append(urls, []string{util.GetQueryPrefix(&isFirstQuery), URLParamSorting, params.Sorting}...)
	}

	if params.Page > 0 {
		urls = append(urls, []string{util.GetQueryPrefix(&isFirstQuery), URLParamPage, strconv.Itoa(int(params.Page))}...)
	}

	urlPayment := strings.Join(urls, "")
	token, err := auth.GetToken()
	if err != nil {
		return nil, err
	}

	res, err := util.CallGet(auth.Client, token, urlPayment)
	if err != nil {
		return nil, err
	}

	paymentRes := payment.PaymentsMerchantUidResponse{}
	err = protojson.Unmarshal(res, &paymentRes)
	if err != nil {
		return nil, err
	}

	return &paymentRes, nil
}

// GetByStatus - GET /payments/status/{payment_status}
// 미결제/결제완료/결제취소/결제실패 상태 별로 검색할 수 있습니다.(20건씩 최신순 페이징)
// 검색기간은 최대 90일까지이며 to파라메터의 기본값은 현재 unix timestamp이고 from파라메터의 기본값은 to파라메터 기준으로 90일 전입니다. 때문에, from/to 파라메터가 없이 호출되면 현재 시점 기준으로 최근 90일 구간에 대한 데이터를 검색하게 됩니다.
// from, to 파라메터를 지정하여 90일 단위로 과거 데이터 조회는 가능합니다.
func GetByStatus(auth *authenticate.Authenticate, params *payment.PaymentStatusRequest) (*payment.PaymentStatusResponse, error) {
	if params.Status == "" {
		params.Status = util.StatusAll
	}
	urls := []string{auth.APIUrl, URLPayments, URLStatus, "/", params.Status}

	isFirstQuery := true
	if params.Page > 0 {
		urls = append(urls, []string{util.GetQueryPrefix(&isFirstQuery), URLParamPage, strconv.Itoa(int(params.Page))}...)
	}

	if params.Limit > 0 {
		urls = append(urls, []string{util.GetQueryPrefix(&isFirstQuery), URLParamLimit, strconv.Itoa(int(params.Limit))}...)
	}

	if params.From > 0 {
		urls = append(urls, []string{util.GetQueryPrefix(&isFirstQuery), URLParamFrom, strconv.Itoa(int(params.From))}...)
	}

	if params.To > 0 {
		urls = append(urls, []string{util.GetQueryPrefix(&isFirstQuery), URLParamTo, strconv.Itoa(int(params.To))}...)
	}

	if params.Sorting != "" {
		urls = append(urls, []string{util.GetQueryPrefix(&isFirstQuery), URLParamSorting, params.Sorting}...)
	}

	urlPayment := strings.Join(urls, "")
	token, err := auth.GetToken()
	if err != nil {
		return nil, err
	}

	res, err := util.CallGet(auth.Client, token, urlPayment)
	if err != nil {
		return nil, err
	}

	paymentRes := payment.PaymentStatusResponse{}
	err = protojson.Unmarshal(res, &paymentRes)
	if err != nil {
		return nil, err
	}

	return &paymentRes, nil
}

// GetBalanceByImpUID - GET /payments/{imp_uid}/balance
// 아임포트 고유번호로 결제수단별 금액 상세정보를 확인합니다.(현재, PAYCO결제수단에 한해 제공되고 있습니다. 타 PG사의 경우 파라메터 검증 등 검토/협의 단계에 있습니다.)
func GetBalanceByImpUID(auth *authenticate.Authenticate, params *payment.PaymentBalanceRequest) (*payment.PaymentBalanceResponse, error) {
	urls := []string{auth.APIUrl, URLPayments, "/", params.ImpUid, URLBalance}
	urlPayment := strings.Join(urls, "")

	token, err := auth.GetToken()
	if err != nil {
		return nil, err
	}

	res, err := util.CallGet(auth.Client, token, urlPayment)
	if err != nil {
		return nil, err
	}

	paymentRes := payment.PaymentBalanceResponse{}
	err = protojson.Unmarshal(res, &paymentRes)
	if err != nil {
		return nil, err
	}

	return &paymentRes, nil
}

// Cancle - POST /payments/cancel
// 승인된 결제를 취소합니다.
// 신용카드/실시간계좌이체/휴대폰소액결제의 경우 즉시 취소처리가 이뤄지게 되며, 가상계좌의 경우는 환불받으실 계좌정보를 같이 전달해주시면 환불정보가 PG사에 등록되어 익영업일에 처리됩니다.(가상계좌 환불관련 특약계약 필요)
func Cancle(auth *authenticate.Authenticate, params *payment.PaymentCancleRequest) (*payment.PaymentCancleResponse, error) {
	urls := []string{auth.APIUrl, URLPayments, URLCancle}
	urlCancle := strings.Join(urls, "")

	token, err := auth.GetToken()
	if err != nil {
		return nil, err
	}

	form := url.Values{}
	if params.ImpUid != "" {
		form.Set(ImpUID, params.ImpUid)
	}
	if params.MerchantUid != "" {
		form.Set(MerchantUID, params.MerchantUid)
	}
	if params.Amount != 0 {
		form.Set(Amount, strconv.FormatFloat(params.Amount, 'f', -1, 64))
	}
	if params.TaxFree != 0 {
		form.Set(TaxFree, strconv.FormatFloat(params.TaxFree, 'f', -1, 64))
	}
	if params.Checksum != 0 {
		form.Set(Checksum, strconv.FormatFloat(params.Checksum, 'f', -1, 64))
	}
	if params.Reason != "" {
		form.Set(Reason, params.Reason)
	}
	if params.RefundHolder != "" {
		form.Set(RefundHolder, params.RefundHolder)
	}
	if params.RefundBank != "" {
		form.Set(RefundBank, params.RefundBank)
	}
	if params.RefundAccount != "" {
		form.Set(RefundAccount, params.RefundAccount)
	}

	res, err := util.CallPostForm(auth.Client, token, urlCancle, form)
	if err != nil {
		return nil, err
	}

	cancleRes := payment.PaymentCancleResponse{}
	err = protojson.Unmarshal(res, &cancleRes)
	if err != nil {
		return nil, err
	}

	return &cancleRes, nil
}

// Prepare - POST /payments/prepare
// (아임포트 javascript사용)인증방식의 결제를 진행할 때 결제금액 위변조시 결제진행자체를 block하기 위해 결제예정금액을 사전등록하는 기능입니다.
// 이 API를 통해 사전등록된 가맹점 주문번호(merchant_uid)에 대해, IMP.request_pay()에 전달된 merchant_uid가 일치하는 주문의 결제금액이 다른 경우 PG사 결제창 호출이 중단됩니다.
func Prepare(auth *authenticate.Authenticate, params *payment.PaymentPrepareRequest) (*payment.PaymentPrepareResponse, error) {
	urls := []string{auth.APIUrl, URLPayments, URLPrepare}
	urlPrepare := strings.Join(urls, "")

	token, err := auth.GetToken()
	if err != nil {
		return nil, err
	}

	form := url.Values{}
	form.Set(MerchantUID, params.MerchantUid)
	form.Set(Amount, strconv.FormatFloat(params.Amount, 'f', -1, 64))

	res, err := util.CallPostForm(auth.Client, token, urlPrepare, form)
	if err != nil {
		return nil, err
	}

	prepareRes := payment.PaymentPrepareResponse{}
	err = protojson.Unmarshal(res, &prepareRes)
	if err != nil {
		return nil, err
	}

	return &prepareRes, nil
}

// GetPrepareByMerchantUID - GET /payments/prepare/{merchant_uid}
// /payments/prepare로 이미 등록되어있는 사전등록 결제정보를 조회합니다
func GetPrepareByMerchantUID(auth *authenticate.Authenticate, params *payment.PaymentPrepareRequest) (*payment.PaymentPrepareResponse, error) {
	urls := []string{auth.APIUrl, URLPayments, URLPrepare, "/", params.MerchantUid}
	urlPrepare := strings.Join(urls, "")

	token, err := auth.GetToken()
	if err != nil {
		return nil, err
	}

	res, err := util.CallGet(auth.Client, token, urlPrepare)
	if err != nil {
		return nil, err
	}

	prepareRes := payment.PaymentPrepareResponse{}
	err = protojson.Unmarshal(res, &prepareRes)
	if err != nil {
		return nil, err
	}

	return &prepareRes, nil
}

// TODO cancle, prepare
