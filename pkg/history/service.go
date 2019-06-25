package history

import (
	"context"

	chttp "github.com/rls/gateway-service/svc/http"
	"github.com/rls/gateway-service/utils/consts"
)

// Service ...
type Service interface {
	Request(context.Context, *historyRequest) (*historyResponse, error)
}

type service struct {
	cHTTP chttp.IHttp
}

func (svc *service) Request(ctx context.Context, r *historyRequest) (*historyResponse, error) {
	resp := svc.cHTTP.Do(r.Method, r.URL, r.Payload, map[string]string{
		consts.RLSReferrer: r.Referrer,
		consts.ContentType: consts.ContentTypeJSON,
	})
	return &historyResponse{
		resp,
	}, nil
}

// NewService creates history service with necessary dependencies.
func NewService(cHTTP chttp.IHttp) Service {
	return &service{cHTTP}
}
