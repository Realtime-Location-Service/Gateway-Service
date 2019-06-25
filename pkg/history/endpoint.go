package history

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/rls/gateway-service/store/model"
)

type historyRequest struct {
	URL             string
	Method          string
	UserID          string
	SubbordinateIDs string
	Referrer        string
	Payload         []byte
}

type historyResponse struct {
	*model.HTTPResponse
}

func makeHistoryProxyEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(historyRequest)
		if !ok {
			return request, nil
		}
		return svc.Request(ctx, &req)
	}
}
