package subscribe

import (
	"strings"

	"google.golang.org/protobuf/encoding/protojson"

	"github.com/iamport/go-iamport/util"

	"github.com/iamport/go-iamport/authenticate"
	"github.com/iamport/interface/gen_src/go/v1/subscribe"
)

const (
	URLSubscribe = "/subscribe"
	URLPayments  = "/payments"
	URLOnetime   = "/onetime"
)

func Onetime(auth *authenticate.Authenticate, params *subscribe.OnetimePaymentRequest) (*subscribe.OnetimePaymentResponse, error) {
	urls := []string{auth.APIUrl, URLSubscribe, URLPayments, URLOnetime}
	url := strings.Join(urls, "")

	token, err := auth.GetToken()
	if err != nil {
		return nil, err
	}

	marshaler := protojson.MarshalOptions{
		UseProtoNames: true,
	}
	jsonBytes, err := marshaler.Marshal(params)
	res, err := util.CallWithForm(auth.Client, token, url, util.POST, jsonBytes)
	if err != nil {
		return nil, err
	}

	ontimeRes := subscribe.OnetimePaymentResponse{}
	err = protojson.Unmarshal(res, &ontimeRes)
	if err != nil {
		return nil, err
	}

	return &ontimeRes, nil
}
