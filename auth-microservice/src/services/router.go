package services

import (
	"encoding/base64"
	"net/http"
	"strings"

	"github.com/auth-user-proxy-microservice/auth-microservice/src/config"
	"github.com/auth-user-proxy-microservice/auth-microservice/src/utils"
	"github.com/gorilla/mux"
)

// NewRouter creates a router for URL-to-service mapping
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(false)

	router.HandleFunc("/auth", AuthHandler).Methods(http.MethodGet)
	return router
}

// AuthHandler used for returning authendication of user
func AuthHandler(w http.ResponseWriter, r *http.Request) {

	defer func() {
		err := recover()
		if err != nil {
			http.Error(w, utils.ToString(err), http.StatusInternalServerError)
			return
		}
	}()

	if config.Config.Authentication {
		authHeader := r.Header.Get("Proxy-Authorization")
		if len(strings.TrimSpace(authHeader)) > 0 {
			headerStr := strings.Split(authHeader, "Basic ")
			if len(headerStr) == 2 {
				credentials, err := base64.StdEncoding.DecodeString(headerStr[1])
				if err == nil {
					decodedCredentials := string(credentials)
					userName := decodedCredentials
					if userName == config.Config.User {
						utils.GenerateResponse(nil, http.StatusOK, w)
						return
					}
				}
			}
		}
		utils.GenerateResponse(nil, http.StatusUnauthorized, w)
		return
	}
	utils.GenerateResponse(nil, http.StatusOK, w)
}
