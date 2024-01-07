package main

import (
	"crypto/tls"
	"crypto/x509"
	"egot-golang/client/systems"
	"egot-golang/ecs"
	"fmt"
	"github.com/Tylores/sep-models/sep"
	"net/http"
	"os"
	"time"
)

func main() {
	world := ecs.NewWorld()
	dcap := world.AddEntity("root")
	dcap.DevicCapability = &sep.DeviceCapability{
		Href:     "/dcap",
		PollRate: 5,
	}

	caCert, _ := os.ReadFile("./ssl/ca.crt")
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	cert, _ := tls.LoadX509KeyPair("./ssl/client.crt", "./ssl/client.key")

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs:      caCertPool,
				Certificates: []tls.Certificate{cert},
			},
		},
	}

	for true {
		start := time.Now().UnixMicro()
		world.Update()
		systems.GetDeviceCapability(world, client)
		systems.GetTimeLink(world, client)
		systems.GetEndDeviceListLink(world, client)
		systems.GetSelfDeviceLink(world, client)
		stop := time.Now().UnixMicro()
		duration := stop - start
		fmt.Println(duration)
		time.Sleep(time.Second - time.Duration(duration))
	}
}
