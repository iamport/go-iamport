package iamport

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	TypeSubscribe "github.com/iamport/interface/gen_src/go/v1/subscribe"

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

var TScheduleParams = &TypeSubscribe.PaymentScheduleParam{
	MerchantUid:   TMerchantUID,
	ScheduleAt:    int32(time.Now().Unix()) + 10000000,
	Amount:        TAmount,
	TaxFree:       0,
	Name:          TName,
	BuyerName:     TBuyerName,
	BuyerEmail:    TBuyerEmail,
	BuyerTel:      TBuyerTel,
	BuyerAddr:     "",
	BuyerPostcode: "",
}

func TestOnetimePayment(t *testing.T) {
	iamport, err := NewIamport(authenticate.BaseURL, authenticate.RestApiKey, authenticate.RestApiSecret)
	assert.NoError(t, err)

	payment, err := iamport.OnetimePayment(
		TMerchantUID+time.Now().String(), TAmount, 0,
		TCardNumber, TExpiry, TBirth, TPwd2Digit,
		"", "",
		TName, TBuyerName, TBuyerEmail, TBuyerTel, "", "",
		0, false, "", "",
	)
	assert.Contains(t, err.Error(), "유효하지않은 카드번호를 입력하셨습니다.")
	assert.Nil(t, payment)
}

func TestAgainPayment(t *testing.T) {
	iamport, err := NewIamport(authenticate.BaseURL, authenticate.RestApiKey, authenticate.RestApiSecret)
	assert.NoError(t, err)

	payment, err := iamport.AgainPayment(
		TCustomerUid+time.Now().String(), TMerchantUID+time.Now().String(), TAmount, 0,
		TName,
		TBuyerName, TBuyerEmail, TBuyerTel, "", "",
		0, false, "", "",
	)
	assert.Contains(t, err.Error(), "등록되지 않은 구매자입니다.")
	assert.Nil(t, payment)
}

func TestSchedulePayment(t *testing.T) {
	iamport, err := NewIamport(authenticate.BaseURL, authenticate.RestApiKey, authenticate.RestApiSecret)
	assert.NoError(t, err)

	schedules, err := iamport.SchedulePayment(
		TCustomerUid+time.Now().String(), 0,
		TCardNumber, TExpiry, TBirth, TPwd2Digit,
		"", []*TypeSubscribe.PaymentScheduleParam{TScheduleParams},
	)
	assert.Contains(t, err.Error(), "유효하지않은 카드번호를 입력하셨습니다.")
	assert.Nil(t, schedules)
}

func TestUnschedulePayment(t *testing.T) {
	iamport, err := NewIamport(authenticate.BaseURL, authenticate.RestApiKey, authenticate.RestApiSecret)
	assert.NoError(t, err)

	schedules, err := iamport.UnschedulePayment(TCustomerUid+time.Now().String(), nil)
	assert.Contains(t, err.Error(), "취소할 예약결제 기록이 존재하지 않습니다.")
	assert.Nil(t, schedules)
}

func TestGetScheduledPaymentByMerchantUID(t *testing.T) {
	iamport, err := NewIamport(authenticate.BaseURL, authenticate.RestApiKey, authenticate.RestApiSecret)
	assert.NoError(t, err)

	schedules, err := iamport.GetScheduledPaymentByMerchantUID(TMerchantUID + time.Now().String())
	assert.Contains(t, err.Error(), "invalid imp_uid")
	assert.Nil(t, schedules)
}

func TestGetScheduledPaymentByCustomerUID(t *testing.T) {
	iamport, err := NewIamport(authenticate.BaseURL, authenticate.RestApiKey, authenticate.RestApiSecret)
	assert.NoError(t, err)

	schedules, err := iamport.GetScheduledPaymentByCustomerUID(
		TCustomerUid+util.GetRandomString(5),
		0,
		int32(time.Now().Unix()-100000),
		int32(time.Now().Unix()),
		"",
	)
	assert.Equal(t, len(schedules.GetList()), 0)
}
