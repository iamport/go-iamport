package payment

import (
	"math/rand"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/iamport/interface/build/go/payment"

	"github.com/joowonyun/go-iamport/authenticate"
)

const (
	ImpUID375245484897 = "imp_375245484897"
	ImpUID619325647049 = "imp_619325647049"

	Merchant1600132246284 = "merchant_1600132246284"
)

func TestGetByImpUID(t *testing.T) {
	auth := authenticate.GetMockBaseAuthenticate()

	params := &payment.PaymentRequest{
		ImpUid: ImpUID375245484897,
	}

	payment, err := GetByImpUID(auth, params)
	assert.NoError(t, err)

	checkPaymentImp375245484897(t, payment.Response)
}

func TestGetByImpUIDs(t *testing.T) {
	auth := authenticate.GetMockBaseAuthenticate()

	params := &payment.PaymentsRequest{
		ImpUid: []string{ImpUID375245484897, ImpUID619325647049},
	}

	payments, err := GetByImpUIDs(auth, params)
	assert.NoError(t, err)

	assert.Equal(t, 2, len(payments.Response))
	checkPaymentImp375245484897(t, payments.Response[0])
	checkPaymentImpUID619325647049(t, payments.Response[1])
}

func TestGetByMerchantUID(t *testing.T) {
	auth := authenticate.GetMockBaseAuthenticate()

	params := &payment.PaymentMerchantUidRequest{
		MerchantUid: Merchant1600132246284,
	}

	payment, err := GetByMerchantUID(auth, params)
	assert.NoError(t, err)

	checkPaymentImp375245484897(t, payment.Response)
}

func TestGetByMerchantUIDWithCorrectStatus(t *testing.T) {
	auth := authenticate.GetMockBaseAuthenticate()

	params := &payment.PaymentMerchantUidRequest{
		MerchantUid: Merchant1600132246284,
		Status:      "paid",
	}

	payment, err := GetByMerchantUID(auth, params)
	assert.NoError(t, err)

	assert.NotNil(t, payment)
}

func TestGetByMerchantUIDWithIncorrectStatus(t *testing.T) {
	auth := authenticate.GetMockBaseAuthenticate()

	params := &payment.PaymentMerchantUidRequest{
		MerchantUid: Merchant1600132246284,
		Status:      "ready",
	}

	payment, err := GetByMerchantUID(auth, params)
	assert.Error(t, err)
	assert.Nil(t, payment)
}

func TestGetByMerchantUIDs(t *testing.T) {
	auth := authenticate.GetMockBaseAuthenticate()

	params := &payment.PaymentsMerchantUidRequest{
		MerchantUid: Merchant1600132246284,
	}

	payment, err := GetByMerchantUIDs(auth, params)
	assert.NoError(t, err)

	checkPaymentImp375245484897(t, payment.Response.List[0])
}

func TestGetByMerchantUIDsWithCorrectStatus(t *testing.T) {
	auth := authenticate.GetMockBaseAuthenticate()

	params := &payment.PaymentsMerchantUidRequest{
		MerchantUid: Merchant1600132246284,
		Status:      "paid",
	}

	payment, err := GetByMerchantUIDs(auth, params)
	assert.NoError(t, err)

	checkPaymentImp375245484897(t, payment.Response.List[0])
}

func TestGetByMerchantUIDsWithInorrectStatus(t *testing.T) {
	auth := authenticate.GetMockBaseAuthenticate()

	params := &payment.PaymentsMerchantUidRequest{
		MerchantUid: Merchant1600132246284,
		Status:      "ready",
	}

	payment, err := GetByMerchantUIDs(auth, params)
	assert.Error(t, err)
	assert.Nil(t, payment)
}

func TestGetByMerchantUIDsWithCorrectPage(t *testing.T) {
	auth := authenticate.GetMockBaseAuthenticate()

	params := &payment.PaymentsMerchantUidRequest{
		MerchantUid: Merchant1600132246284,
		Page:        int32(1),
	}

	payment, err := GetByMerchantUIDs(auth, params)
	assert.NoError(t, err)
	checkPaymentImp375245484897(t, payment.Response.List[0])
}

func TestGetByMerchantUIDsWithIncorrectPage(t *testing.T) {
	auth := authenticate.GetMockBaseAuthenticate()

	params := &payment.PaymentsMerchantUidRequest{
		MerchantUid: Merchant1600132246284,
		Page:        int32(2),
	}

	payment, err := GetByMerchantUIDs(auth, params)
	assert.Error(t, err)
	assert.Nil(t, payment)
}

func TestGetByMerchantUIDsWithStatusAndPage(t *testing.T) {
	auth := authenticate.GetMockBaseAuthenticate()

	params := &payment.PaymentsMerchantUidRequest{
		MerchantUid: Merchant1600132246284,
		Status:      "paid",
		Page:        int32(1),
	}

	payment, err := GetByMerchantUIDs(auth, params)
	assert.NoError(t, err)
	checkPaymentImp375245484897(t, payment.Response.List[0])
}

