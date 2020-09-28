package iamport

import (
	"math/rand"
	"strings"
	"testing"
	"time"

	"github.com/joowonyun/go-iamport/authenticate"
	"github.com/stretchr/testify/assert"
)

func TestGetPaymentImpUID(t *testing.T) {
	iamport, err := NewIamport(authenticate.BaseURL, authenticate.RestApiKey, authenticate.RestApiSecret)
	assert.NoError(t, err)

	payment, err := iamport.GetPaymentImpUID("imp_375245484897")
	assert.NoError(t, err)
	assert.NotNil(t, payment)
}

func TestGetPaymentImpUIDWithoutImpUid(t *testing.T) {
	iamport, err := NewIamport(authenticate.BaseURL, authenticate.RestApiKey, authenticate.RestApiSecret)
	assert.NoError(t, err)

	payment, err := iamport.GetPaymentImpUID("")
	assert.Error(t, err)
	assert.Nil(t, payment)
}

func TestGetPaymentImpUIDs(t *testing.T) {
	iamport, err := NewIamport(authenticate.BaseURL, authenticate.RestApiKey, authenticate.RestApiSecret)
	assert.NoError(t, err)

	payment, err := iamport.GetPaymentsImpUIDs([]string{"imp_375245484897"})
	assert.NoError(t, err)
	assert.NotNil(t, payment)
}

func TestGetPaymentsImpUIDsWithoutImpUIds(t *testing.T) {
	iamport, err := NewIamport(authenticate.BaseURL, authenticate.RestApiKey, authenticate.RestApiSecret)
	assert.NoError(t, err)

	payment, err := iamport.GetPaymentsImpUIDs([]string{})
	assert.Error(t, err)
	assert.Nil(t, payment)
}

func TestGetPaymentMerchantUID(t *testing.T) {
	iamport, err := NewIamport(authenticate.BaseURL, authenticate.RestApiKey, authenticate.RestApiSecret)
	assert.NoError(t, err)

	payment, err := iamport.GetPaymentMerchantUID("merchant_1600132246284", "", "")
	assert.NoError(t, err)
	assert.NotNil(t, payment)
}

func TestGetPaymentMerchantUIDWithoutMerchantUID(t *testing.T) {
	iamport, err := NewIamport(authenticate.BaseURL, authenticate.RestApiKey, authenticate.RestApiSecret)
	assert.NoError(t, err)

	payment, err := iamport.GetPaymentMerchantUID("", "", "")
	assert.Error(t, err)
	assert.Nil(t, payment)
}

func TestGetPaymentMerchantUIDInvalidStatus(t *testing.T) {
	iamport, err := NewIamport(authenticate.BaseURL, authenticate.RestApiKey, authenticate.RestApiSecret)
	assert.NoError(t, err)

	payment, err := iamport.GetPaymentMerchantUID("merchant_1600132246284", "error", "")
	assert.Error(t, err)
	assert.Nil(t, payment)
}

func TestGetPaymentMerchantUIDInvalidSort(t *testing.T) {
	iamport, err := NewIamport(authenticate.BaseURL, authenticate.RestApiKey, authenticate.RestApiSecret)
	assert.NoError(t, err)

	payment, err := iamport.GetPaymentMerchantUID("merchant_1600132246284", "", "error")
	assert.Error(t, err)
	assert.Nil(t, payment)
}

func TestGetPaymentsMerchantUID(t *testing.T) {
	iamport, err := NewIamport(authenticate.BaseURL, authenticate.RestApiKey, authenticate.RestApiSecret)
	assert.NoError(t, err)

	payment, err := iamport.GetPaymentsMerchantUID("merchant_1600132246284", "", "", 0)
	assert.NoError(t, err)
	assert.NotNil(t, payment)
}

func TestGetPaymentsMerchantUIDWithoutMerchantUID(t *testing.T) {
	iamport, err := NewIamport(authenticate.BaseURL, authenticate.RestApiKey, authenticate.RestApiSecret)
	assert.NoError(t, err)

	payment, err := iamport.GetPaymentsMerchantUID("", "", "", 0)
	assert.Error(t, err)
	assert.Nil(t, payment)
}

func TestGetPaymentsMerchantUIDInvalidStatus(t *testing.T) {
	iamport, err := NewIamport(authenticate.BaseURL, authenticate.RestApiKey, authenticate.RestApiSecret)
	assert.NoError(t, err)

	payment, err := iamport.GetPaymentsMerchantUID("merchant_1600132246284", "error", "", 0)
	assert.Error(t, err)
	assert.Nil(t, payment)
}

func TestGetPaymentsMerchantUIDInvalidSort(t *testing.T) {
	iamport, err := NewIamport(authenticate.BaseURL, authenticate.RestApiKey, authenticate.RestApiSecret)
	assert.NoError(t, err)

	payment, err := iamport.GetPaymentsMerchantUID("merchant_1600132246284", "", "error", 0)
	assert.Error(t, err)
	assert.Nil(t, payment)
}

func TestGetPaymentsMerchantUIDInvalidPage(t *testing.T) {
	iamport, err := NewIamport(authenticate.BaseURL, authenticate.RestApiKey, authenticate.RestApiSecret)
	assert.NoError(t, err)

	payment, err := iamport.GetPaymentsMerchantUID("merchant_1600132246284", "", "", -1)
	assert.Error(t, err)
	assert.Nil(t, payment)
}

func TestGetPaymentsStatus(t *testing.T) {
	iamport, err := NewIamport(authenticate.BaseURL, authenticate.RestApiKey, authenticate.RestApiSecret)
	assert.NoError(t, err)

	utc, err := time.LoadLocation("UTC")
	assert.NoError(t, err)
	payment, err := iamport.GetPaymentsStatus("", 0, 0, time.Date(2020, 9, 1, 0, 0, 0, 0, utc), time.Date(2020, 9, 2, 0, 0, 0, 0, utc), "")
	assert.NoError(t, err)
	assert.NotNil(t, payment)
}

