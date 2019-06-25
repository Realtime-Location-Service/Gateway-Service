package meta

import (
	"context"

	chttp "github.com/rls/gateway-service/svc/http"
	"github.com/rls/gateway-service/utils/consts"
)

// Service  ...
type Service interface {
	Request(context.Context, *metaRequest) (*metaResponse, error)
}

type service struct {
	cHTTP chttp.IHttp
}

func (svc *service) Request(ctx context.Context, r *metaRequest) (*metaResponse, error) {
	resp := svc.cHTTP.Do(r.Method, r.URL, r.Payload, map[string]string{
		consts.RLSReferrer: r.Referrer,
		consts.ContentType: consts.ContentTypeJSON,
	})
	return &metaResponse{
		resp,
	}, nil
}

// NewService ...
func NewService(cHTTP chttp.IHttp) Service {
	return &service{cHTTP}
}
