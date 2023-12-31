package main

import (
	"crypto/sha256"
	"crypto/tls"
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
	"sep-golang/sep"
)

func dcapHandler(w http.ResponseWriter, req *http.Request) {
	cert := req.TLS.PeerCertificates[0]
	fingerprint := fmt.Sprintf("%X", sha256.Sum256(cert.Raw))

	fmt.Printf("%s", fingerprint) // to make sure it's a hex string

	dcap := &sep.DeviceCapability{
		Href:       "/dcap",
		Poll_rate:  900,
		Time:       sep.TimeLink{Href: "/tm"},
		Self:       sep.SelfDeviceLink{Href: "/sdev"},
		EndDevices: sep.EndDeviceListLink{Href: "/edev", All: 1},
	}

	out, err := xml.MarshalIndent(dcap, " ", "  ")
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/sep-xml")
	w.Write([]byte(out))
}

func main() {
	http.HandleFunc("/dcap", dcapHandler)
	cfg := &tls.Config{
		MinVersion: tls.VersionTLS12,
		ClientAuth: tls.RequireAndVerifyClientCert,
	}
	server := http.Server{
		Addr:      ":4443",
		TLSConfig: cfg,
	}
	err := server.ListenAndServeTLS("./ssl/srv.crt", "./ssl/srv.key")
	if err != nil {
		log.Fatal(err)
	}
}
