package services

import (
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/auth-user-proxy-service/proxy-microservice/src/config"
	"github.com/auth-user-proxy-service/proxy-microservice/src/model"
	"github.com/auth-user-proxy-service/proxy-microservice/src/utils"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
)

// NewRouter creates a router for URL-to-service mapping
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(false)

	router.HandleFunc("/", HomeHandler)
	return router
}

// HomeHandler used for returning version of service
func HomeHandler(w http.ResponseWriter, _ *http.Request) {

	defer func() {
		err := recover()
		if err != nil {
			http.Error(w, utils.ToString(err), http.StatusInternalServerError)
			return
		}
	}()

	//creating the proxyURL
	proxyStr := "http://localhost:8081"
	proxyURL, err := url.Parse(proxyStr)
	if err != nil {
		log.Println(err)
	}

	//creating the URL to be loaded through the proxy
	urlStr := "http://localhost:8081/auth"
	url, err := url.Parse(urlStr)
	if err != nil {
		log.Println(err)
	}

	//generating the HTTP GET request
	request, err := http.NewRequest(http.MethodGet, url.String(), nil)
	if err != nil {
		log.Println(err)
	}

	//adding proxy authentication
	auth := "root"
	basicAuth := "Basic " + base64.StdEncoding.EncodeToString([]byte(auth))
	request.Header.Add("Proxy-Authorization", basicAuth)

	//adding the proxy settings to the Transport object
	transport := &http.Transport{
		Proxy:           http.ProxyURL(proxyURL),
		TLSClientConfig: &tls.Config{InsecureSkipVerify: false},
	}
	transport.ProxyConnectHeader = request.Header

	//adding the Transport object to the http Client
	client := &http.Client{
		Transport: transport,
	}

	//printing the request to the console
	dump, _ := httputil.DumpRequest(request, true)
	fmt.Println(string(dump))

	//calling the URL

	resp, err := client.Do(request)
	if err != nil {
		log.Println(err)
	}

	if resp.StatusCode == http.StatusOK {
		data, _ := checkUserProfile(w, auth)
		jsonByte, err := json.Marshal(data)
		if err != nil {
			http.Error(w, model.Servererrstr, http.StatusInternalServerError)
		}
		utils.GenerateResponse(jsonByte, http.StatusOK, w)
	} else {
		data, _ := checkMicroserviceName(w)
		jsonByte, err := json.Marshal(data)
		if err != nil {
			http.Error(w, model.Servererrstr, http.StatusInternalServerError)
		}
		utils.GenerateResponse(jsonByte, http.StatusOK, w)
	}

}

func checkUserProfile(w http.ResponseWriter, username string) (config.Profile, error) {
	profile := config.Profile{}
	URL := "http://localhost:8083/user/profile"
	req, err := http.NewRequest(http.MethodGet, URL, nil)
	if err != nil {
		return profile, errors.Wrapf(errors.New("ErrHTTPGet"), fmt.Sprintf("Error in newrequest : %s", req), err)
	}
	req.Header.Add("user", username)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return profile, errors.Wrapf(errors.New("ErrHTTPGet"), fmt.Sprintf("Error in newrequest : %s", req), err)
	}
	if resp.StatusCode == http.StatusOK {
		defer func() {
			_ = resp.Body.Close()
		}()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return profile, err
		}

		data := config.Profile{}
		err = json.Unmarshal(body, &data)
		if err != nil {
			return profile, err
		}
		return data, nil
	}
	return profile, nil
}

func checkMicroserviceName(w http.ResponseWriter) (string, error) {
	URL := "http://localhost:8083/microservice/name"
	req, err := http.NewRequest(http.MethodGet, URL, nil)
	if err != nil {
		return "", errors.Wrapf(errors.New("ErrHTTPGet"), fmt.Sprintf("Error in request : %s", req), err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", errors.Wrapf(errors.New("ErrHTTPGet"), fmt.Sprintf("Error in Unmarshelling : %s", resp), err)
	}
	if resp.StatusCode == http.StatusOK {
		defer func() {
			_ = resp.Body.Close()
		}()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return "", err
		}
		var data string
		err = json.Unmarshal(body, &data)
		if err != nil {
			return "", err
		}
		return string(data), nil
	}
	return "", nil
}
