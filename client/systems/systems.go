package systems

import (
	"egot-golang/ecs"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"time"
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

	if e.DevicCapability.Time != nil {
		e.TimeLink = e.DevicCapability.Time
	}

	if e.DevicCapability.EndDevices != nil {
		e.EndDeviceListLink = e.DevicCapability.EndDevices
	}

	if e.DevicCapability.SelfDevice != nil {
		e.SelfDeviceLink = e.DevicCapability.SelfDevice
	}

}

func GetTimeLink(world *ecs.World, client *http.Client) {
	e, found := world.GetTaggedEntity("root")
	if !found {
		fmt.Println("No entities found")
		return
	}

	if e.TimeLink == nil {
		return
	}

	resp, err := client.Get("https://localhost:4443" + e.TimeLink.Href)
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

	//err = xml.Unmarshal(body, e.Time)
	if err != nil {
		fmt.Println(err)
		return
	}

	e.TimeLink = nil

}

func GetSelfDeviceLink(world *ecs.World, client *http.Client) {
	e, found := world.GetTaggedEntity("root")
	if !found {
		fmt.Println("No entities found")
		return
	}

	if e.SelfDeviceLink == nil {
		return
	}

	resp, err := client.Get("https://localhost:4443" + e.SelfDeviceLink.Href)
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

	//err = xml.Unmarshal(body, e.SelfDevice)
	if err != nil {
		fmt.Println(err)
		return
	}

	e.SelfDeviceLink = nil

}

func GetEndDeviceListLink(world *ecs.World, client *http.Client) {
	e, found := world.GetTaggedEntity("root")
	if !found {
		fmt.Println("No entities found")
		return
	}

	if e.EndDeviceListLink == nil {
		return
	}

	resp, err := client.Get("https://localhost:4443" + e.EndDeviceListLink.Href)
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

	//err = xml.Unmarshal(body, e.EndDeviceList)
	if err != nil {
		fmt.Println(err)
		return
	}

	e.EndDeviceListLink = nil

}