func TestGetByMerchantUIDsWithStatusAndPageAndSorting(t *testing.T) {
	auth := authenticate.GetMockBaseAuthenticate()

	params := &payment.PaymentsMerchantUidRequest{
		MerchantUid: Merchant1600132246284,
		Status:      "paid",
		Page:        int32(1),
		Sorting:     "paid",
	}

	payment, err := GetByMerchantUIDs(auth, params)
	assert.NoError(t, err)
	checkPaymentImp375245484897(t, payment.Response.List[0])
}

func TestGetByStatus(t *testing.T) {
	auth := authenticate.GetMockBaseAuthenticate()

	params := &payment.PaymentStatusRequest{
		Status: "all",
	}

	payment, err := GetByStatus(auth, params)
	assert.NoError(t, err)
	assert.Equal(t, 20, len(payment.Response.List))
}

func TestGetByStatusWithLimit(t *testing.T) {
	auth := authenticate.GetMockBaseAuthenticate()

	params := &payment.PaymentStatusRequest{
		Status: "all",
		Limit:  10,
	}

	payment, err := GetByStatus(auth, params)
	assert.NoError(t, err)
	assert.Equal(t, 10, len(payment.Response.List))
}

func TestGetByStatusWithSpecificStatus(t *testing.T) {
	auth := authenticate.GetMockBaseAuthenticate()

	params := &payment.PaymentStatusRequest{
		Status: "paid",
		Limit:  10,
	}

	payment, err := GetByStatus(auth, params)
	assert.NoError(t, err)
	for _, pay := range payment.Response.GetList() {
		assert.Equal(t, "paid", pay.Status)
	}
}

func TestGetByStatusWithDate(t *testing.T) {
	auth := authenticate.GetMockBaseAuthenticate()

	params := &payment.PaymentStatusRequest{
		Status: "paid",
		Limit:  10,
		From:   int32(time.Date(2020, 9, 1, 0, 0, 0, 0, time.UTC).Unix()),
		To:     int32(time.Date(2020, 9, 2, 0, 0, 0, 0, time.UTC).Unix()),
	}

	payment, err := GetByStatus(auth, params)
	assert.NoError(t, err)
	for _, pay := range payment.Response.GetList() {
		assert.True(t, time.Unix(int64(pay.PaidAt), 0).Before(time.Date(2020, 9, 2, 0, 0, 0, 0, time.UTC)))
		assert.True(t, time.Unix(int64(pay.PaidAt), 0).After(time.Date(2020, 9, 1, 0, 0, 0, 0, time.UTC)))
	}
}

func TestGetByStatusWithSortPaidTime(t *testing.T) {
	auth := authenticate.GetMockBaseAuthenticate()

	params := &payment.PaymentStatusRequest{
		Status:  "paid",
		Limit:   10,
		Sorting: "-paid",
		From:    int32(time.Date(2020, 9, 1, 0, 0, 0, 0, time.UTC).Unix()),
		To:      int32(time.Date(2020, 9, 2, 0, 0, 0, 0, time.UTC).Unix()),
	}

	payment, err := GetByStatus(auth, params)
	assert.NoError(t, err)
	paidTime := time.Now()
	for _, pay := range payment.Response.GetList() {
		current := time.Unix(int64(pay.PaidAt), 0)
		assert.True(t, current.Before(paidTime))
		paidTime = current
	}
}

// TODO 테스트 데이터 필요 (KCP or Payco)
func xTestGetBalanceByImpUID(t *testing.T) {
	auth := authenticate.GetMockBaseAuthenticate()

	params := &payment.PaymentBalanceRequest{
		ImpUid: "imp_088621754304",
	}

	payment, err := GetBalanceByImpUID(auth, params)
	assert.NoError(t, err)
	assert.Equal(t, 123, payment.Response.Amount)
}

func TestPrepare(t *testing.T) {
	auth := authenticate.GetMockBaseAuthenticate()

	rand.Seed(time.Now().UnixNano())
	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVEWXYZabcdefghijklmnopqrstuvewxyz0123456789")
	var merchantBytes strings.Builder
	for i := 0; i < 20; i++ {
		merchantBytes.WriteRune(chars[rand.Intn(len(chars))])
	}

	merchantUID := merchantBytes.String()
	amount := rand.Intn(10000)

	params := &payment.PaymentPrepareRequest{
		MerchantUid: merchantUID,
		Amount:      float64(amount),
	}

	res, err := Prepare(auth, params)
	assert.NoError(t, err)
	assert.Equal(t, merchantUID, res.Response.MerchantUid)
	assert.Equal(t, amount, int(res.Response.Amount))

	params = &payment.PaymentPrepareRequest{
		MerchantUid: merchantUID,
	}

	res, err = GetPrepareByMerchantUID(auth, params)
	assert.NoError(t, err)
	assert.Equal(t, merchantUID, res.Response.MerchantUid)
	assert.Equal(t, amount, int(res.Response.Amount))
}

