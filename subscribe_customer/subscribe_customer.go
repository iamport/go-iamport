package subscribe_customer

import (
	urllib "net/url"
	"strconv"
	"strings"

	"google.golang.org/protobuf/encoding/protojson"

	"github.com/iamport/go-iamport/util"

	"github.com/iamport/go-iamport/authenticate"
	subscribe_dup "github.com/iamport/interface/gen_src/go/v1/subscribe"
	subscribe "github.com/iamport/interface/gen_src/go/v1/subscribe_customers"
)

const (
	URLSubscribe = "/subscribe"
	URLCustomers = "/customers"
	URLPayments  = "/payments"
	URLSchedules = "/schedules"

	URLParamCustomerUID    = "customer_uid[]="
	URLParamReason         = "reason="
	URLParamRequester      = "requester="
	URLParamPage           = "page="
	URLParamFrom           = "from="
	URLParamTo             = "to="
	URLParamScheduleStatus = "schedule-status="
)

// GetMultipleBillingKeysByCustomer - GET /subscribe/customers
// 여러 빌링키를 한 번에 조회하는 API
func GetMultipleBillingKeysByCustomer(auth *authenticate.Authenticate, params *subscribe.GetMultipleCustomerBillingKeyRequest) (*subscribe.GetMultipleCustomerBillingKeyResponse, error) {
	urls := []string{auth.APIUrl, URLSubscribe, URLCustomers}

	token, err := auth.GetToken()
	if err != nil {
		return nil, err
	}

	isFirstQuery := true
	for _, customerUID := range params.GetCustomerUid() {
		urls = append(urls, []string{util.GetQueryPrefix(&isFirstQuery), URLParamCustomerUID, customerUID}...)
	}
	urlGetBillingkeys := strings.Join(urls, "")

	res, err := util.Call(auth.Client, token, urlGetBillingkeys, util.GET)
	if err != nil {
		return nil, err
	}

	getBillingKeysRes := subscribe.GetMultipleCustomerBillingKeyResponse{}
	err = protojson.Unmarshal(res, &getBillingKeysRes)
	if err != nil {
		return nil, err
	}

	return &getBillingKeysRes, nil
}

// DeleteBillingKey - DELETE /subscribe/customers/{customer_uid}
// 해당 빌링키 삭제
func DeleteBillingKey(auth *authenticate.Authenticate, params *subscribe.DeleteCustomerBillingKeyRequest) (*subscribe.DeleteCustomerBillingKeyResponse, error) {
	urls := []string{auth.APIUrl, URLSubscribe, URLCustomers, "/", params.GetCustomerUid()}

	token, err := auth.GetToken()
	if err != nil {
		return nil, err
	}

	isFirstQuery := true
	urls = append(urls, []string{util.GetQueryPrefix(&isFirstQuery), URLParamReason, urllib.PathEscape(params.GetReason())}...)
	urls = append(urls, []string{util.GetQueryPrefix(&isFirstQuery), URLParamRequester, urllib.PathEscape(params.GetRequester())}...)
	urlDeleleBillingkey := strings.Join(urls, "")

	res, err := util.Call(auth.Client, token, urlDeleleBillingkey, util.DELETE)
	if err != nil {
		return nil, err
	}

	deleteBillingKeyRes := subscribe.DeleteCustomerBillingKeyResponse{}
	err = protojson.Unmarshal(res, &deleteBillingKeyRes)
	if err != nil {
		return nil, err
	}

	return &deleteBillingKeyRes, nil
}

// GetBillingKeyByCustomer - GET /subscribe/customers/{customer_uid}
// 구매자 빌링키 조회하는 API
func GetBillingKeyByCustomer(auth *authenticate.Authenticate, params *subscribe.GetCustomerBillingKeyRequest) (*subscribe.GetCustomerBillingKeyResponse, error) {
	urls := util.GetJoinString(auth.APIUrl, URLSubscribe, URLCustomers, "/", params.GetCustomerUid())

	token, err := auth.GetToken()
	if err != nil {
		return nil, err
	}

	res, err := util.Call(auth.Client, token, urls, util.GET)
	if err != nil {
		return nil, err
	}

	getBillingKeyRes := subscribe.GetCustomerBillingKeyResponse{}
	err = protojson.Unmarshal(res, &getBillingKeyRes)
	if err != nil {
		return nil, err
	}

	return &getBillingKeyRes, nil
}

