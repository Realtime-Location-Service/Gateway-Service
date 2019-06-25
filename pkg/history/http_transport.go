package history

import (
	"context"
	"io/ioutil"
	"net/http"

	"github.com/go-chi/chi"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/rls/gateway-service/pkg/config"
	httpSvc "github.com/rls/gateway-service/svc/http"
	"github.com/rls/gateway-service/utils/consts"
	httputil "github.com/rls/gateway-service/utils/http"
)

// MakeHandler returns a handler for the history service.
func MakeHandler(svc Service) http.Handler {
	reqHandler := kithttp.NewServer(
		makeHistoryProxyEndpoint(svc),
		decodeRequest,
		encodeResponse,
	)

	r := chi.NewRouter()
	r.Group(func(r chi.Router) {
		r.Method("GET", "/", reqHandler)
	})

	return r
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	res := response.(*historyResponse)
	return httpSvc.NewHTTP(0).SendResponse(w, res.HTTPResponse)
}

func decodeRequest(_ context.Context, r *http.Request) (i interface{}, e error) {
	var req historyRequest
	req.Referrer = r.Header.Get(consts.RLSReferrer)
	req.UserID = r.Header.Get(consts.UserID)
	req.SubbordinateIDs = r.Header.Get(consts.SubordinateIDs)
	req.Method = r.Method
	req.URL = httputil.BuildURL(config.HistoryCfg().BaseURL, r.URL.String(), "/api")
	payload, err := ioutil.ReadAll(r.Body)
	if err != nil {
		hr := historyResponse{}
		hr.StatusCode = http.StatusInternalServerError
		hr.Err = err
		return hr, nil
	}
	req.Payload = payload
	return req, nil
}
