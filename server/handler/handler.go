package handler

import (
	"crypto/sha256"
	"egot-golang/ecs"
	"encoding/xml"
	"fmt"
	"github.com/Tylores/sep-models/sep"
	"net/http"
)

func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
    w.WriteHeader(status)
    if status == http.StatusNotFound {
        fmt.Fprint(w, "custom 404")
    }
}

func DeviceCapabilityHandler(world *ecs.World) http.HandlerFunc {
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

func TimeHandler(world *ecs.World) http.HandlerFunc {
	fn := func(w http.ResponseWriter, req *http.Request) {
		cert := req.TLS.PeerCertificates[0]
		fingerprint := fmt.Sprintf("%X", sha256.Sum256(cert.Raw))[0:40]

		e, found := world.GetTaggedEntity(fingerprint)
		if !found {
			req.	
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
func DeviceCapabilityHandler(world *ecs.World) http.HandlerFunc {
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
func DeviceCapabilityHandler(world *ecs.World) http.HandlerFunc {
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
