package subscribe

import (
	"strings"
	"testing"

	"github.com/iamport/go-iamport/util"

	"github.com/stretchr/testify/assert"

	"github.com/iamport/interface/gen_src/go/v1/subscribe"

	"github.com/iamport/go-iamport/authenticate"
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
	TodoCustomValue = ""
)

// 카드 정보 실제 값 입력시 테스트 가능
func xTestOneTime(t *testing.T) {

	auth := authenticate.GetMockBingBongAuthenticate()

	params := &subscribe.OnetimePaymentRequest{
		MerchantUid: TMerchantUID + util.GetRandomString(20), // require
		Amount:      TAmount,                                 // require
		CardNumber:  TCardNumber,                             // require
		Expiry:      TExpiry,                                 // require
		Birth:       TBirth,                                  // require
		Pwd_2Digit:  TPwd2Digit,
		Name:        TName,
		BuyerName:   TBuyerName,
		BuyerEmail:  TBuyerEmail,
		BuyerTel:    TBuyerTel,
		// TaxFree:                float64(0),
		// CustomerUid:            TodoCustomValue,
		// Pg:                     TodoCustomValue,
		// BuyerAddr:              TodoCustomValue,
		// BuyerPostcode:          TodoCustomValue,
		// CardQuota:              0,
		// InterestFreeByMerchant: false,
		// CustomData:             TodoCustomValue,
		// NoticeUrl:              TodoCustomValue,
	}

	onetimeRes, err := Onetime(auth, params)
	assert.NoError(t, err)
	checkOnetimePaymentParameter(t, params, onetimeRes)
}

// 일반 카드 결제 파라미터 체크
func checkOnetimePaymentParameter(t *testing.T, req *subscribe.OnetimePaymentRequest, onetimeRes *subscribe.OnetimePaymentResponse) {
	assert.NotEqual(t, nil, onetimeRes.Code)
	assert.Equal(t, int32(0), onetimeRes.Code)

	res := onetimeRes.Response
	assert.Equal(t, req.GetMerchantUid(), res.GetMerchantUid())
	assert.Equal(t, req.GetAmount(), float64(res.GetAmount()))

	requstCardNum := req.GetCardNumber()
	requstCardNum = strings.Replace(requstCardNum, "-", "", -1)
	responseCardNum := res.GetCardNumber()
	assert.True(t, strings.HasPrefix(responseCardNum, requstCardNum[:6]))
	assert.True(t, strings.HasSuffix(responseCardNum, requstCardNum[len(requstCardNum)-4:]))

	assert.Equal(t, req.GetName(), res.GetName())
	assert.Equal(t, req.GetBuyerName(), res.GetBuyerName())
	assert.Equal(t, req.GetBuyerEmail(), res.GetBuyerEmail())
	assert.Equal(t, req.GetBuyerTel(), res.GetBuyerTel())
	assert.Equal(t, req.GetCustomerUid(), res.GetCustomerUid())

	// TODO 기타 파라미터?
}

/*
{
  "code": 0,
  "message": null,
  "response": {
    "amount": 1000,
    "apply_num": "63889323",
    "buyer_email": "example@example.com",
    "buyer_name": "홍길동",
    "buyer_tel": "01012341234",
    "card_code": "361",
    "card_name": "BC카드",
    "card_number": "95404900****1569",
    "card_type": 1,
    "channel": "api",
    "currency": "KRW",
    "imp_uid": "imps_157956638714",
    "merchant_uid": "merchant_901w6xDqDHUoR6r8FJ5R",
    "name": "아임포트 GO SDK 테스트",
    "paid_at": 1603417956,
    "pay_method": "card",
    "pg_id": "nictest04m",
    "pg_provider": "nice",
    "pg_tid": "nictest04m01162010231052355996",
    "receipt_url": "https://npg.nicepay.co.kr/issue/IssueLoader.do?TID=nictest04m01162010231052355996&type=0&InnerWin=Y",
    "started_at": 1603417956,
    "status": "paid",
    "user_agent": "Go-http-client/1.1"
  }
}
*/
