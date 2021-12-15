package main

import (
	"flag"
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

var secretsPath = "./certs/private.json"
var certPath = "./certs/butlerBurtonCert.pfx"

func main() {
	port := flag.String("port", "6666", "set port to run on")
	flag.Parse()
	http.HandleFunc("/secrets", secretsHandler)
	http.HandleFunc("/cert", certHandler)
	err := http.ListenAndServe(":"+*port, nil)
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
	w.Write(file)
}

func secretsHandler(w http.ResponseWriter, r *http.Request) {
	json, err := ioutil.ReadFile(secretsPath)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	w.Header().Add("content-type", "application/json")
	w.Write(json)
}
