package payment

import (
	"math/rand"
	"testing"
	"time"

	"github.com/iamport/go-iamport/util"

	"github.com/stretchr/testify/assert"

	"github.com/iamport/interface/gen_src/go/payment"

	"github.com/iamport/go-iamport/authenticate"
)

const (
	ImpUID785510843101 = "imp_785510843101"
	ImpUID338103167934 = "imp_338103167934"

	MerchantUIDORD20180131_0009728 = "ORD20180131-0009728"
)

func TestGetByImpUID(t *testing.T) {
	auth := authenticate.GetMockBaseAuthenticate()

	params := &payment.PaymentRequest{
		ImpUid: ImpUID785510843101,
	}

	payment, err := GetByImpUID(auth, params)
	assert.NoError(t, err)

	checkPaymentImp375245484897(t, payment.Response)
}

func TestGetByImpUIDs(t *testing.T) {
	auth := authenticate.GetMockBaseAuthenticate()

	params := &payment.PaymentsRequest{
		ImpUid: []string{ImpUID785510843101, ImpUID338103167934},
	}

	payments, err := GetByImpUIDs(auth, params)
	assert.NoError(t, err)

	assert.Equal(t, 2, len(payments.Response))
	checkPaymentImp375245484897(t, payments.Response[1])
	checkPaymentImpUID338103167934(t, payments.Response[0])
}

func TestGetByMerchantUID(t *testing.T) {
	auth := authenticate.GetMockBaseAuthenticate()

	params := &payment.PaymentMerchantUidRequest{
		MerchantUid: MerchantUIDORD20180131_0009728,
	}

	payment, err := GetByMerchantUID(auth, params)
	assert.NoError(t, err)

	checkPaymentImp375245484897(t, payment.Response)
}

func TestGetByMerchantUIDWithCorrectStatus(t *testing.T) {
	auth := authenticate.GetMockBaseAuthenticate()

	params := &payment.PaymentMerchantUidRequest{
		MerchantUid: MerchantUIDORD20180131_0009728,
		Status:      "paid",
	}

	payment, err := GetByMerchantUID(auth, params)
	assert.NoError(t, err)

	assert.NotNil(t, payment)
}

func TestGetByMerchantUIDWithIncorrectStatus(t *testing.T) {
	auth := authenticate.GetMockBaseAuthenticate()

	params := &payment.PaymentMerchantUidRequest{
		MerchantUid: MerchantUIDORD20180131_0009728,
		Status:      "123",
	}

	payment, err := GetByMerchantUID(auth, params)
	assert.Error(t, err)
	assert.Nil(t, payment)
}

func TestGetByMerchantUIDs(t *testing.T) {
	auth := authenticate.GetMockBaseAuthenticate()

	params := &payment.PaymentsMerchantUidRequest{
		MerchantUid: MerchantUIDORD20180131_0009728,
	}

	payment, err := GetByMerchantUIDs(auth, params)
	assert.NoError(t, err)

	checkPaymentImp375245484897(t, payment.Response.List[0])
}

func TestGetByMerchantUIDsWithCorrectStatus(t *testing.T) {
	auth := authenticate.GetMockBaseAuthenticate()

	params := &payment.PaymentsMerchantUidRequest{
		MerchantUid: MerchantUIDORD20180131_0009728,
		Status:      "paid",
	}

	payment, err := GetByMerchantUIDs(auth, params)
	assert.NoError(t, err)

	checkPaymentImp375245484897(t, payment.Response.List[0])
}

func TestGetByMerchantUIDsWithInorrectStatus(t *testing.T) {
	auth := authenticate.GetMockBaseAuthenticate()

	params := &payment.PaymentsMerchantUidRequest{
		MerchantUid: MerchantUIDORD20180131_0009728,
		Status:      "123",
	}

	payment, err := GetByMerchantUIDs(auth, params)
	assert.Error(t, err)
	assert.Nil(t, payment)
}

func TestGetByMerchantUIDsWithCorrectPage(t *testing.T) {
	auth := authenticate.GetMockBaseAuthenticate()

	params := &payment.PaymentsMerchantUidRequest{
		MerchantUid: MerchantUIDORD20180131_0009728,
		Page:        int32(1),
	}

	payment, err := GetByMerchantUIDs(auth, params)
	assert.NoError(t, err)
	checkPaymentImp375245484897(t, payment.Response.List[0])
}

func TestGetByMerchantUIDsWithIncorrectPage(t *testing.T) {
	auth := authenticate.GetMockBaseAuthenticate()

	params := &payment.PaymentsMerchantUidRequest{
		MerchantUid: MerchantUIDORD20180131_0009728,
		Page:        int32(2),
	}

	payment, err := GetByMerchantUIDs(auth, params)
	assert.Error(t, err)
	assert.Nil(t, payment)
}

