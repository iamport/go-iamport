package subscribe_customer

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	subscribe_dup "github.com/iamport/interface/gen_src/go/v1/subscribe"
	subscribe "github.com/iamport/interface/gen_src/go/v1/subscribe_customers"

	"github.com/iamport/go-iamport/authenticate"
	"github.com/iamport/go-iamport/util"
)

const (
	TMerchantUID    = "merchant_"
	TAmount         = 1000
	TName           = "아임포트 GO SDK 테스트"
	TBuyerEmail     = "example@example.com"
	TBuyerTel       = "01012341234"
	TBuyerName      = "홍길동"
	TBirth          = "911125"              // 실제값 입력
	TPwd2Digit      = "11"                  // 실제값 입력
	TCardNumber     = "1111-2222-3333-4444" // 실제값 입력
	TExpiry         = "2025-08"             // 실제값 입력
	TCustomerUid    = "testcustomer_gosdk_"
	TodoCustomValue = ""
)

func TestGetMultipleBillingKeysByCustomer(t *testing.T) {
	auth := authenticate.GetMockBingBongAuthenticate()
	params := &subscribe.GetMultipleCustomerBillingKeyRequest{
		CustomerUid: []string{
			TCustomerUid + util.GetRandomString(5),
			TCustomerUid + util.GetRandomString(5),
		},
	}

	token, err := auth.GetToken()
	if err != nil {
		t.Error(err)
	}

	res, err := GetMultipleBillingKeysByCustomer(auth.Client, auth.APIUrl, token, params)
	fmt.Println(res)
	assert.Error(t, err)
}

func TestDeleteBillingKey(t *testing.T) {
	auth := authenticate.GetMockBingBongAuthenticate()
	params := &subscribe.DeleteCustomerBillingKeyRequest{
		CustomerUid: TCustomerUid + util.GetRandomString(5),
		Reason:      "just test",
		Requester:   "you",
	}

	token, err := auth.GetToken()
	if err != nil {
		t.Error(err)
	}

	res, err := DeleteBillingKey(auth.Client, auth.APIUrl, token, params)
	fmt.Println(res)
	assert.NoError(t, err)
}

func TestGetBillingKeyByCustomer(t *testing.T) {
	auth := authenticate.GetMockBingBongAuthenticate()
	params := &subscribe.GetCustomerBillingKeyRequest{
		CustomerUid: TCustomerUid + util.GetRandomString(5),
	}

	token, err := auth.GetToken()
	if err != nil {
		t.Error(err)
	}

	res, err := GetBillingKeyByCustomer(auth.Client, auth.APIUrl, token, params)
	fmt.Println(res)
	assert.NoError(t, err)
}

func TestInsertBillingKeyByCustomer(t *testing.T) {
	auth := authenticate.GetMockBingBongAuthenticate()
	params := &subscribe.InsertCustomerBillingKeyRequest{
		CustomerUid:   TCustomerUid + util.GetRandomString(5),
		CardNumber:    TCardNumber,
		Expiry:        TExpiry,
		Birth:         TBirth,
		Pwd_2Digit:    TPwd2Digit,
		CustomerName:  TBuyerName,
		CustomerTel:   TBuyerTel,
		CustomerEmail: TBuyerEmail,
	}

	token, err := auth.GetToken()
	if err != nil {
		t.Error(err)
	}

	res, err := InsertBillingKeyByCustomer(auth.Client, auth.APIUrl, token, params)
	fmt.Println(res)
	assert.NoError(t, err)
}

func TestGetPaymentsByCustomer(t *testing.T) {
	auth := authenticate.GetMockBingBongAuthenticate()
	params := &subscribe.GetPaidByBillingKeyListRequest{
		CustomerUid: TCustomerUid + util.GetRandomString(5),
		Page:        1,
	}

	token, err := auth.GetToken()
	if err != nil {
		t.Error(err)
	}

	res, err := GetPaymentsByCustomer(auth.Client, auth.APIUrl, token, params)
	fmt.Println(res)
	assert.NoError(t, err)
}

func TestGetScheduledPaymentByCustomerUID(t *testing.T) {
	auth := authenticate.GetMockBingBongAuthenticate()
	params := &subscribe_dup.GetPaymentScheduleByCustomerRequest{
		CustomerUid: TCustomerUid + util.GetRandomString(5),
		Page:        1,
		From:        int32(time.Now().Unix() - 100000),
		To:          int32(time.Now().Unix()),
	}

	token, err := auth.GetToken()
	if err != nil {
		t.Error(err)
	}

	res, err := GetScheduledPaymentByCustomerUID(auth.Client, auth.APIUrl, token, params)
	fmt.Println(res)
	assert.NoError(t, err)
}
