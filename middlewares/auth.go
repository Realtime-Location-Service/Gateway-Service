package middlewares

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/rls/gateway-service/pkg/config"
	"github.com/rls/gateway-service/store/model"
	httpSvc "github.com/rls/gateway-service/svc/http"
	"github.com/rls/gateway-service/utils/consts"
)

// ResolveUser will fetch user by appkey
// returns unauthorized otherthan 200 statuscode
func ResolveUser(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				return
			}
		}()

		w.Header().Set("Content-Type", consts.ContentTypeJSON)
		appKey := r.Header.Get(consts.AppKey)
		if appKey == "" {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		authCfg := config.AuthCfg()
		// url := chttp.BuildURL(cfg.URL, r.URL.String())
		resp := httpSvc.NewHTTP(authCfg.RequestTimeout).Get(authCfg.AuthURL+"?appKey="+appKey, map[string]string{
			consts.AppKey: appKey,
		})

		if resp.Err != nil || resp.StatusCode != http.StatusOK {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}

		defer resp.Body.Close()

		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}

		user := model.User{}
		if err := json.Unmarshal(b, &user); err != nil {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}
		r.Header.Set(consts.RLSReferrer, user.Domain)
		// r.Header.Add(consts.UserID, user.UserID)
		// r.Header.Add(consts.SubordinateIDs, user.SubordinateIDs())
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
