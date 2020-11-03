package subscribe

import (
	"net/http"
	"strconv"
	"strings"

	"google.golang.org/protobuf/encoding/protojson"

	"github.com/iamport/go-iamport/util"

	"github.com/iamport/interface/gen_src/go/v1/subscribe"
)

const (
	URLSubscribe  = "/subscribe"
	URLPayments   = "/payments"
	URLOnetime    = "/onetime"
	URLAgain      = "/again"
	URLSchedule   = "/schedule"
	URLUnschedule = "/unschedule"
	URLCustomers  = "/customers"

	URLParamPage = "page="
	URLParamFrom = "from="
	URLParamTo   = "to="
)

// Onetime - POST /subscribe/payments/onetime
// 구매자로부터 별도의 인증과정을 거치지 않고, 카드정보만으로 결제를 진행하는 API입니다(아임포트 javascript가 필요없습니다).
// customer_uid를 전달해주시면 결제 후 다음 번 결제를 위해 성공된 결제에 사용된 빌링키를 저장해두게되고, customer_uid가 없는 경우 저장되지 않습니다.
// 동일한 merchant_uid는 재사용이 불가능하며 고유한 값을 전달해주셔야 합니다.
// 빌링키 저장 시, buyer_email, buyer_name 등의 정보는 customer 부가정보인 customer_email, customer_name 등으로 함께 저장됩니다.
// /subscribe/customers/{customer_uid} 참조
func Onetime(client *http.Client, apiDomain string, token string, params *subscribe.OnetimePaymentRequest) (*subscribe.OnetimePaymentResponse, error) {
	url := util.GetJoinString(apiDomain, URLSubscribe, URLPayments, URLOnetime)

	marshaler := protojson.MarshalOptions{
		UseProtoNames: true,
	}
	jsonBytes, err := marshaler.Marshal(params)
	if err != nil {
		return nil, err
	}

	res, err := util.CallWithForm(client, token, url, util.POST, jsonBytes)
	if err != nil {
		return nil, err
	}

	onetimeRes := subscribe.OnetimePaymentResponse{}
	err = protojson.Unmarshal(res, &onetimeRes)
	if err != nil {
		return nil, err
	}

	return &onetimeRes, nil
}

// Again - POST /subscribe/payments/again
// 저장된 빌링키로 재결제를 하는 경우 사용됩니다. /subscribe/payments/onetime 또는 /subscribe/customers/{customer_uid} 로 등록된 빌링키가 있을 때 매칭되는 customer_uid로 재결제를 진행할 수 있습니다.
func Again(client *http.Client, apiDomain string, token string, params *subscribe.AgainPaymentRequest) (*subscribe.AgainPaymentResponse, error) {
	url := util.GetJoinString(apiDomain, URLSubscribe, URLPayments, URLAgain)

	marshaler := protojson.MarshalOptions{
		UseProtoNames: true,
	}
	jsonBytes, err := marshaler.Marshal(params)
	if err != nil {
		return nil, err
	}

	res, err := util.CallWithForm(client, token, url, util.POST, jsonBytes)
	if err != nil {
		return nil, err
	}

	againRes := subscribe.AgainPaymentResponse{}
	err = protojson.Unmarshal(res, &againRes)
	if err != nil {
		return nil, err
	}

	return &againRes, nil
}

// Schedule payemnt - POST /subscribe/payments/schedule
// 지정된 스케줄에 결제를 예약합니다
func Schedule(client *http.Client, apiDomain string, token string, params *subscribe.SchedulePayemntRequest) (*subscribe.SchedulePaymentResponse, error) {
	url := util.GetJoinString(apiDomain, URLSubscribe, URLPayments, URLSchedule)

	marshaler := protojson.MarshalOptions{
		UseProtoNames: true,
	}
	jsonBytes, err := marshaler.Marshal(params)
	if err != nil {
		return nil, err
	}

	res, err := util.CallWithJson(client, token, url, util.POST, jsonBytes)
	if err != nil {
		return nil, err
	}

	scheduleRes := subscribe.SchedulePaymentResponse{}
	err = protojson.Unmarshal(res, &scheduleRes)
	if err != nil {
		return nil, err
	}

	return &scheduleRes, nil
}

// Unschedule payemnt - POST /subscribe/payments/unschedule
// 예약한 결제를 취소합니다
func Unschedule(client *http.Client, apiDomain string, token string, params *subscribe.UnschedulePaymentRequest) (*subscribe.UnschedulePaymentResponse, error) {
	url := util.GetJoinString(apiDomain, URLSubscribe, URLPayments, URLUnschedule)

	marshaler := protojson.MarshalOptions{
		UseProtoNames: true,
	}
	jsonBytes, err := marshaler.Marshal(params)
	if err != nil {
		return nil, err
	}

	res, err := util.CallWithJson(client, token, url, util.POST, jsonBytes)
	if err != nil {
		return nil, err
	}

	unscheduleRes := subscribe.UnschedulePaymentResponse{}
	err = protojson.Unmarshal(res, &unscheduleRes)
	if err != nil {
		return nil, err
	}

	return &unscheduleRes, nil
}

// GetScheduledPaymentByMerchantUID - GET /subscribe/payments/schedule/{merchant_uid}
// 예약한 결제 내역을 가져옵니다
func GetScheduledPaymentByMerchantUID(client *http.Client, apiDomain string, token string, params *subscribe.GetPaymentScheduleRequest) (*subscribe.GetPaymentScheduleResponse, error) {
	url := util.GetJoinString(apiDomain, URLSubscribe, URLPayments, URLSchedule, "/", params.GetMerchantUid())

	res, err := util.Call(client, token, url, util.GET)
	if err != nil {
		return nil, err
	}

	scheduleRes := subscribe.GetPaymentScheduleResponse{}
	err = protojson.Unmarshal(res, &scheduleRes)
	if err != nil {
		return nil, err
	}

	return &scheduleRes, nil
}

// GetScheduledPaymentByCustomerUID - GET /subscribe/payments/schedule/customers/{merchant_uid}
// 예약한 결제 내역을 가져옵니다
func GetScheduledPaymentByCustomerUID(client *http.Client, apiDomain string, token string, params *subscribe.GetPaymentScheduleByCustomerRequest) (*subscribe.GetPaymentScheduleByCustomerResponse, error) {
	urls := []string{apiDomain, URLSubscribe, URLPayments, URLSchedule, URLCustomers, "/", params.GetCustomerUid()}

	isFirstQuery := true
	urls = append(urls, []string{util.GetQueryPrefix(&isFirstQuery), URLParamPage, strconv.Itoa(int(params.GetPage()))}...)
	urls = append(urls, []string{util.GetQueryPrefix(&isFirstQuery), URLParamFrom, strconv.Itoa(int(params.GetFrom()))}...)
	urls = append(urls, []string{util.GetQueryPrefix(&isFirstQuery), URLParamTo, strconv.Itoa(int(params.GetTo()))}...)
	urlGetSchedule := strings.Join(urls, "")

	res, err := util.Call(client, token, urlGetSchedule, util.GET)
	if err != nil {
		return nil, err
	}

	scheduleRes := subscribe.GetPaymentScheduleByCustomerResponse{}
	err = protojson.Unmarshal(res, &scheduleRes)
	if err != nil {
		return nil, err
	}

	return &scheduleRes, nil
}
