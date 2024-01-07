package main

import (
	"crypto/tls"
	"egot-golang/ecs"
	"egot-golang/server/handler"
	"log"
	"net/http"
)

func main() {
	world := ecs.NewWorld()

	http.HandleFunc("/dcap", handler.DeviceCapabilityHandler(world))
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
