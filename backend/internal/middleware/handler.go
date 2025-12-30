package middleware

import (
	"encoding/base64"
	"net/http"
	"strings"

	"github.com/rafaeldepontes/gopher-sub/internal/logger"
)

func AuthenticationFilter(next http.Handler) http.Handler {

	// I'm using base64 to enconde and decode the token because the focus here is not
	// security, I'm just trying to make a system with an email funcionality that it
	// will never be in production...
	//
	// If you want to see something that really uses the newest tecnologies for
	// token protection, check: github.com/rafaeldepontes/fauthless-go
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var accessType []byte

		dirtyToken := r.Header.Get("Authorization")
		if dirtyToken == "" {
			logger.GetLogger().Errorln("Missing token")
			http.Error(w, "Missing token", http.StatusUnauthorized)
			return
		}

		token := cleanUpToken(dirtyToken)

		accessType, err := base64.RawURLEncoding.DecodeString(token)
		if err != nil {
			logger.GetLogger().Errorln(err)
			http.Error(w, "Something went wrong, try again later...", http.StatusInternalServerError)
			return
		}

		if string(accessType) != "allowed" {
			logger.GetLogger().Errorln("User not allowed")
			http.Error(w, "User not allowed...", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func cleanUpToken(tk string) string {
	result, found := strings.CutPrefix(tk, "Bearer ")
	if !found {
		return tk
	}
	return result
}
