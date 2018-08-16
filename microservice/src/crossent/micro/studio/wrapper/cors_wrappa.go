package wrappa

import (
	"net/http"
	"strings"
	"crossent/micro/studio/domain"
	"fmt"
	"github.com/pivotal-cf/uaa-sso-golang/uaa"
)

func HttpWrap(handler http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		fmt.Println(fmt.Sprintf("[hyo] r.Method = %s, r.RequestURI = %s", r.Method, r.RequestURI))
		if origin := r.Header.Get("Origin"); origin != "" {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, PATCH, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, X-XSRF-TOKEN," +
				" U-TOKEN, U-X-TOKEN, Accept-Encoding, Authorization")
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			w.Header().Set("Access-Control-Expose-Headers", "X-XSRF-TOKEN, U-TOKEN, U-X-TOKEN")
		}
		if r.Method == "OPTIONS" {
			return
		}

		if !strings.Contains(r.RequestURI, "/web/") &&
			!strings.Contains(r.RequestURI, "/login") &&
			!strings.Contains(r.RequestURI, "/ping") {

			session := domain.SessionManager.Load(r)
			token, err := session.GetString(domain.UAA_TOKEN_NAME)
			if token == "" || err != nil {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusUnauthorized)
				var errStr string
				if err != nil {
					errStr = err.Error()
				}
				b := []byte(fmt.Sprintf("UAA Token does not exist. %v", errStr))
				w.Write(b)
				return
			} else {
				var uaaToken uaa.Token
				uaaToken.Access = token
				expired, err := uaaToken.IsExpired()
				if err != nil {
					session.Remove(w, domain.UAA_TOKEN_NAME)
					b := []byte(fmt.Sprintf("UAA token Error. %v", err))
					w.WriteHeader(http.StatusUnauthorized)
					w.Write(b)
					return
				}
				if expired {
					session.Remove(w, domain.UAA_TOKEN_NAME)
					b := []byte("UAA token has expired.")
					w.WriteHeader(http.StatusUnauthorized)
					w.Write(b)
					return
				}
			}
		}

		handler.ServeHTTP(w, r)
	}
}