package meta

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/rls/gateway-service/store/model"
)

type metaRequest struct {
	URL             string
	Method          string
	UserID          string
	SubbordinateIDs string
	Referrer        string
	Payload         []byte
}

type metaResponse struct {
	*model.HTTPResponse
}

func makeMetaProxyEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(metaRequest)
		if !ok {
			return request, nil
		}
		return svc.Request(ctx, &req)
	}
}