func TestGetByMerchantUIDsWithStatusAndPage(t *testing.T) {
	auth := authenticate.GetMockBaseAuthenticate()

	params := &payment.PaymentsMerchantUidRequest{
		MerchantUid: MerchantUIDORD20180131_0009728,
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
		MerchantUid: MerchantUIDORD20180131_0009728,
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
	assert.Equal(t, 9, len(payment.Response.List))
}

func TestGetByStatusWithLimit(t *testing.T) {
	auth := authenticate.GetMockBaseAuthenticate()

	params := &payment.PaymentStatusRequest{
		Status: "all",
		Limit:  10,
	}

	payment, err := GetByStatus(auth, params)
	assert.NoError(t, err)
	assert.Equal(t, 9, len(payment.Response.List))
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

	merchantUID := util.GetRandomString(20)
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
	assert.Equal(t, 100, int(p.Amount))
	assert.Equal(t, "서울특별시 강남구 신사동", p.BuyerAddr)
	assert.Equal(t, "gildong@gmail.com", p.BuyerEmail)
	assert.Equal(t, "홍길동", p.BuyerName)
	assert.Equal(t, "01181", p.BuyerPostcode)
	assert.Equal(t, "010-4242-4242", p.BuyerTel)
	assert.False(t, p.CashReceiptIssued)
	assert.Equal(t, "pc", p.Channel)
	assert.Equal(t, "KRW", p.Currency)
	assert.True(t, p.Escrow)
	assert.Equal(t, 0, int(p.FailedAt))
	assert.Equal(t, "imp_785510843101", p.ImpUid)
	assert.Equal(t, "ORD20180131-0009728", p.MerchantUid)
	assert.Equal(t, "노르웨이 회전 의자", p.Name)
	assert.Equal(t, 1600828565, int(p.PaidAt))
	assert.Equal(t, "card", p.PayMethod)
	assert.Equal(t, "T0000", p.PgId)
	assert.Equal(t, "kcp", p.PgProvider)
	assert.Equal(t, "20517935943730", p.PgTid)
	assert.Equal(t, "https://admin8.kcp.co.kr/assist/bill.BillActionNew.do?cmd=card_bill&tno=20517935943730&order_no=imp_785510843101&trade_mony=100", p.ReceiptUrl)
	assert.Equal(t, 1600828510, int(p.StartedAt))
	assert.Equal(t, "paid", p.Status)
	assert.Equal(t, "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.83 Safari/537.36", p.UserAgent)
	assert.Equal(t, "", p.VbankCode)
	assert.Equal(t, 0, int(p.VbankDate))
	assert.Equal(t, "", p.VbankHolder)
	assert.Equal(t, 0, int(p.VbankIssuedAt))
	assert.Equal(t, "", p.VbankName)
	assert.Equal(t, "", p.VbankNum)
}

func checkPaymentImpUID338103167934(t *testing.T, p *payment.Payment) {
	assert.Equal(t, 200, int(p.Amount))
	assert.Equal(t, "서울특별시 강남구 삼성동", p.BuyerAddr)
	assert.Equal(t, "johny@chai.finance", p.BuyerEmail)
	assert.Equal(t, "구매자이름", p.BuyerName)
	assert.Equal(t, "123-456", p.BuyerPostcode)
	assert.Equal(t, "010-1234-5678", p.BuyerTel)
	assert.False(t, p.CashReceiptIssued)
	assert.Equal(t, "pc", p.Channel)
	assert.Equal(t, "KRW", p.Currency)
	assert.False(t, p.Escrow)
	assert.Equal(t, 0, int(p.FailedAt))
	assert.Equal(t, "imp_338103167934", p.ImpUid)
	assert.Equal(t, "merchant_1601964102920", p.MerchantUid)
	assert.Equal(t, "주문명:결제테스트", p.Name)
	assert.Equal(t, 1601964198, int(p.PaidAt))
	assert.Equal(t, "vbank", p.PayMethod)
	assert.Equal(t, "INIpayTest", p.PgId)
	assert.Equal(t, "html5_inicis", p.PgProvider)
	assert.Equal(t, "StdpayVBNKINIpayTest20201006150251995341", p.PgTid)
	assert.Equal(t, "https://iniweb.inicis.com/DefaultWebApp/mall/cr/cm/mCmReceipt_head.jsp?noTid=&noMethod=1", p.ReceiptUrl)
	assert.Equal(t, 1601964103, int(p.StartedAt))
	assert.Equal(t, "paid", p.Status)
	assert.Equal(t, "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.121 Safari/537.36", p.UserAgent)
	assert.Equal(t, "004", p.VbankCode)
	assert.Equal(t, 1604588399, int(p.VbankDate))
	assert.Equal(t, "（주）케이지이니시", p.VbankHolder)
	assert.Equal(t, 1601964172, int(p.VbankIssuedAt))
	assert.Equal(t, "KB 국민은행", p.VbankName)
	assert.Equal(t, "79959078731512", p.VbankNum)
}
