# go-iamport

Go Language 아임포트 Rest API Client  
https://api.iamport.kr

## 설치

    $ go get github.com/joowonyun/go-iamport

## 예제
    iam, err := iamport.NewIamport("https://api.iamport.kr", "<your_api_key>", "<your_api_secret>")
    if err != nil {
      return err
    }

    pay, err := iam.GetPaymentImpUID("<some imp_uid>")
    if err != nil {
      fmt.Println(err)
      return
    }

    fmt.Println(pay.Amount)
    fmt.Println(pay.MerchantUID)

## 구현되어있는 기능 - https://api.iamport.kr

- authenticate
  - POST /users/getToken
- payments  
  - GET /payments/{imp_uid}
  - GET /payments
  - GET /payments/find/{merchant_uid}
  - GET /payments/findAll/{merchant_uid}/{payment_status}
  - GET /payments/status/{payment_status}
  - GET /payments/{imp_uid}/balance

### TODO
- payments
  - POST /payments/cancel
- payments.validation
  - POST /payments/prepare
  - GET /payments/prepare/merchant_uid
- subscribe
  - POST /subscribe/payments/ontime
  - POST /subscribe/payments/again
  - POST /subscribe/payments/schedule
  - POST /subscribe/payments/unschedule
- subscribe.customer
  - DELETE /subscribe/customers/{customer_uid}
  - GET /subscribe/customers/{customer_uid}
  - POST /subscribe/customers/{customer_uid}