func TestGetPaymentsStatusWithInvalidStatus(t *testing.T) {
	iamport, err := NewIamport(authenticate.BaseURL, authenticate.RestApiKey, authenticate.RestApiSecret)
	assert.NoError(t, err)

	utc, err := time.LoadLocation("UTC")
	assert.NoError(t, err)
	payment, err := iamport.GetPaymentsStatus("error", 0, 0, time.Date(2020, 9, 1, 0, 0, 0, 0, utc), time.Date(2020, 9, 2, 0, 0, 0, 0, utc), "")
	assert.Error(t, err)
	assert.Nil(t, payment)
}

func TestGetPaymentsStatusWithInvalidPage(t *testing.T) {
	iamport, err := NewIamport(authenticate.BaseURL, authenticate.RestApiKey, authenticate.RestApiSecret)
	assert.NoError(t, err)

	utc, err := time.LoadLocation("UTC")
	assert.NoError(t, err)
	payment, err := iamport.GetPaymentsStatus("", -1, 0, time.Date(2020, 9, 1, 0, 0, 0, 0, utc), time.Date(2020, 9, 2, 0, 0, 0, 0, utc), "")
	assert.Error(t, err)
	assert.Nil(t, payment)
}

func TestGetPaymentsStatusWithInvalidLimit(t *testing.T) {
	iamport, err := NewIamport(authenticate.BaseURL, authenticate.RestApiKey, authenticate.RestApiSecret)
	assert.NoError(t, err)

	utc, err := time.LoadLocation("UTC")
	assert.NoError(t, err)
	payment, err := iamport.GetPaymentsStatus("", 0, -1, time.Date(2020, 9, 1, 0, 0, 0, 0, utc), time.Date(2020, 9, 2, 0, 0, 0, 0, utc), "")
	assert.Error(t, err)
	assert.Nil(t, payment)
}

func TestGetPaymentsStatusWithInvalidTime(t *testing.T) {
	iamport, err := NewIamport(authenticate.BaseURL, authenticate.RestApiKey, authenticate.RestApiSecret)
	assert.NoError(t, err)

	utc, err := time.LoadLocation("UTC")
	assert.NoError(t, err)
	payment, err := iamport.GetPaymentsStatus("", 0, 0, time.Date(2020, 10, 1, 0, 0, 0, 0, utc), time.Date(2020, 9, 2, 0, 0, 0, 0, utc), "")
	assert.Error(t, err)
	assert.Nil(t, payment)
}

func TestGetPaymentsStatusWithMoreThan3Month(t *testing.T) {
	iamport, err := NewIamport(authenticate.BaseURL, authenticate.RestApiKey, authenticate.RestApiSecret)
	assert.NoError(t, err)

	utc, err := time.LoadLocation("UTC")
	assert.NoError(t, err)
	payment, err := iamport.GetPaymentsStatus("", 0, 0, time.Date(2020, 1, 1, 0, 0, 0, 0, utc), time.Date(2020, 10, 2, 0, 0, 0, 0, utc), "")
	assert.Error(t, err)
	assert.Nil(t, payment)
}

// TODO 테스트 데이터 필요 (KCP or Payco)
func xTestGetPaymentBalanceImpUID(t *testing.T) {
	iamport, err := NewIamport(authenticate.BaseURL, authenticate.RestApiKey, authenticate.RestApiSecret)
	assert.NoError(t, err)

	payment, err := iamport.GetPaymentBalanceImpUID("imp_088621754304")
	assert.NoError(t, err)
	assert.NotNil(t, payment)
}

func TestPreparePayment(t *testing.T) {
	iamport, err := NewIamport(authenticate.BaseURL, authenticate.RestApiKey, authenticate.RestApiSecret)
	assert.NoError(t, err)

	rand.Seed(time.Now().UnixNano())
	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVEWXYZabcdefghijklmnopqrstuvewxyz0123456789")
	var merchantBytes strings.Builder
	for i := 0; i < 20; i++ {
		merchantBytes.WriteRune(chars[rand.Intn(len(chars))])
	}

	merchantUID := merchantBytes.String()
	amount := rand.Intn(10000)

	payment, err := iamport.PreparePayment(merchantUID, float64(amount))
	assert.NoError(t, err)
	assert.NotNil(t, payment)

	res, err := iamport.GetPreparePayment(merchantUID)
	assert.NoError(t, err)
	assert.NotNil(t, res)
}

func TestPreparePaymentWithoutMerchantUID(t *testing.T) {
	iamport, err := NewIamport(authenticate.BaseURL, authenticate.RestApiKey, authenticate.RestApiSecret)
	assert.NoError(t, err)

	payment, err := iamport.PreparePayment("", float64(1))
	assert.Error(t, err)
	assert.Nil(t, payment)
}

func TestPreparePaymentWithInvalidAmount(t *testing.T) {
	iamport, err := NewIamport(authenticate.BaseURL, authenticate.RestApiKey, authenticate.RestApiSecret)
	assert.NoError(t, err)

	payment, err := iamport.PreparePayment("abc", float64(-1))
	assert.Error(t, err)
	assert.Nil(t, payment)
}

func TestGetPreparePaymentWithoutMerchantUID(t *testing.T) {
	iamport, err := NewIamport(authenticate.BaseURL, authenticate.RestApiKey, authenticate.RestApiSecret)
	assert.NoError(t, err)

	payment, err := iamport.GetPreparePayment("")
	assert.Error(t, err)
	assert.Nil(t, payment)
}
