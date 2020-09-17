package util

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	CodeOK = 0

	ErrStatusUnauthorized = "iamport: unauthorized"
	ErrStatusNotFound     = "iamport: invalid imp_uid"
	ErrUnknown            = "iamport: unknown error"

	HeaderContentType     = "Content-Type"
	HeaderContentTypeForm = "application/x-www-form-urlencoded"
	HeaderAuthorization   = "Authorization"

	GET  = "GET"
	POST = "POST"
)

func CallGet(client *http.Client, token string, url string) ([]byte, error) {
	req, err := http.NewRequest(GET, url, nil)
	if err != nil {
		return []byte{}, err
	}
	req.Header.Set(HeaderAuthorization, token)

	res, err := client.Do(req)
	err = errorHandler(res)
	if err != nil {
		return []byte{}, err
	}

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return []byte{}, err
	}

	return resBody, nil
}

func CallPostForm(client *http.Client, token string, url string, form url.Values) ([]byte, error) {
	req, err := http.NewRequest(POST, url, bytes.NewBufferString(form.Encode()))
	if err != nil {
		return []byte{}, err
	}
	req.Header.Set(HeaderContentType, HeaderContentTypeForm)
	req.Header.Set(HeaderAuthorization, token)

	res, err := client.Do(req)
	err = errorHandler(res)
	if err != nil {
		return []byte{}, err
	}

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return []byte{}, err
	}

	return resBody, nil
}

func GetQueryPrefix(isFirst *bool) string {
	if *isFirst {
		*isFirst = false
		return "?"
	} else {
		return "&"
	}
}

func ValidateStatusParameter(src string) bool {
	if src == "" || src == StatusReady || src == StatusAll || src == StatusPaid || src == StatusFailed || src == StatusCanceled {
		return true
	}

	return false
}

func ValidateSortParameter(src string) bool {
	if src == "" || src == SortDESCStarted || src == SortASCStarted || src == SortDESCPaid || src == SortASCPaid || src == SortASCUpdated || src == SortDESCUpdated {
		return true
	}

	return false
}

func errorHandler(res *http.Response) error {
	switch res.StatusCode {
	case http.StatusOK:
		return nil
	case http.StatusUnauthorized:
		return errors.New(ErrStatusUnauthorized)
	case http.StatusNotFound:
		return errors.New(ErrStatusNotFound)
	default:
		return errors.New(ErrUnknown)
	}
}
