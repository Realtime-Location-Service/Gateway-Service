package ping

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

// MakeHandler returns a handler for the pass service.
func MakeHandler(svc Service) http.Handler {
	reqHandler := kithttp.NewServer(
		makePingProxyEndpoint(svc),
		decodeGetPassRequest,
		encodePassResponse,
	)

	r := chi.NewRouter()
	r.Group(func(r chi.Router) {
		r.Method("POST", "/", reqHandler)
		r.Method("GET", "/", reqHandler)
		r.Method("GET", "/users", reqHandler)
	})

	return r
}

func encodePassResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	res := response.(*pingResponse)
	return httpSvc.NewHTTP(0).SendResponse(w, res.HTTPResponse)
}

func decodeGetPassRequest(_ context.Context, r *http.Request) (i interface{}, e error) {
	var req pingRequest
	req.Referrer = r.Header.Get(consts.RLSReferrer)
	req.UserID = r.Header.Get(consts.UserID)
	req.SubbordinateIDs = r.Header.Get(consts.SubordinateIDs)
	req.Method = r.Method
	req.URL = httputil.BuildURL(config.PingCfg().BaseURL, r.URL.String())
	payload, err := ioutil.ReadAll(r.Body)
	if err != nil {
		pr := pingResponse{}
		pr.StatusCode = http.StatusInternalServerError
		pr.Err = err
		return pr, nil
	}
	req.Payload = payload
	return req, nil
}
