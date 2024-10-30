package middlewares

import (
	"net"
	"net/http"
	"strconv"

	"github.com/PereverzevIvan/razrabotka-biznes-prilojenii-konditerskaya/backend/hello/configs"
)

func IpWhitelistMiddleware(allowedOrigins []configs.AllowedOrigin, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		remoteIP, port, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		isAllowed := false
		for _, origin := range allowedOrigins {
			if remoteIP != origin.Host {
				continue
			}
			if origin.Port != 0 && port != strconv.Itoa(origin.Port) {
				continue
			}

			isAllowed = true
			break
		}

		if !isAllowed {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	})
}
