package payment

import (
	"net/http"
	"strconv"
	"strings"

	"google.golang.org/protobuf/encoding/protojson"

	"github.com/iamport/go-iamport/util"

	"github.com/iamport/interface/gen_src/go/v1/payment"
)

const (
	URLPayments = "/payments"
	URLFind     = "/find"
	URLFindAll  = "/findAll"
	URLBalance  = "/balance"
	URLStatus   = "/status"
	URLCancel   = "/cancel"
	URLPrepare  = "/prepare"

	URLParamSorting = "sorting="
	URLParamPage    = "page="
	URLParamLimit   = "limit="
	URLParamFrom    = "from="
	URLParamTo      = "to="
	URLParamImpUids = "imp_uid[]="
)

// GetByImpUID - GET /payments/{imp_uid}
// 아임포트 고유번호로 결제내역을 확인합니다
func GetByImpUID(client *http.Client, apiDomain string, token string, params *payment.PaymentRequest) (*payment.PaymentResponse, error) {
	urls := []string{apiDomain, URLPayments, "/", params.GetImpUid()}
	urlPayment := strings.Join(urls, "")

	res, err := util.Call(client, token, urlPayment, util.GET)
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
func GetByImpUIDs(client *http.Client, apiDomain string, token string, params *payment.PaymentsRequest) (*payment.PaymentsResponse, error) {
	urls := []string{apiDomain, URLPayments}

	isFirstQuery := true
	for _, impUID := range params.GetImpUid() {
		urls = append(urls, []string{util.GetQueryPrefix(&isFirstQuery), URLParamImpUids, impUID}...)
	}
	urlPayment := strings.Join(urls, "")

	res, err := util.Call(client, token, urlPayment, util.GET)
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
func GetByMerchantUID(client *http.Client, apiDomain string, token string, params *payment.PaymentMerchantUidRequest) (*payment.PaymentMerchantUidResponse, error) {
	urls := []string{apiDomain, URLPayments, URLFind, "/", params.GetMerchantUid(), "/"}

	if params.Status != "" {
		urls = append(urls, params.GetStatus())
	}

	isFirstQuery := true
	if params.Sorting != "" {
		urls = append(urls, []string{util.GetQueryPrefix(&isFirstQuery), URLParamSorting, params.GetSorting()}...)
	}

	urlPayment := strings.Join(urls, "")

	res, err := util.Call(client, token, urlPayment, util.GET)
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
func GetByMerchantUIDs(client *http.Client, apiDomain string, token string, params *payment.PaymentsMerchantUidRequest) (*payment.PaymentsMerchantUidResponse, error) {
	urls := []string{apiDomain, URLPayments, URLFindAll, "/", params.GetMerchantUid(), "/"}

	if params.Status != "" {
		urls = append(urls, params.GetStatus())
	}

	isFirstQuery := true
	if params.Sorting != "" {
		urls = append(urls, []string{util.GetQueryPrefix(&isFirstQuery), URLParamSorting, params.GetSorting()}...)
	}

	if params.Page > 0 {
		urls = append(urls, []string{util.GetQueryPrefix(&isFirstQuery), URLParamPage, strconv.Itoa(int(params.GetPage()))}...)
	}

	urlPayment := strings.Join(urls, "")

	res, err := util.Call(client, token, urlPayment, util.GET)
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
func GetByStatus(client *http.Client, apiDomain string, token string, params *payment.PaymentStatusRequest) (*payment.PaymentStatusResponse, error) {
	if params.Status == "" {
		params.Status = util.StatusAll
	}
	urls := []string{apiDomain, URLPayments, URLStatus, "/", params.GetStatus()}

	isFirstQuery := true
	if params.GetPage() > 0 {
		urls = append(urls, []string{util.GetQueryPrefix(&isFirstQuery), URLParamPage, strconv.Itoa(int(params.GetPage()))}...)
	}

	if params.GetLimit() > 0 {
		urls = append(urls, []string{util.GetQueryPrefix(&isFirstQuery), URLParamLimit, strconv.Itoa(int(params.GetLimit()))}...)
	}

	if params.GetFrom() > 0 {
		urls = append(urls, []string{util.GetQueryPrefix(&isFirstQuery), URLParamFrom, strconv.Itoa(int(params.GetFrom()))}...)
	}

	if params.GetTo() > 0 {
		urls = append(urls, []string{util.GetQueryPrefix(&isFirstQuery), URLParamTo, strconv.Itoa(int(params.GetTo()))}...)
	}

	if params.GetSorting() != "" {
		urls = append(urls, []string{util.GetQueryPrefix(&isFirstQuery), URLParamSorting, params.GetSorting()}...)
	}

	urlPayment := strings.Join(urls, "")

	res, err := util.Call(client, token, urlPayment, util.GET)
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
func GetBalanceByImpUID(client *http.Client, apiDomain string, token string, params *payment.PaymentBalanceRequest) (*payment.PaymentBalanceResponse, error) {
	urls := []string{apiDomain, URLPayments, "/", params.ImpUid, URLBalance}
	urlPayment := strings.Join(urls, "")

	res, err := util.Call(client, token, urlPayment, util.GET)
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

// Cancel - POST /payments/cancel
// 승인된 결제를 취소합니다.
// 신용카드/실시간계좌이체/휴대폰소액결제의 경우 즉시 취소처리가 이뤄지게 되며, 가상계좌의 경우는 환불받으실 계좌정보를 같이 전달해주시면 환불정보가 PG사에 등록되어 익영업일에 처리됩니다.(가상계좌 환불관련 특약계약 필요)
func Cancel(client *http.Client, apiDomain string, token string, params *payment.PaymentCancelRequest) (*payment.PaymentCancelResponse, error) {
	urls := []string{apiDomain, URLPayments, URLCancel}
	urlCancel := strings.Join(urls, "")

	marshaler := protojson.MarshalOptions{
		UseProtoNames: true,
	}
	jsonBytes, err := marshaler.Marshal(params)
	if err != nil {
		return nil, err
	}

	res, err := util.CallWithForm(client, token, urlCancel, util.POST, jsonBytes)
	if err != nil {
		return nil, err
	}

	cancelRes := payment.PaymentCancelResponse{}
	err = protojson.Unmarshal(res, &cancelRes)
	if err != nil {
		return nil, err
	}

	return &cancelRes, nil
}

// Prepare - POST /payments/prepare
// (아임포트 javascript사용)인증방식의 결제를 진행할 때 결제금액 위변조시 결제진행자체를 block하기 위해 결제예정금액을 사전등록하는 기능입니다.
// 이 API를 통해 사전등록된 가맹점 주문번호(merchant_uid)에 대해, IMP.request_pay()에 전달된 merchant_uid가 일치하는 주문의 결제금액이 다른 경우 PG사 결제창 호출이 중단됩니다.
func Prepare(client *http.Client, apiDomain string, token string, params *payment.PaymentPrepareRequest) (*payment.PaymentPrepareResponse, error) {
	urls := []string{apiDomain, URLPayments, URLPrepare}
	urlPrepare := strings.Join(urls, "")

	marshaler := protojson.MarshalOptions{
		UseProtoNames: true,
	}

	jsonBytes, err := marshaler.Marshal(params)

	res, err := util.CallWithForm(client, token, urlPrepare, util.POST, jsonBytes)
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
func GetPrepareByMerchantUID(client *http.Client, apiDomain string, token string, params *payment.PaymentGetPrepareRequest) (*payment.PaymentPrepareResponse, error) {
	urls := []string{apiDomain, URLPayments, URLPrepare, "/", params.GetMerchantUid()}
	urlPrepare := strings.Join(urls, "")

	res, err := util.Call(client, token, urlPrepare, util.GET)
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
