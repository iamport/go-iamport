package iamport

import (
	"errors"

	"github.com/iamport/go-iamport/certification"
	"github.com/iamport/go-iamport/util"
	TypeCertification "github.com/iamport/interface/gen_src/go/v1/certification"
)

func (iamport *Iamport) GetCertificationByImpUID(impUID string) (*TypeCertification.Certification, error) {
	if impUID == "" {
		return nil, errors.New(ErrMustExistImpUID)
	}

	token, err := iamport.Authenticate.GetToken()
	if err != nil {
		return nil, err
	}

	req := &TypeCertification.CertificationRequest{
		ImpUid: impUID,
	}

	res, err := certification.GetCertifications(
		iamport.Authenticate.Client,
		iamport.Authenticate.APIUrl,
		token,
		req,
	)

	if err != nil {
		return nil, err
	}
	if res.Code != util.CodeOK {
		return nil, errors.New(res.Message)
	}

	return res.Response, nil
}

func (iamport *Iamport) DeleteCertificationByImpUID(impUID string) (*TypeCertification.Certification, error) {
	if impUID == "" {
		return nil, errors.New(ErrMustExistImpUID)
	}

	token, err := iamport.Authenticate.GetToken()
	if err != nil {
		return nil, err
	}

	req := &TypeCertification.CertificationRequest{
		ImpUid: impUID,
	}

	res, err := certification.DeleteCertifications(
		iamport.Authenticate.Client,
		iamport.Authenticate.APIUrl,
		token,
		req,
	)

	if err != nil {
		return nil, err
	}
	if res.Code != util.CodeOK {
		return nil, errors.New(res.Message)
	}

	return res.Response, nil
}
