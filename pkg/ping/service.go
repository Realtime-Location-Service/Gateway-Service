package ping

import (
	"context"

	chttp "github.com/rls/gateway-service/svc/http"
	"github.com/rls/gateway-service/utils/consts"
)

// Service is the interface that provides passes.
type Service interface {
	Request(context.Context, *pingRequest) (*pingResponse, error)
}

type service struct {
	cHTTP chttp.IHttp
}

func (svc *service) Request(ctx context.Context, r *pingRequest) (*pingResponse, error) {
	resp := svc.cHTTP.Do(r.Method, r.URL, r.Payload, map[string]string{
		consts.RLSReferrer: r.Referrer,
		consts.ContentType: consts.ContentTypeJSON,
	})
	return &pingResponse{
		resp,
	}, nil
}

// NewService creates pass service with necessary dependencies.
func NewService(cHTTP chttp.IHttp) Service {
	return &service{cHTTP}
}
