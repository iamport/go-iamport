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
	HeaderContentTypeJson = "application/json"
	HeaderAuthorization   = "Authorization"

	GET  = "GET"
	POST = "POST"
	PUT  = "PUT"
)

type Method string

func Call(client *http.Client, token string, url string, method Method) ([]byte, error) {
	req, err := http.NewRequest(string(method), url, nil)

	if err != nil {
		return []byte{}, err
	}

	req.Header.Set(HeaderAuthorization, token)

	res, err := call(client, req)

	return res, nil
}

func CallWithForm(client *http.Client, token string, url string, method Method, form *url.Values) ([]byte, error) {

	req, err := http.NewRequest(string(method), url, bytes.NewBufferString(form.Encode()))
	if err != nil {
		return []byte{}, err
	}

	req.Header.Set(HeaderAuthorization, token)
	req.Header.Set(HeaderContentType, HeaderContentTypeForm)

	res, err := call(client, req)
	if err != nil {
		return []byte{}, err
	}

	return res, nil
}

func CallWithJson(client *http.Client, token string, url string, method Method, param []byte) ([]byte, error) {
	req, err := http.NewRequest(string(method), url, bytes.NewReader(param))
	if err != nil {
		return []byte{}, err
	}

	req.Header.Set(HeaderAuthorization, token)
	req.Header.Set(HeaderContentType, HeaderContentTypeJson)

	res, err := call(client, req)
	if err != nil {
		return []byte{}, err
	}

	return res, nil
}

func call(client *http.Client, req *http.Request) ([]byte, error) {
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
