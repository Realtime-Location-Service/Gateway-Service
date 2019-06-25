package ping

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/rls/gateway-service/store/model"
)

type pingRequest struct {
	URL             string
	Method          string
	UserID          string
	SubbordinateIDs string
	Referrer        string
	Payload         []byte
}

type pingResponse struct {
	*model.HTTPResponse
}

func makePingProxyEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(pingRequest)
		if !ok {
			return request, nil
		}
		return svc.Request(ctx, &req)
	}
}
