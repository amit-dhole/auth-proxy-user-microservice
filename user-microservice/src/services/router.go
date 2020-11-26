package services

import (
	"encoding/json"
	"net/http"

	"github.com/auth-user-proxy-microservice/user-microservice/src/config"
	"github.com/auth-user-proxy-microservice/user-microservice/src/model"
	"github.com/auth-user-proxy-microservice/user-microservice/src/utils"
	"github.com/gorilla/mux"
)

// NewRouter creates a router for URL-to-service mapping
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(false)

	router.HandleFunc("/user/profile", UserProfileHandler).Methods(http.MethodGet)
	router.HandleFunc("/microservice/name", ServiceNameHandler).Methods(http.MethodGet)
	return router
}

// UserProfileHandler used for returning user profile
func UserProfileHandler(w http.ResponseWriter, r *http.Request) {
	defer func() {
		err := recover()
		if err != nil {
			http.Error(w, utils.ToString(err), http.StatusInternalServerError)
			return
		}
	}()

	userName := r.Header.Get("user")
	if userName == config.Config.User {
		profile := config.Config.Profile
		jsonByte, err := json.MarshalIndent(profile, " ", " ")
		if err != nil {
			http.Error(w, model.Servererrstr, http.StatusInternalServerError)
			return
		}
		utils.GenerateResponse(jsonByte, http.StatusOK, w)
	}
	return
}

// ServiceNameHandler used for returning name of service
func ServiceNameHandler(w http.ResponseWriter, _ *http.Request) {
	defer func() {
		err := recover()
		if err != nil {
			http.Error(w, utils.ToString(err), http.StatusInternalServerError)
			return
		}
	}()

	jsonByte, err := json.MarshalIndent("user-miscroservice", " ", " ")
	if err != nil {
		http.Error(w, model.Servererrstr, http.StatusInternalServerError)
		return
	}
	utils.GenerateResponse(jsonByte, http.StatusOK, w)
	return
}
