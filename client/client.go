package main

import (
	"bufio"
	"crypto/tls"
	"crypto/x509"
	"encoding/xml"
	"fmt"
	"net/http"
	"os"
)

type CustomerAccountListLink struct {
	XMLName xml.Name `xml:"CustomerAccountListLink"`
	all     int      `xml:"all,attr"`
	href    string   `xml:"href,attr"`
}

type DemandResponseProgramListLink struct {
	XMLName xml.Name `xml:"DemandResponseProgramListLink"`
	all     int      `xml:"all,attr"`
	href    string   `xml:"href,attr"`
}

type DERProgramListLink struct {
	XMLName xml.Name `xml:"DERProgramListLink"`
	all     int      `xml:"all,attr"`
	href    string   `xml:"href,attr"`
}

type FileListLink struct {
	XMLName xml.Name `xml:"FileListLink"`
	all     int      `xml:"all,attr"`
	href    string   `xml:"href,attr"`
}

type MessagingProgramListLink struct {
	XMLName xml.Name `xml:"MessagingProgramListLink"`
	all     int      `xml:"all,attr"`
	href    string   `xml:"href,attr"`
}

type PrepaymentListLink struct {
	XMLName xml.Name `xml:"PrepaymentListLink"`
	all     int      `xml:"all,attr"`
	href    string   `xml:"href,attr"`
}

type ResponseSetListLink struct {
	XMLName xml.Name `xml:"ResponseSetListLink"`
	all     int      `xml:"all,attr"`
	href    string   `xml:"href,attr"`
}

type TariffProfileListLink struct {
	XMLName xml.Name `xml:"TariffProfileListLink"`
	all     int      `xml:"all,attr"`
	href    string   `xml:"href,attr"`
}

type TimeLink struct {
	XMLName xml.Name `xml:"TimeLink"`
	href    string   `xml:"href,attr"`
}

type UsagePointListLink struct {
	XMLName xml.Name `xml:"UsagePointListLink"`
	all     int      `xml:"all,attr"`
	href    string   `xml:"href,attr"`
}

type EndDeviceListLink struct {
	XMLName xml.Name `xml:"EndDeviceListLink"`
	all     int      `xml:"all,attr"`
	href    string   `xml:"href,attr"`
}

type MirrorUsagePointListLink struct {
	XMLName xml.Name `xml:"MirrorUsagePointListLink"`
	all     int      `xml:"all,attr"`
	href    string   `xml:"href,attr"`
}

type SelfDeviceLink struct {
	XMLName xml.Name `xml:"SelfDeviceLink"`
	href    string   `xml:"href,attr"`
}

type DeviceCapability struct {
	XMLName   xml.Name `xml:"DeviceCapability"`
	poll_rate int      `xml:"pollRate,attr"`
	href      string   `xml:"href,attr"`
	ca_ll     CustomerAccountListLink
	drp_ll    DemandResponseProgramListLink
	derp_ll   DERProgramListLink
	file_ll   FileListLink
	msg_ll    MessagingProgramListLink
	pp_ll     PrepaymentListLink
	rsps_ll   ResponseSetListLink
	tariff_ll TariffProfileListLink
	time      TimeLink
	up_ll     UsagePointListLink
	edev_ll   EndDeviceListLink
	mup_ll    MirrorUsagePointListLink
	sdev      SelfDeviceLink
}

func main() {
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

	resp, err := client.Get("https://localhost:4443/dcap")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("Response status:", resp.Status)

	scanner := bufio.NewScanner(resp.Body)
	for i := 0; scanner.Scan() && i < 5; i++ {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
