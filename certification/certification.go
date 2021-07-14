package certification

import (
	"net/http"
	"strings"

	"github.com/iamport/go-iamport/util"
	"github.com/iamport/interface/gen_src/go/v1/certification"
	"google.golang.org/protobuf/encoding/protojson"
)

const (
	URLCertification = "/certifications"
)

func GetCertifications(client *http.Client, apiDomain string, token string, params *certification.CertificationRequest) (*certification.CertificationResponse, error) {
	urls := []string{apiDomain, URLCertification, "/", params.ImpUid}
	url := strings.Join(urls, "")

	res, err := util.Call(client, token, url, util.GET)
	if err != nil {
		return nil, err
	}

	certificationRes := certification.CertificationResponse{}
	err = protojson.Unmarshal(res, &certificationRes)
	if err != nil {
		return nil, err
	}

	return &certificationRes, nil
}

func DeleteCertifications(client *http.Client, apiDomain string, token string, params *certification.CertificationRequest) (*certification.CertificationResponse, error) {
	urls := []string{apiDomain, URLCertification, "/", params.ImpUid}
	url := strings.Join(urls, "")

	res, err := util.Call(client, token, url, util.DELETE)
	if err != nil {
		return nil, err
	}

	certificationRes := certification.CertificationResponse{}
	err = protojson.Unmarshal(res, &certificationRes)
	if err != nil {
		return nil, err
	}

	return &certificationRes, nil
}