// InsertBillingKeyByCustomer - POST /subscribe/customers/{customer_uid}
// 구매자 빌링키 입력하는 API
func InsertBillingKeyByCustomer(auth *authenticate.Authenticate, params *subscribe.InsertCustomerBillingKeyRequest) (*subscribe.InsertCustomerBillingKeyResponse, error) {
	urls := util.GetJoinString(auth.APIUrl, URLSubscribe, URLCustomers, "/", params.GetCustomerUid())

	token, err := auth.GetToken()
	if err != nil {
		return nil, err
	}

	marshaler := protojson.MarshalOptions{
		UseProtoNames: true,
	}
	jsonBytes, err := marshaler.Marshal(params)
	if err != nil {
		return nil, err
	}

	res, err := util.CallWithJson(auth.Client, token, urls, util.POST, jsonBytes)
	if err != nil {
		return nil, err
	}

	insertBillingKeyRes := subscribe.InsertCustomerBillingKeyResponse{}
	err = protojson.Unmarshal(res, &insertBillingKeyRes)
	if err != nil {
		return nil, err
	}

	return &insertBillingKeyRes, nil
}

// GetPaymentsByCustomer - GET /subscribe/customers/{customer_uid}/payments
// 여러 빌링키를 한 번에 조회하는 API
func GetPaymentsByCustomer(auth *authenticate.Authenticate, params *subscribe.GetPaidByBillingKeyListRequest) (*subscribe.GetPaidByBillingKeyListResponse, error) {
	urls := []string{auth.APIUrl, URLSubscribe, URLCustomers, "/", params.GetCustomerUid(), URLPayments}

	token, err := auth.GetToken()
	if err != nil {
		return nil, err
	}

	isFirstQuery := true
	urls = append(urls, []string{util.GetQueryPrefix(&isFirstQuery), URLParamPage, strconv.Itoa(int(params.GetPage()))}...)
	urlGetBillingkeys := strings.Join(urls, "")

	res, err := util.Call(auth.Client, token, urlGetBillingkeys, util.GET)
	if err != nil {
		return nil, err
	}

	getPaymentRecordRes := subscribe.GetPaidByBillingKeyListResponse{}
	err = protojson.Unmarshal(res, &getPaymentRecordRes)
	if err != nil {
		return nil, err
	}

	return &getPaymentRecordRes, nil
}

// GetScheduledPaymentByCustomerUID - GET /subscribe/customers/{customer_uid}/schedules
// 예약한 결제 내역을 가져옵니다
func GetScheduledPaymentByCustomerUID(auth *authenticate.Authenticate, params *subscribe_dup.GetPaymentScheduleByCustomerRequest) (*subscribe_dup.GetPaymentScheduleByCustomerResponse, error) {
	urls := []string{auth.APIUrl, URLSubscribe, URLCustomers, "/", params.GetCustomerUid(), URLSchedules}

	isFirstQuery := true
	urls = append(urls, []string{util.GetQueryPrefix(&isFirstQuery), URLParamPage, strconv.Itoa(int(params.GetPage()))}...)
	urls = append(urls, []string{util.GetQueryPrefix(&isFirstQuery), URLParamFrom, strconv.Itoa(int(params.GetFrom()))}...)
	urls = append(urls, []string{util.GetQueryPrefix(&isFirstQuery), URLParamTo, strconv.Itoa(int(params.GetTo()))}...)
	urls = append(urls, []string{util.GetQueryPrefix(&isFirstQuery), URLParamScheduleStatus, params.GetScheduleStatus()}...)
	urlGetSchedule := strings.Join(urls, "")

	token, err := auth.GetToken()
	if err != nil {
		return nil, err
	}

	res, err := util.Call(auth.Client, token, urlGetSchedule, util.GET)
	if err != nil {
		return nil, err
	}

	scheduleRes := subscribe_dup.GetPaymentScheduleByCustomerResponse{}
	err = protojson.Unmarshal(res, &scheduleRes)
	if err != nil {
		return nil, err
	}

	return &scheduleRes, nil
}
