package main

import (
	"crypto/tls"
	"crypto/x509"
	"egot-golang/ecs"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/Tylores/sep-models/sep"
)

func GetDeviceCapability(world *ecs.World, client *http.Client) {
	e, found := world.GetTaggedEntity("root")
	if !found {
		fmt.Println("No entities found")
		return
	}

	if time.Now().Unix()%int64(e.DevicCapability.PollRate) != 0 {
		fmt.Println("Not time yet")
		return
	}
	resp, err := client.Get("https://localhost:4443" + e.DevicCapability.Href)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		fmt.Println("Response status:", resp.Status)
		return
	}

	if resp.Header["Content-Type"][0] != "application/sep-xml" {
		fmt.Println("Response headers:", resp.Header["Content-Type"])
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(body))

	err = xml.Unmarshal(body, e.DevicCapability)
	if err != nil {
		fmt.Println(err)
		return
	}
}

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
		world.Update()
		GetDeviceCapability(world, client)
		time.Sleep(time.Second)
	}
}
