package services

import (
	"encoding/json"
	"net/http"

	"github.com/auth-user-proxy-service/user-microservice/src/config"
	"github.com/auth-user-proxy-service/user-microservice/src/model"
	"github.com/auth-user-proxy-service/user-microservice/src/utils"
	"github.com/gorilla/mux"
)

// NewRouter creates a router for URL-to-service mapping
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(false)

	router.HandleFunc("/user/profile", UserProfileHandler).Methods(http.MethodGet)
	router.HandleFunc("/microservice/name", ServiceNameHandler).Methods(http.MethodGet)
	return router
}

// UserProfileHandler used for returning version of service
func UserProfileHandler(w http.ResponseWriter, r *http.Request) {
	defer func() {
		err := recover()
		if err != nil {
			http.Error(w, utils.ToString(err), http.StatusInternalServerError)
			return
		}
	}()

	userName := r.Header.Get("user")
	if userName == config.Config.Username {
		profile := config.Config.Profile
		jsonByte, err := json.Marshal(profile)
		if err != nil {
			http.Error(w, model.Servererrstr, http.StatusInternalServerError)
		}
		utils.GenerateResponse(jsonByte, http.StatusOK, w)
	}
	return
}

// ServiceNameHandler used for returning version of service
func ServiceNameHandler(w http.ResponseWriter, _ *http.Request) {
	defer func() {
		err := recover()
		if err != nil {
			http.Error(w, utils.ToString(err), http.StatusInternalServerError)
			return
		}
	}()

	jsonByte, err := json.Marshal("user-miscroservice")
	if err != nil {
		http.Error(w, model.Servererrstr, http.StatusInternalServerError)
	}
	utils.GenerateResponse(jsonByte, http.StatusOK, w)
	return
}
