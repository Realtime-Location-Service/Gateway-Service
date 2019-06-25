package meta

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/go-chi/chi"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/rls/gateway-service/pkg/config"
	httpSvc "github.com/rls/gateway-service/svc/http"
	"github.com/rls/gateway-service/utils/consts"
	httputil "github.com/rls/gateway-service/utils/http"
)

// MakeHandler returns a handler for the  meta service.
func MakeHandler(svc Service) http.Handler {
	reqHandler := kithttp.NewServer(
		makeMetaProxyEndpoint(svc),
		decodeRequest,
		encodeResponse,
	)

	r := chi.NewRouter()
	r.Group(func(r chi.Router) {
		r.Method("POST", "/meta", reqHandler)
		r.Method("PATCH", "/{id}/meta", reqHandler)
		r.Method("GET", "/", reqHandler)
		r.Method("POST", "/meta/search", reqHandler)
	})

	return r
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	res := response.(*metaResponse)
	return httpSvc.NewHTTP(0).SendResponse(w, res.HTTPResponse)
}

func decodeRequest(_ context.Context, r *http.Request) (i interface{}, e error) {
	var req metaRequest
	req.Referrer = r.Header.Get(consts.RLSReferrer)
	req.UserID = r.Header.Get(consts.UserID)
	req.SubbordinateIDs = r.Header.Get(consts.SubordinateIDs)
	req.Method = r.Method
	req.URL = httputil.BuildURL(config.MetaCfg().BaseURL, r.URL.String(), "/api")

	fmt.Println("req.URL", req.URL)

	payload, err := ioutil.ReadAll(r.Body)
	if err != nil {
		mr := metaResponse{}
		mr.StatusCode = http.StatusInternalServerError
		mr.Err = err
		return mr, nil
	}
	req.Payload = payload
	return req, nil
}
