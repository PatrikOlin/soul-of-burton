package main

import (
	"io/ioutil"
	"net/http"
)

type Secrets struct {
	SiteURL  string `json:"siteUrl"`
	TenantID string `json:"tenantId"`
	ClientID string `json:"clientId"`
	CertPath string `json:"certPath"`
	CertPass string `json:"certPass"`
}

var secretsPath = "./certs/secrets.json"
var certPath = "./certs/butlerBurtonCert.pfx"

func main() {
	http.HandleFunc("/secrets", secretsHandler)
	http.HandleFunc("/cert", certHandler)
	err := http.ListenAndServe(":6666", nil)
	if err != nil {
		panic(err)
	}
}

func certHandler(w http.ResponseWriter, r *http.Request) {
	file, err := ioutil.ReadFile(certPath)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	w.Header().Add("content-type", "application/octet-stream")
	w.WriteHeader(http.StatusOK)
	w.Write(file)
}

func secretsHandler(w http.ResponseWriter, r *http.Request) {
	json, err := ioutil.ReadFile(secretsPath)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json)
}
