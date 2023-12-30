package main

import (
	"encoding/xml"
	"log"
	"net/http"
)

type EndDeviceListLink struct {
	XMLName xml.Name `xml:"EndDeviceListLink"`
	All     int      `xml:"all,attr"`
	Href    string   `xml:"href,attr"`
}

type TimeLink struct {
	XMLName xml.Name `xml:"TimeLink"`
	Href    string   `xml:"href,attr"`
}

type SelfDeviceLink struct {
	XMLName xml.Name `xml:"SelfDeviceLink"`
	Href    string   `xml:"href,attr"`
}

type DeviceCapability struct {
	XMLName    xml.Name          `xml:"DeviceCapability"`
	Poll_rate  int               `xml:"pollRate,attr"`
	Href       string            `xml:"href,attr"`
	Time       TimeLink          `xml:"TimeLink"`
	Self       SelfDeviceLink    `xml:"SelfDeviceLink"`
	EndDevices EndDeviceListLink `xml:"EndDeviceListLink"`
}

func dcapHandler(w http.ResponseWriter, req *http.Request) {
	dcap := &DeviceCapability{
		Href:       "/dcap",
		Poll_rate:  900,
		Time:       TimeLink{Href: "/tm"},
		Self:       SelfDeviceLink{Href: "/sdev"},
		EndDevices: EndDeviceListLink{Href: "/edev", All: 1},
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
	err := http.ListenAndServeTLS(":4443", "./ssl/srv.crt", "./ssl/srv.key", nil)
	if err != nil {
		log.Fatal(err)
	}
}
