package iamport

import (
	"testing"
	"time"

	"github.com/iamport/go-iamport/authenticate"
	"github.com/iamport/go-iamport/util"
	"github.com/stretchr/testify/assert"
)

func TestGetMultipleBillingKeysByCustomer(t *testing.T) {
	iamport, err := NewIamport(authenticate.BaseURL, authenticate.RestApiKey, authenticate.RestApiSecret)
	assert.NoError(t, err)

	schedules, err := iamport.GetMultipleBillingKeysByCustomer(
		[]string{TCustomerUid + util.GetRandomString(5), TCustomerUid + util.GetRandomString(5)},
	)
	assert.Equal(t, len(schedules), 0)
}

func TestDeleteBillingKey(t *testing.T) {
	iamport, err := NewIamport(authenticate.BaseURL, authenticate.RestApiKey, authenticate.RestApiSecret)
	assert.NoError(t, err)

	res, err := iamport.DeleteBillingKey(TCustomerUid+util.GetRandomString(5), "just test", "")
	assert.Nil(t, res)
	assert.Contains(t, err.Error(), "등록된 정보를 찾을 수 없습니다.")
}

func TestGetBillingKeyByCustomer(t *testing.T) {
	iamport, err := NewIamport(authenticate.BaseURL, authenticate.RestApiKey, authenticate.RestApiSecret)
	assert.NoError(t, err)

	res, err := iamport.GetBillingKeyByCustomer(TCustomerUid + util.GetRandomString(5))
	assert.Nil(t, res)
	assert.Contains(t, err.Error(), "등록된 정보를 찾을 수 없습니다.")
}

func TestInsertBillingKeyByCustomer(t *testing.T) {
	iamport, err := NewIamport(authenticate.BaseURL, authenticate.RestApiKey, authenticate.RestApiSecret)
	assert.NoError(t, err)

	res, err := iamport.InsertBillingKeyByCustomer(
		TCustomerUid+util.GetRandomString(5),
		"",
		TCardNumber, TExpiry, TBirth, TPwd2Digit,
		TBuyerName, TBuyerTel, TBuyerEmail, "", "",
	)
	assert.Nil(t, res)
	assert.Contains(t, err.Error(), "카드번호")
}

func TestGetPaymentsByCustomer(t *testing.T) {
	iamport, err := NewIamport(authenticate.BaseURL, authenticate.RestApiKey, authenticate.RestApiSecret)
	assert.NoError(t, err)

	res, err := iamport.GetPaymentsByCustomer(TCustomerUid+util.GetRandomString(5), 0)
	assert.Nil(t, err)
	assert.Equal(t, len(res.GetList()), 0)
}

func TestGetScheduledPaymentListByCustomerUID(t *testing.T) {
	iamport, err := NewIamport(authenticate.BaseURL, authenticate.RestApiKey, authenticate.RestApiSecret)
	assert.NoError(t, err)

	schedules, err := iamport.GetScheduledPaymentListByCustomerUID(
		TCustomerUid+util.GetRandomString(5),
		0,
		int32(time.Now().Unix()-100000),
		int32(time.Now().Unix()),
		"",
	)
	assert.Nil(t, err)
	assert.Equal(t, len(schedules.GetList()), 0)
}
