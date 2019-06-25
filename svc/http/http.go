package http

import (
	"bytes"
	"io"
	shttp "net/http"
	"time"

	"github.com/rls/gateway-service/utils/consts"
	"github.com/rls/gateway-service/utils/errors"

	"github.com/rls/gateway-service/store/model"
)

// SupportedDoMethods for Do request
var SupportedDoMethods = map[string]bool{
	shttp.MethodGet:   true,
	shttp.MethodPost:  true,
	shttp.MethodPatch: true,
}

// IHttp ...
type IHttp interface {
	Get(url string, headers map[string]string) *model.HTTPResponse
	Do(method, url string, data []byte, headers map[string]string) *model.HTTPResponse
	SendResponse(shttp.ResponseWriter, *model.HTTPResponse) error
}

type http struct {
	client *shttp.Client
}

func (h *http) Get(url string, headers map[string]string) *model.HTTPResponse {
	return h.Do(shttp.MethodGet, url, nil, headers)
}

func (h *http) Do(method, url string, data []byte, headers map[string]string) *model.HTTPResponse {

	if _, ok := SupportedDoMethods[method]; !ok {
		return &model.HTTPResponse{
			Body:       nil,
			Err:        errors.New("invalid method"),
			StatusCode: shttp.StatusBadRequest,
		}
	}

	req, err := shttp.NewRequest(method, url, bytes.NewBuffer(data))
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	resp, err := h.client.Do(req)
	if err != nil {
		return &model.HTTPResponse{
			Body:       nil,
			Err:        err,
			StatusCode: shttp.StatusInternalServerError,
			Header:     nil,
		}
	}

	return &model.HTTPResponse{
		Body:       resp.Body,
		Err:        err,
		StatusCode: resp.StatusCode,
		Header:     resp.Header,
	}
}

func (h *http) SendResponse(w shttp.ResponseWriter, resp *model.HTTPResponse) error {
	if resp == nil {
		return errors.New("response is nil")
	}
	w.Header().Set(consts.ContentType, consts.ContentTypeJSON)
	w.WriteHeader(resp.StatusCode)

	if resp.Err != nil {
		return nil
	}
	for k, v := range resp.Header {
		w.Header().Set(k, v[0])
	}

	io.Copy(w, resp.Body)
	resp.Body.Close()
	return nil
}

// NewHTTP returns custom http interface
func NewHTTP(requestTimeout time.Duration) IHttp {
	return &http{
		&shttp.Client{
			Timeout: requestTimeout,
		},
	}
}
