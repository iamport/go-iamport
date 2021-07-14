package certification

import (
	"testing"

	"github.com/iamport/go-iamport/authenticate"
	"github.com/iamport/interface/gen_src/go/v1/certification"
	"github.com/stretchr/testify/assert"
)

func xTestGetCertifications(t *testing.T) {
	auth := authenticate.GetMockBaseAuthenticate()

	token, err := auth.GetToken()
	assert.NoError(t, err)

	params := &certification.CertificationRequest{
		ImpUid: "imp_939214393899",
	}

	res, err := GetCertifications(auth.Client, auth.APIUrl, token, params)
	assert.NoError(t, err)
	assert.Equal(t, 0, res.Code)
	assert.Equal(t, true, res.Response.Certified)
}
