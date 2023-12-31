package sep

import (
	"encoding/xml"
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
