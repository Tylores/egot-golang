package main

import (
	"crypto/sha256"
	"crypto/tls"
	"egot-golang/ecs"
	"encoding/xml"
	"fmt"
	"github.com/Tylores/sep-models/sep"
	"log"
	"net/http"
)

func dcapHandler(w http.ResponseWriter, req *http.Request) {
	cert := req.TLS.PeerCertificates[0]
	fingerprint := fmt.Sprintf("%X", sha256.Sum256(cert.Raw))

	fmt.Printf("%s", fingerprint) // to make sure it's a hex string

	dcap := &sep.DeviceCapability{}

	out, err := xml.MarshalIndent(dcap, " ", "  ")
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/sep-xml")
	w.Write([]byte(out))
}

func worldHandler(world *ecs.World) http.HandlerFunc {
	fn := func(w http.ResponseWriter, req *http.Request) {
		cert := req.TLS.PeerCertificates[0]
		fingerprint := fmt.Sprintf("%X", sha256.Sum256(cert.Raw))[0:40]

		e, found := world.GetTaggedEntity(fingerprint)
		if !found {
			e = world.AddEntity(fingerprint)

			e.DevicCapability = &sep.DeviceCapability{
				Href:     "/dcap",
				PollRate: 9,
				FunctionSetAssignmentsBase: sep.FunctionSetAssignmentsBase{
					Time: &sep.TimeLink{
						Href: "/tm",
					},
				},
				SelfDevice: &sep.SelfDeviceLink{Href: "/sdev"},
				EndDevices: &sep.EndDeviceListLink{Href: "/edev", All: 1},
			}
		}

		out, err := xml.MarshalIndent(e.DevicCapability, " ", "  ")
		if err != nil {
			fmt.Println(err)
		}
		w.Header().Set("Content-Type", "application/sep-xml")
		w.Write([]byte(out))
	}

	return http.HandlerFunc(fn)
}
func main() {
	world := ecs.NewWorld()

	http.HandleFunc("/dcap", worldHandler(world))
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