func checkPaymentImp375245484897(t *testing.T, p *payment.Payment) {
	assert.Equal(t, 14000, int(p.Amount))
	assert.Equal(t, "서울특별시 강남구 삼성동", p.BuyerAddr)
	assert.Equal(t, "iamport@siot.do", p.BuyerEmail)
	assert.Equal(t, "구매자이름", p.BuyerName)
	assert.Equal(t, "123-456", p.BuyerPostcode)
	assert.Equal(t, "010-1234-5678", p.BuyerTel)
	assert.False(t, p.CashReceiptIssued)
	assert.Equal(t, "pc", p.Channel)
	assert.Equal(t, "KRW", p.Currency)
	assert.False(t, p.Escrow)
	assert.Equal(t, 0, int(p.FailedAt))
	assert.Equal(t, "imp_375245484897", p.ImpUid)
	assert.Equal(t, "merchant_1600132246284", p.MerchantUid)
	assert.Equal(t, "주문명:결제테스트", p.Name)
	assert.Equal(t, 1600132785, int(p.PaidAt))
	assert.Equal(t, "vbank", p.PayMethod)
	assert.Equal(t, "INIpayTest", p.PgId)
	assert.Equal(t, "html5_inicis", p.PgProvider)
	assert.Equal(t, "StdpayVBNKINIpayTest20200915101120578600", p.PgTid)
	assert.Equal(t, "https://iniweb.inicis.com/DefaultWebApp/mall/cr/cm/mCmReceipt_head.jsp?noTid=StdpayVBNKINIpayTest20200915101120578600&noMethod=1", p.ReceiptUrl)
	assert.Equal(t, 1600132245, int(p.StartedAt))
	assert.Equal(t, "paid", p.Status)
	assert.Equal(t, "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.102 Safari/537.36", p.UserAgent)
	assert.Equal(t, "011", p.VbankCode)
	assert.Equal(t, 1602773999, int(p.VbankDate))
	assert.Equal(t, "（주）케이지이니시", p.VbankHolder)
	assert.Equal(t, 1600132280, int(p.VbankIssuedAt))
	assert.Equal(t, "농협중앙회", p.VbankName)
	assert.Equal(t, "79014407657133", p.VbankNum)
}

func checkPaymentImpUID619325647049(t *testing.T, p *payment.Payment) {
	assert.Equal(t, 200, int(p.Amount))
	assert.Equal(t, "34009097", p.ApplyNum)
	assert.Equal(t, "lepus2073@gmail.com", p.BuyerEmail)
	assert.Equal(t, "", p.BuyerName)
	assert.Equal(t, "", p.BuyerPostcode)
	assert.Equal(t, "", p.BuyerTel)
	assert.Equal(t, 0, int(p.CancelAmount))
	assert.Equal(t, 0, int(p.CancelledAt))
	assert.Equal(t, "361", p.CardCode)
	assert.Equal(t, "BC카드", p.CardName)
	assert.Equal(t, "910023*********3", p.CardNumber)
	assert.Equal(t, 0, int(p.CardQuota))
	assert.Equal(t, 0, int(p.CardType))
	assert.False(t, p.CashReceiptIssued)
	assert.Equal(t, "mobile", p.Channel)
	assert.Equal(t, "KRW", p.Currency)
	assert.False(t, p.Escrow)
	assert.Equal(t, 0, int(p.FailedAt))
	assert.Equal(t, "imp_619325647049", p.ImpUid)
	assert.Equal(t, "ifOp5OHWV6fuoTOqlCLW9b7WLnr2_16000963253", p.MerchantUid)
	assert.Equal(t, "XXXX 상담 비용", p.Name)
	assert.Equal(t, 1600096351, int(p.PaidAt))
	assert.Equal(t, "card", p.PayMethod)
	assert.Equal(t, "INIpayTest", p.PgId)
	assert.Equal(t, "html5_inicis", p.PgProvider)
	assert.Equal(t, "INIMX_ISP_INIpayTest20200915001231335553", p.PgTid)
	assert.Equal(t, "https://iniweb.inicis.com/DefaultWebApp/mall/cr/cm/mCmReceipt_head.jsp?noTid=INIMX_ISP_INIpayTest20200915001231335553&noMethod=1", p.ReceiptUrl)
	assert.Equal(t, 1600096325, int(p.StartedAt))
	assert.Equal(t, "paid", p.Status)
	assert.Equal(t, "Mozilla/5.0 (iPhone; CPU iPhone OS 13_7 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148", p.UserAgent)
	assert.Equal(t, 0, int(p.VbankDate))
}
