package escrow

import (
	"errors"
	"strings"

	"github.com/joowonyun/go-iamport/util"

	"google.golang.org/protobuf/encoding/protojson"

	"github.com/iamport/interface/gen_src/go/escrow"
	"github.com/joowonyun/go-iamport/authenticate"
)

const (
	URLEscrows = "/escrows"
	URLLogis   = "/logis"

	ParamsImpUID   = "imp_uid"
	ParamsSender   = "sender"
	ParamsReceiver = "receiver"
	ParamsLogis    = "logis"

	ErrUnknownMethod = "Unknown method"
)

type Method int

const (
	Register = iota
	Update
)

func Escrow(auth *authenticate.Authenticate, impUID string, params *escrow.EscrowRequest, method Method) (*escrow.EscrowResponse, error) {
	urls := []string{auth.APIUrl, URLEscrows, URLLogis, "/", impUID}
	urlEscrow := strings.Join(urls, "")

	token, err := auth.GetToken()
	if err != nil {
		return nil, err
	}

	option := protojson.MarshalOptions{
		UseProtoNames: true,
	}
	paramsBytes, err := option.Marshal(params)
	if err != nil {
		return nil, err
	}

	var restMethod util.Method
	switch method {
	case Register:
		restMethod = util.POST
		break
	case Update:
		restMethod = util.PUT
		break
	default:
		return nil, errors.New(ErrUnknownMethod)
	}

	res, err := util.CallWithJson(auth.Client, token, urlEscrow, restMethod, paramsBytes)
	if err != nil {
		return nil, err
	}

	escrowRes := escrow.EscrowResponse{}
	err = protojson.Unmarshal(res, &escrowRes)
	if err != nil {
		return nil, err
	}

	return &escrowRes, nil
}
